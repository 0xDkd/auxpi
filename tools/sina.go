package tools

import (
	"auxpi/bootstrap"
	"fmt"
	"hash/crc32"
)

func CheckPid(pid string,imgType string) string {
	sinaNumber := fmt.Sprint((crc32.ChecksumIEEE([]byte(pid)) & 3) + 1)
	//从配置文件中获取
	size := bootstrap.SiteConfig.SiteUploadWay.SinaAccount.DefultPicSize
	n := len(imgType)
	rs := []rune(imgType)
	suffix := string(rs[6:n])
	if suffix != "gif" {
		suffix = "jpg"
	}
	sinaUrl := "https://ws" + sinaNumber + ".sinaimg.cn/" + size + "/" + pid + "." + suffix
	return sinaUrl

}
