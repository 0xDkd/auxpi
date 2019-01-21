package server

import (
	"auxpi/auxpiAll"
	"auxpi/tools"
)

type Ali struct {
}

func (this *Ali) UploadToAli(img []byte, imgInfo string, imgType string) string {
	url := "https://kfupload.alibaba.com/mupload"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "file",
		Value: img,
		Type:  imgType,
	}
	//var header map[string]string
	data := tools.AliFormPost(file, url)
	j := auxpi.AliResp{}
	j.UnmarshalJSON([]byte(data))
	return j.Url
}
