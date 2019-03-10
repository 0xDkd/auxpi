package server

import (
	"auxpi/auxpiAll"
	"auxpi/tools"

	"github.com/astaxie/beego"
)

type VimCN struct {

}

func (this *VimCN) UploadToVimCN(img []byte, imgInfo string, imgType string) string {
	url := "https://img.vim-cn.com/"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "image",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	beego.Alert(data)
	return data
}