package server

import (
	"auxpi/auxpiAll"
	"auxpi/tools"
	"regexp"
)

type Qihoo struct {
}

func (this *Qihoo) UploadToQihoo(img []byte, imgInfo string, imgType string) string {
	url := "http://st.so.com/stu"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "upload",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	var re = regexp.MustCompile(`(?m)data-imgkey="(.*)"`)
	imgKey := re.FindAllStringSubmatch(data,-1)[0][1]
	url = "https://ps.ssl.qhmsg.com/"+imgKey
	return url
}
