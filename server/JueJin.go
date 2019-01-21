package server

import (
	"auxpi/auxpiAll"
	"auxpi/tools"
)

type JueJin struct {
}

func (this *JueJin) UploadToJueJin(img []byte, imgInfo string, imgType string) string {
	url := "https://cdn-ms.juejin.im/v1/upload?bucket=gold-user-assets"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "file",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	j := auxpi.JueJinResp{}
	j.UnmarshalJSON([]byte(data))

	//神奇三断言 : )
	reJ,_ := j.D.(map[string]interface{})
	urls,_ := reJ["url"].(map[string]interface {} )
	httpUrl,_ := urls["https"].(string)

	return httpUrl
}
