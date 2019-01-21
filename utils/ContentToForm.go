package utils

import (
	"auxpi/auxpiAll"
	"io/ioutil"
	"net/http"
	"strings"
)

func FormPost(file *auxpi.FormFile, url string, header map[string]string) string {
	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"" + file.Key + "\"; filename=\"" + file.Name + "\"\r\nContent-Type: " + file.Type + "\r\n\r\n\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	if len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
