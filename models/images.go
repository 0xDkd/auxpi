package models

import (
	"auxpi/auxpiAll"
	"auxpi/bootstrap"
	"os"

	"github.com/astaxie/beego"
)

type Image struct {
	Model

	Name string `json:"name"`
	Link string `gorm:"UNIQUE" json:"link" `

	//属于
	StoreID int `gorm:"UNIQUE_INDEX" json:"store_id"`
	UserID  int `gorm:"UNIQUE_INDEX" json:"user_id"`

	//本地，CC ,SM 可以使用
	Delete string `gorm:"size:255;UNIQUE" json:"delete"`

	//本地储存位置 仅本地可用
	Path string `gorm:"UNIQUE" json:"path"`

	//获取外键
	User  User  `json:"user"`
	Store Store `json:"store"`
}

//首页统计各个图床占比返回
type AllImageStore struct {
	Sid   uint   `json:"-"`
	Name  string `json:"name"`
	Total int    `json:"value"`
}

func AddImage(image *auxpi.ImageJson) bool {

	db.Create(&Image{
		Name:    image.Name,
		Link:    image.Url,
		UserID:  image.UserID,
		StoreID: image.StoreID,
		Delete:  image.Delete,
		Path:    image.Path,
	})

	return true
}

func GetStoreNameByImageID(id int) string {
	var s Store
	db.Select("name").Where("id=?", id).First(&s)

	return s.Name
}

func GetImages(pageNum int, pageSize int, maps interface{}) (images []Image, count int) {
	db.Preload("User").
		Preload("Store").
		Model(&Image{}).
		Where(maps).
		Count(&count).
		Offset(pageNum).Limit(pageSize).
		Find(&images)

	return
}

func GetImagesByUserId(pageNum int, pageSize int, id uint) (image []Image, count int) {
	err := db.Preload("User").
		Preload("Store").
		Model(&Image{}).
		Where("user_id=?", id).
		Count(&count).
		Offset(pageNum).Limit(pageSize).
		Find(&image).Error

	modelsError(auxpi.ErrorToString(err))

	return
}

func GetAllImagesStoreNumber() (result []AllImageStore) {
	err := db.Model(&Image{}).
		Select("`store_id` AS `sid`, store.`name` as `name` , count(*) as `total`").
		Joins("left join `" +
			bootstrap.SiteConfig.DbOption.TablePrefix +
			"store` AS `store` on store.`id`=" +
			bootstrap.SiteConfig.DbOption.TablePrefix +
			"image.`store_id` ").
		Group("store_id").
		Scan(&result).Error

	modelsError(auxpi.ErrorToString(err))

	return

}

func GetAllImagesReport() (report []Report) {
	err := db.Model(&Image{}).
		Select("COUNT(*) AS `number` , created_day AS `date`").
		Order("`created_day`").
		Group("`created_day`").
		Limit(7).
		Scan(&report).Error

	modelsError(auxpi.ErrorToString(err))

	return
}

func GetLocalImageReport() (report []Report) {
	err := db.Model(&Image{}).
		Select("COUNT(*) AS `number` , created_day AS `date`").
		Where("store_id=?", 12).
		Order("`created_day` ASC ").
		Group("`created_day`").
		Limit(7).
		Scan(&report).Error

	modelsError(auxpi.ErrorToString(err))

	return
}

func DelImageByPath(ids []int) (images []Image) {
	err := db.Model(&Image{}).Select("path").Where(ids).Find(&images).Error
	if err != nil {
		beego.Alert("[Delete Image Record From DataBase Error]:", err)
		return
	}
	//删除图片
	for _, value := range images {
		if value.Path != "" {
			err := os.Remove(value.Path)
			beego.Alert("remove :? )")
			if err != nil {
				AddLog("IMAGE_DELETE", auxpi.ErrorToString(err), "SYSTEM", "ERROR")
				beego.Alert("[Delete Images localStore Error]:", err)
			}
		}
	}

	return
}

func DelImages(ids []int) error {
	DelImageByPath(ids)
	return db.
		Where(ids).
		Delete(&Image{}).
		Error

}

func MigrateImages() error {
	if db.HasTable(&Image{}) {
		err :=db.DropTable(&Image{}).Error
		err = db.CreateTable(&Image{}).Error
		return err
	} else {
		err :=db.CreateTable(&Image{}).Error
		return err
	}
}
