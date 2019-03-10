package server

import (
	"auxpi/auxpiAll"
	"auxpi/tools"

	"github.com/astaxie/beego"
)

//upload.cc
type CC struct {
}

func (this *CC) UploadToCC(img []byte, imgInfo string, imgType string) (string, string) {
	url := "https://upload.cc/image_upload"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "uploaded_file[]",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	j := auxpi.CCResp{}
	j.UnmarshalJSON([]byte(data))
	beego.Alert(j)
	mj, _ := j.SuccessImage[0].(map[string]interface{})
	smj, _ := mj["url"].(string)
	del, _ := mj["delete"].(string)

	url = "https://upload.cc/" + smj

	deleteJson := `[{"path":"` + smj + `",key":"` + del + `"}]`
	return url, deleteJson
}
