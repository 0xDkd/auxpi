package server

import (
	"auxpi/auxpiAll"
	"auxpi/tools"
	"math/rand"
	"regexp"
	"strconv"

	"github.com/astaxie/beego"
)

type Jd struct {
}

func (this *Jd) UploadToJd(img []byte, imgInfo string, imgType string) string {
	url := "https://search.jd.com/image?op=upload"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "file",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)

	var re = regexp.MustCompile(`(?m)\("(.*")\)`)
	imgFix := re.FindAllStringSubmatch(data, -1)[0][1]
	url = "https://img" + strconv.Itoa(rand.Intn(3)+11) + ".360buyimg.com/img/" + imgFix
	beego.Alert(url)
	return imgFix
}
