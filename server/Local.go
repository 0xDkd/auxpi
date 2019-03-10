package server

import (
	"auxpi/bootstrap"
	"bufio"
	"os"
	"time"

	"github.com/astaxie/beego"
)

type Local struct {
}

//获取本地图片链接
func (this *Local) UpLoadToLocal(name string, fileContent []byte) (string, string, string, string) {
	//
	if !bootstrap.SiteConfig.SiteUploadWay.LocalStore.Open {
		return "", "", "", ""
	}

	host := &bootstrap.SiteConfig.SiteUrl
	storeLocation := &bootstrap.SiteConfig.SiteUploadWay.LocalStore.StorageLocation
	softLink := &bootstrap.SiteConfig.SiteUploadWay.LocalStore.Link
	//修正URL
	bootstrap.FormatUrl(softLink)
	bootstrap.FormatUrl(host)
	bootstrap.FormatUrl(storeLocation)

	suffix := this.storeImage(*storeLocation, name, fileContent)
	url := *host + *softLink + suffix
	beego.Alert(url)
	backup := *host + "backup/" + suffix
	str := `ZXCVBNMASDFGHJKLQWERTYUIOPzxcvbnmasdfghjklqwertyuiop1234567890`
	randomStr := bootstrap.GetRandomString(16, str)
	return url, backup, randomStr, *storeLocation + suffix
}

//储存图片
func (this *Local) storeImage(path string, n string, fileContent []byte) string {
	nowTime := beego.Date(time.Now(), "Y/m/d/")
	str := `ZXCVBNMASDFGHJKLQWERTYUIOPzxcvbnmasdfghjklqwertyuiop1234567890`
	suffix := bootstrap.GetRandomString(16, str) + "." + this.getImageSuffix(n)
	dir := path + nowTime
	file := dir + suffix
	bootstrap.CheckPath(dir)
	var f *os.File
	f, err := os.Create(file)
	if err != nil {
		beego.Alert("File Create Error:", err)
	}
	w := bufio.NewWriter(f)
	_, err = w.Write(fileContent)
	if err != nil {
		beego.Alert("File Create Error:", err)
	}
	w.Flush()
	f.Close()

	return nowTime + suffix
}

//获取图片后缀
func (this *Local) getImageSuffix(name string) (suffix string) {
	n := len(name)
	rs := []rune(name)
	suffix = string(rs[n-3 : n])
	beego.Alert(suffix)
	if suffix == "peg" {
		suffix = "jpeg"
	}

	return suffix
}
