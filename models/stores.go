package models

import (
	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/bootstrap"
)

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
	Status bool `json:"status"` //是否开启此图床
	Weight int  `json:"weight"` // v2 分发使用
	Rank   int  `json:"rank"`   //首页排名使用
	//是否需要登录才能使用 0->不需要 1->需要并且开放给所有登录用户 2->仅仅给管理员使用
	Auth uint `json:"auth"`

	//包含多个
	Images []Image
}

func GetStores() (stores []Store) {
	//缓存处理
	if value, status := getCacheStores("all", false); status {
		return value
	}
	err := db.Model(&Store{}).
		Order("`rank` ASC").
		Find(&stores).
		Error
	if err != nil {
		setCacheStores("all", &stores)
	}
	modelsError(auxpi.ErrorToString(err))
	return
}

func GetActiveStore() (stores []Store) {
	if value, status := getCacheStores("active", false); status {
		return value
	}

	err := db.Model(&Store{}).
		Where("status=?", true).
		Order("`rank` ASC").
		Find(&stores).Error
	if err != nil {
		setCacheStores("active", &stores)
	}
	modelsError(auxpi.ErrorToString(err))
	return
}

func GetCloseStore() (stores []Store) {
	err := db.Model(&Store{}).
		Where("status=?", false).
		Order("`rank` DESC").
		Find(&stores).Error
	modelsError(auxpi.ErrorToString(err))
	return
}

func GetStoreInfoByRouter(router string) (store Store) {
	err := db.Model(&Store{}).
		Where("router=?", router).
		First(&store).Error
	modelsError(auxpi.ErrorToString(err))
	return
}

func GetStoreInfoByApi(api string) (store Store) {
	err := db.Model(&Store{}).
		Where("api=?", api).
		First(&store).Error
	modelsError(auxpi.ErrorToString(err))
	return
}

func DisableStores(receive auxpi.MenuReceive) bool {
	var l = len(receive.Disable)
	ids := make([]int, l)
	for k, v := range receive.Disable {
		ids[k] = v.ID
	}
	err := db.Model(&Store{}).
		Where("id IN (?)", ids).Update("status", false).
		Error
	if err != nil {
		//刷新缓存
		deleteCacheAllStores()
		as := GetStores()
		ats := GetActiveStore()
		setCacheStores("active", &ats)
		setCacheStores("all", &as)
	}
	return modelsError(auxpi.ErrorToString(err))
}

func EnableStores(receive auxpi.MenuReceive) bool {
	var l = len(receive.Enable)
	ids := make([]int, l)
	for k, v := range receive.Enable {
		ids[k] = v.ID
	}
	err := db.Model(&Store{}).
		Where("id IN (?)", ids).Update("status", true).
		Error
	if err != nil {
		//刷新缓存
		deleteCacheAllStores()
		as := GetStores()
		ats := GetActiveStore()
		setCacheStores("active", &ats)
		setCacheStores("all", &as)
	}
	return modelsError(auxpi.ErrorToString(err))
}

func RankStores(receive auxpi.MenuReceive) bool {
	for k, v := range receive.Enable {
		err := db.Model(&Store{}).
			Where("id=?", v.ID).Update("rank", k+1).
			Error
		if err != nil {
			return modelsError(auxpi.ErrorToString(err))
		}
		modelsError(auxpi.ErrorToString(err))
	}
	return true
}

func UpdateStore(store Store) bool {
	err := db.Save(&store).Error
	if err != nil {
		deleteCacheAllStores()
		as := GetStores()
		ats := GetActiveStore()
		setCacheStores("active", &ats)
		setCacheStores("all", &as)
	}
	return modelsError(auxpi.ErrorToString(err))
}

//TODO: callback auto cache
func getCacheStores(key string, single bool) ([]Store, bool) {
	//缓存
	s := bootstrap.Cache.Get("store_" + key)
	list := []Store{}
	if s != nil {
		list, _ = s.([]Store)
		return list, true
	}
	return list, false
}

func setCacheStores(key string, value *[]Store) bool {
	err := bootstrap.Cache.Put("store_"+key, value, 3600)
	return modelsError(auxpi.ErrorToString(err))
}

