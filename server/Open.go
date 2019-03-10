package server

import (
	"auxpi/auxpiAll"
	"auxpi/tools"

	"github.com/astaxie/beego/logs"
)

type OpenApi struct {
}

//存在 api 限制问题，暂时不考虑接入
func (this *OpenApi) UpLoadToPublicSina(img []byte, imgInfo string,imgType string) string {
	url := "https://apis.yum6.cn/api/5bd44dc94bcfc?token=f07b711396f9a05bc7129c4507fb65c5"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "file",
		Value: img,
	}
	var header map[string]string
	data :=tools.FormPost(file,url,header)
	j := auxpi.SinaPublicResponse{}
	j.UnmarshalJSON([]byte(data))
	pid,ok := j.Data["pid"].(string)
	if !ok {
		logs.Alert("上传公共图床出错")
		return ""
	}
	url = tools.CheckPid(pid,imgType)
	return url
}
