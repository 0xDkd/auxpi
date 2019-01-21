package server

import (
	"auxpi/auxpiAll"
	"auxpi/tools"
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/astaxie/beego"
)

type Smms struct {
}

//上传 SM 图床 返回图片 URL
func (this *Smms) UpLoatToSmms(img []byte, imgInfo string) string {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	content_type := w.FormDataContentType()
	beego.Alert(imgInfo)
	name := tools.GetFileNameByMimeType(imgInfo)
	file, _ := w.CreateFormFile("smfile", name)
	file.Write(img)
	w.Close()
	req, _ := http.NewRequest("POST", "https://sm.ms/api/upload", body)
	req.Header.Set("Content-Type", content_type)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	sm := auxpi.SmResponse{}
	beego.Alert(string(data))
	sm.UnmarshalJSON([]byte(string(data)))
	return string(sm.Data.Url)
}
