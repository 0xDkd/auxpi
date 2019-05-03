package models

import (
	"github.com/auxpi/auxpiAll"

	"github.com/astaxie/beego/logs"
)

type Permission struct {
	Model

	Name        string `gorm:"unique;not null VARCHAR(191)" json:"name"`
	DisplayName string `gorm:"VARCHAR(191)" json:"display_name"`
	Description string `gorm:"VARCHAR(191)" json:"description"`
	//匹配的路由黑名单
	//RouterBlock string
}

func GetPermissionsByIds(ids []uint) (ps []*Permission) {
	err := db.Model(&Permission{}).Where(ids).Find(&ps).Error
	if err != nil {
		logs.Alert("[DataBase Error] :", err)
		return
	}
	return
}

func GetPermissionsById(id uint) (p *Permission) {
	err := db.Model(&Permission{}).Where("id=?", id).First(&p).Error

	modelsError(auxpi.ErrorToString(err))

	return
}

func createAdminPermissions() {
	var p = new(Permission)
	p.ID = 1
	p.Name = "All"
	p.Description = "具有所有权限"
	p.DisplayName = "所有权限"
	db.Create(p)

	p.ID = 2
	p.Name = "edit_user"
	p.Description = "edit_user"
	p.DisplayName = "修改用户"
	db.Create(p)

	p.ID = 3
	p.Name = "edit_image"
	p.Description = "edit_image"
	p.DisplayName = "修改图片"
	db.Create(p)

	p.ID = 4
	p.Name = "upload_image"
	p.Description = "upload_image"
	p.DisplayName = "上传图片"
	db.Create(p)

	p.ID = 5
	p.Name = "delete_image"
	p.Description = "delete_image"
	p.DisplayName = "删除图片"
	db.Create(p)

}

func MigratePermissions() error {
	if db.HasTable(&Permission{}) {
		err := db.DropTable(&Permission{}).Error
		err = db.CreateTable(&Permission{}).Error
		createAdminPermissions()
		return err
	} else {
		err := db.CreateTable(&Permission{}).Error
		createAdminPermissions()
		return err
	}
}
