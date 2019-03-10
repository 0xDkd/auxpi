package models

type Store struct {
	Model

	Name   string `gorm:"size:32" json:"name"`
	Desc   string `json:"desc"`
	Url    string `json:"url"`
	Icon   string `gorm:"size:32" json:"icon"`
	Color  string `gorm:"size:32" json:"color"`
	Api    string `gorm:"size:32" json:"api"`
	Router string `gorm:"size:32" json:"router"`
	//是否开启
	Status bool `json:"Status"`
	//是否需要登录才能使用 0->不需要 1->需要并且开放给所有登录用户 2->仅仅给管理员使用
	Auth uint `json:"Auth"`

	//包含多个
	Images []Image
}

func MigrateStores() error {
	if db.HasTable(&Store{}) {
		err := db.DropTable(&Store{}).Error
		err = db.CreateTable(&Store{}).Error
		initStores()
		return err
	} else {
		err := db.CreateTable(&Store{}).Error
		initStores()
		return err
	}

}


func GetStores() (stores []Store) {
	db.Model(&Store{}).Find(&stores)
	return
}

//初始化储存方案
func initStores() {
	s := &Store{}
	s.Name = "搜狗图床"
	s.Icon = "sougou"
	s.Url = "http://pic.sogou.com/pic/upload_pic.jsp"
	s.Desc = "搜狗图床,有防盗链,会定期清除违规图片,建议仅作为备份"
	s.Color = "orange"
	s.Api = "SouGou"
	s.Router = "/"
	s.Status = true
	s.Auth = 0
	s.ID = 1
	db.Create(s)

	s.Name = "微博图床"
	s.ID = 2
	s.Icon = "xinlang"
	s.Url = "http://picupload.service.weibo.com/interface/pic_upload.php" +
		"?ori=1&mime=image%2Fjpeg&data=base64&url=0&markpos=1&logo=&nick=0&marks=1&app=miniblog"
	s.Desc = "新浪图床需要登录,对于违规图片有限制,但是不会清理图片"
	s.Color = "red"
	s.Api = "Sina"
	s.Router = "/Sina"
	s.Status = true
	s.Auth = 0
	db.Create(s)

	s.Name = "SMMS 图床"
	s.Icon = "sm"
	s.Url = "https://sm.ms/api/upload"
	s.Desc = "SMMS 兽兽家的图床，速度很不错，相当稳定~"
	s.Color = "blue"
	s.Api = "Smms"
	s.Router = "/Smms"
	s.Status = true
	s.Auth = 0
	s.ID = 3
	db.Create(s)

	s.Name = "CC 图床"
	s.Icon = "cc"
	s.Url = "https://upload.cc/image_upload"
	s.Desc = "CC 图床，服务器在香港，速度还可以"
	s.Color = "green"
	s.Api = "CC"
	s.Router = "/cc"
	s.Status = true
	s.Auth = 0
	s.ID = 4
	db.Create(s)

	s.Name = "Flickr 图床"
	s.Icon = "flickr"
	s.Url = ""
	s.Desc = "Flickr 图床是雅虎旗下的，图片可以在国内访问速度不错，但是 api 只能国外访问"
	s.Color = "pink darken-1"
	s.Api = "Flickr"
	s.Router = "/Flickr"
	s.Status = true
	s.Auth = 0
	s.ID = 5
	db.Create(s)

	s.Name = "Baidu 图床"
	s.Icon = "baidu1"
	s.Url = ""
	s.Desc = "百度图床,已经绕过防盗链,会定期清理"
	s.Color = "blue-300"
	s.Api = "Baidu"
	s.Router = "/Baidu"
	s.Status = true
	s.Auth = 0
	s.ID = 6
	db.Create(s)

	s.Name = "360 图床"
	s.Icon = "logo-"
	s.Url = ""
	s.Desc = "360 图床,已经绕过防盗链,会定期清理"
	s.Color = "green"
	s.Api = "Qihoo"
	s.Router = "/360"
	s.Status = true
	s.Auth = 0
	s.ID = 7
	db.Create(s)

	s.Name = "网易图床"
	s.Icon = "wangyi"
	s.Url = ""
	s.Desc = "百度图床,已经绕过防盗链,会定期清理"
	s.Color = "red"
	s.Api = "NetEasy"
	s.Router = "/NetEasy"
	s.Status = true
	s.Auth = 0
	s.ID = 8
	db.Create(s)

	s.Name = "京东 图床"
	s.Icon = "jingdong-"
	s.Url = ""
	s.Desc = "京东图床,已经绕过防盗链,会定期清理"
	s.Color = "orange-800"
	s.Api = "Qihoo"
	s.Router = "/Jd"
	s.Status = true
	s.Auth = 0
	s.ID = 9
	db.Create(s)

	s.Name = "掘金 图床"
	s.Icon = "juejin"
	s.Url = ""
	s.Desc = "京东图床,已经绕过防盗链,会定期清理"
	s.Color = "blue"
	s.Api = "JueJin"
	s.Router = "/JueJin"
	s.Status = true
	s.Auth = 0
	s.ID = 10
	db.Create(s)

	s.Name = "阿里 图床"
	s.Icon = "ali"
	s.Url = ""
	s.Desc = "阿里图床,已经绕过防盗链,会定期清理"
	s.Color = "orange"
	s.Api = "Ali"
	s.Router = "/Ali"
	s.Status = true
	s.Auth = 0
	s.ID = 11
	db.Create(s)

	s.Name = "本地图床"
	s.Icon = "auxly"
	s.Url = ""
	s.Desc = "本地图床没有任何限制，但是会占用服务器的储存空间"
	s.Color = "purple"
	s.Api = "Local"
	s.Router = "/Local"
	s.Status = true
	s.Auth = 0
	s.ID = 12
	db.Create(s)

}
