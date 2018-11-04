package models

type SiteDefault struct {
	Url string
	Api string
}

type SiteConfig struct {
	Name string
	//最大上传的图片个数
	UploadMaxNumber string
	//最大图片规格 MB
	UpLoadMaxSize int
	//图床储存的一些配置
	UploadWay UploadConfig
}

type UploadConfig struct {
	//是否开启本地上传
	LocalStore bool
	//选择上传的图床 Sina | SouGou | All
	StorePic string
	//
}
