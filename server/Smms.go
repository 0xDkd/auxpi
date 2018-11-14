package server

import (
	"auxpi/auxpiAll"
	"bytes"
	"github.com/astaxie/beego"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"regexp"
)

type Smms struct {
}

//上传 SM 图床 返回图片 URL
func (this *Smms) UpLoatToSmms(img []byte, imgInfo string) string {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	content_type := w.FormDataContentType()
	pat := `filename="(.*)"`
	res := regexp.MustCompile(pat)
	name := res.FindAllStringSubmatch(imgInfo, -1)
	file, _ := w.CreateFormFile("smfile", name[0][1])
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
