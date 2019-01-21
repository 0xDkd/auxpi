package server

import (
	"auxpi/auxpiAll"
	"auxpi/tools"

	"github.com/astaxie/beego"
)

type NetEasy struct {
}

func (this *NetEasy) UploadToNetEasy(img []byte, imgInfo string, imgType string) string {
	url := "http://you.163.com/xhr/file/upload.json"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "file",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	j := auxpi.NetEasyResp{}
	j.UnmarshalJSON([]byte(data))
	beego.Alert(j.Data[0])
	return j.Data[0]
}
