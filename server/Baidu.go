package server

import (
	"auxpi/auxpiAll"
	"auxpi/tools"
)

type Baidu struct {
}

func (this *Baidu) UploadToBaidu(img []byte, imgInfo string, imgType string) string {
	url := "http://image.baidu.com/pcdutu/a_upload?fr=html5&target=pcSearchImage&needJson=true"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "file",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	j := auxpi.BaiduResp{}
	j.UnmarshalJSON([]byte(data))
	url = "https://image.baidu.com/search/down?tn=download&url=" + j.Url
	return url
}
