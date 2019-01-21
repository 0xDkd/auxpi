package tools

import (
	"auxpi/auxpiAll"
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

//自定义 post 表单上传
func FormPost(fileContent *auxpi.FormFile, url string, header map[string]string) string {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	contentType := w.FormDataContentType()

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fileContent.Key), escapeQuotes(fileContent.Name)))
	h.Set("Content-Type", "application/"+fileContent.Type)

	file, _ := w.CreatePart(h)
	file.Write(fileContent.Value)

	w.Close()
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", contentType)

	if len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	return string(data)
}


func AliFormPost(fileContent *auxpi.FormFile, url string) string {
	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"file\"; filename=\"" + fileContent.Name + "\"\r\nContent-Type: " + fileContent.Type + "\r\n\r\n" + string(fileContent.Value) + "\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"scene\"\r\n\r\naeMessageCenterV2ImageRule\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"name\"\r\n\r\n" + fileContent.Name + "\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--\r\n")
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
