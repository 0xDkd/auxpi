package tools

import (
	"regexp"
)

//通过 MimeType 信息获取文件名称
func GetFileNameByMimeType(info string) string {
	//beego.Alert(info)
	pat := `filename="(.*)"`
	res := regexp.MustCompile(pat)
	name := res.FindAllStringSubmatch(info, -1)
	return name[0][1]
}