func deleteCacheAllStores() {
	bootstrap.Cache.Delete("store_all")
	bootstrap.Cache.Delete("store_active")
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

//初始化储存方案
func initStores() {
	s := &Store{}
	s.Name = "搜狗图床"
	s.Icon = "sougou"
	s.Url = "http://pic.sogou.com/pic/upload_pic.jsp"
	s.Desc = "搜狗图床,有防盗链,会定期清除违规图片,建议仅作为备份"
	s.Color = "orange"
	s.Api = "sougou"
	s.Router = "SouGou"
	s.Status = true
	s.Auth = 0
	s.ID = 1
	s.Weight = 16
	db.Create(s)

	s.Name = "微博图床"
	s.ID = 2
	s.Icon = "weibo"
	s.Url = "http://picupload.service.weibo.com/interface/pic_upload.php" +
		"?ori=1&mime=image%2Fjpeg&data=base64&url=0&markpos=1&logo=&nick=0&marks=1&app=miniblog"
	s.Desc = "新浪图床需要登录,对于违规图片有限制,但是不会清理图片"
	s.Color = "red"
	s.Api = "sina"
	s.Router = "Sina"
	s.Status = true
	s.Auth = 0
	s.Weight = 1
	db.Create(s)

	s.Name = "SMMS 图床"
	s.Icon = "sm"
	s.Url = "https://sm.ms/api/upload"
	s.Desc = "SMMS 兽兽家的图床，速度很不错，相当稳定~"
	s.Color = "blue"
	s.Api = "smms"
	s.Router = "Smms"
	s.Status = true
	s.Auth = 0
	s.ID = 3
	s.Weight = 2
	db.Create(s)

	s.Name = "CC 图床"
	s.Icon = "cc"
	s.Url = "https://upload.cc/image_upload"
	s.Desc = "CC 图床，服务器在香港，速度还可以"
	s.Color = "green"
	s.Api = "cc"
	s.Router = "cc"
	s.Status = true
	s.Auth = 0
	s.ID = 4
	s.Weight = 3
	db.Create(s)

	s.Name = "Flickr 图床"
	s.Icon = "flickr"
	s.Url = ""
	s.Desc = "Flickr 图床是雅虎旗下的，图片可以在国内访问速度不错，但是 api 只能国外访问"
	s.Color = "pink darken-1"
	s.Api = "flickr"
	s.Router = "Flickr"
	s.Status = true
	s.Auth = 0
	s.ID = 5
	s.Weight = 4
	db.Create(s)

	//接替百度的位置
	s.Name = "Imgur"
	s.Icon = "imgur"
	s.Url = ""
	s.Desc = "imgur 图床会永久储存，可以使用节点进行代理"
	s.Color = "purple"
	s.Api = "imgur"
	s.Router = "Imgur"
	s.Status = true
	s.Auth = 0
	s.ID = 6
	s.Weight = 15
	db.Create(s)

	//接替 360 的位置，360 会在以后加上
	s.Name = "Prnt"
	s.Icon = "prnt"
	s.Url = ""
	s.Desc = "本地图床没有任何限制，但是会占用服务器的储存空间"
	s.Color = "black"
	s.Api = "prnt"
	s.Router = "Prnt"
	s.Status = true
	s.Auth = 0
	s.ID = 7
	s.Weight = 14
	db.Create(s)

	//s.Name = "360 图床"
	//s.Icon = "logo-"
	//s.Url = ""
	//s.Desc = "360 图床,已经绕过防盗链,会定期清理"
	//s.Color = "green"
	//s.Api = "Qihoo"
	//s.Router = "/360"
	//s.Status = true
	//s.Auth = 0
	//s.ID = 7
	//s.Weight = 999
	//db.Create(s)

	s.Name = "网易图床"
	s.Icon = "wangyi"
	s.Url = ""
	s.Desc = "网易严选 CDN,暂无防盗链,目前发现不会清理"
	s.Color = "red"
	s.Api = "neteasy"
	s.Router = "NetEasy"
	s.Status = true
	s.Auth = 0
	s.ID = 8
	s.Weight = 5
	db.Create(s)

	s.Name = "京东 图床"
	s.Icon = "jingdong"
	s.Url = ""
	s.Desc = "京东图床，不允许上传过小的图片，其它正常"
	s.Color = "orange-800"
	s.Api = "jd"
	s.Router = "Jd"
	s.Status = true
	s.Auth = 0
	s.ID = 9
	s.Weight = 6
	db.Create(s)

	s.Name = "掘金 图床"
	s.Icon = "juejin"
	s.Url = ""
	s.Desc = "掘金图床,暂无防盗链，一切正常"
	s.Color = "blue"
	s.Api = "juejin"
	s.Router = "JueJin"
	s.Status = true
	s.Auth = 0
	s.ID = 10
	s.Weight = 7
	db.Create(s)

	s.Name = "阿里 图床"
	s.Icon = "ali"
	s.Url = ""
	s.Desc = "阿里图床,推荐使用"
	s.Color = "orange"
	s.Api = "ali"
	s.Router = "Ali"
	s.Status = true
	s.Auth = 0
	s.ID = 11
	s.Weight = 8
	db.Create(s)

	s.Name = "本地图床"
	s.Icon = "auxpi"
	s.Url = ""
	s.Desc = "本地图床没有任何限制，但是会占用服务器的储存空间"
	s.Color = "purple"
	s.Api = "local"
	s.Router = "Local"
	s.Status = true
	s.Auth = 0
	s.ID = 12
	s.Weight = 9
	db.Create(s)

	s.Name = "小米图床"
	s.Icon = "xiaomi"
	s.Url = ""
	s.Desc = "小米图床不是很稳定，返回的 json 会改变"
	s.Color = "orange"
	s.Api = "xiaomi"
	s.Router = "XiaoMi"
	s.Status = true
	s.Auth = 0
	s.ID = 14
	s.Weight = 10
	db.Create(s)

	s.Name = "苏宁图床"
	s.Icon = "suning"
	s.Url = ""
	s.Desc = "建议观察后使用"
	s.Color = "orange"
	s.Api = "suning"
	s.Router = "SuNing"
	s.Status = true
	s.Auth = 0
	s.ID = 13
	s.Weight = 11
	db.Create(s)

	s.Name = "Vim图床"
	s.Icon = "vim"
	s.Url = ""
	s.Desc = "国外图床，非常稳定，cf 的 CDN"
	s.Color = "green"
	s.Api = "vim"
	s.Router = "Vim"
	s.Status = true
	s.Auth = 0
	s.ID = 15
	s.Weight = 12
	db.Create(s)

	s.Name = "OOXX"
	s.Icon = "ooxx"
	s.Url = ""
	s.Desc = "国外图床，稳定性未知，cf 提供 CDN"
	s.Color = "red"
	s.Api = "ooxx"
	s.Router = "Local"
	s.Status = true
	s.Auth = 0
	s.ID = 16
	s.Weight = 13
	db.Create(s)

	s.Name = "搜狐"
	s.Icon = "souhu"
	s.Url = ""
	s.Desc = "搜狐畅言图床，目前看来还不错"
	s.Color = "orange"
	s.Api = "souhu"
	s.Router = "SouHu"
	s.Status = true
	s.Auth = 0
	s.ID = 17
	s.Weight = 13
	db.Create(s)

	s.Name = "Github"
	s.Icon = "github"
	s.Url = ""
	s.Desc = "github 的仓库为我们提供图床，可以当做根节点使用"
	s.Color = "black"
	s.Api = "github"
	s.Router = "Github"
	s.Status = true
	s.Auth = 0
	s.ID = 18
	s.Weight = 13
	db.Create(s)

	s.Name = "头条"
	s.Icon = "toutiao"
	s.Url = ""
	s.Desc = "头条 CDN，稳定性位置"
	s.Color = "red"
	s.Api = "toutiao"
	s.Router = "TouTiao"
	s.Status = true
	s.Auth = 0
	s.ID = 19
	s.Weight = 13
	db.Create(s)

	s.Name = "Gitee"
	s.Icon = "gitee"
	s.Url = ""
	s.Desc = "gitee 很稳定，希望大家不要滥用"
	s.Color = "red"
	s.Api = "gitee"
	s.Router = "Gitee"
	s.Status = true
	s.Auth = 0
	s.ID = 20
	s.Weight = 13
	db.Create(s)

}
