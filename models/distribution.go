package models

import (
	"github.com/auxpi/auxpiAll"
)

type Distribution struct {
	Model

	Hash        string `gorm:"INDEX;" json:"hash"`
	Url         string `json:"url"`
	IP          string `json:"ip"`
	Short       string `json:"short"`
	DispatchUrl string `json:"dispatch_url"`

	StoreID int   `json:"store_id"`
	Store   Store `json:"store"`

	UserID int  `json:"user_id"`
	User   User `json:"user"`

	RootID  int    `json:"root_id"`
	Root    Store  `gorm:"ForeignKey:RootID";json:"root"`
	RootUrl string `json:"root_url"`
	Number  int    `json:"number"`
}

func CreateDistribution(distribution Distribution) bool {
	err := db.Create(&Distribution{
		Hash:    distribution.Hash,
		Url:     distribution.Url,
		IP:      distribution.IP,
		StoreID: distribution.StoreID,
		RootID:  distribution.RootID,
		Short:   distribution.Short,
		RootUrl: distribution.RootUrl,
		Number:  0,
	}).Error
	return modelsError(auxpi.ErrorToString(err))
}

func ChangeUrl(hash string, url string, storeId int) bool {
	var d = &Distribution{}
	err := db.Model(&Distribution{}).
		Where("hash=?", hash).
		First(d).Error

	if d.ID <= 0 {
		return false
	}

	d.Number++
	d.Url = url
	d.StoreID = storeId
	err = db.Save(d).Error

	return modelsError(auxpi.ErrorToString(err))
}

//图片回源处理
func BackToSource(hash string) bool {
	var d = &Distribution{}
	err := db.Model(&Distribution{}).
		Where("hash=?", hash).
		First(d).Error

	if d.ID <= 0 {
		return false
	}

	d.Number = -1
	d.Url = d.RootUrl
	d.StoreID = d.RootID
	err = db.Save(d).Error

	return modelsError(auxpi.ErrorToString(err))
}

func GetUrlByHash(hash string) (dis Distribution, s bool) {
	err := db.Model(&Distribution{}).
		Where("hash=?", hash).
		First(&dis).
		Error
	if dis.ID <= 0 {
		s = false
		return
	}
	s = modelsError(auxpi.ErrorToString(err))
	return
}

//分页获取列表
func GetDistributionList(pageNum int, pageSize int, maps interface{}, sort string) (distribution []Distribution, count int) {
	db.Preload("User").
		Preload("Store").
		Preload("Root").
		Model(&Distribution{}).
		Where(maps).
		Order("`ID` " + sort).
		Count(&count).
		Offset(pageNum).
		Limit(pageSize).
		Find(&distribution)
	return
}

func MigrateDistribution() error {
	if db.HasTable(&Distribution{}) {
		err := db.DropTable(&Distribution{}).Error
		err = db.CreateTable(&Distribution{}).Error
		return err
	} else {
		err := db.CreateTable(&Distribution{}).Error
		return err
	}
}
