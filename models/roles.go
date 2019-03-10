package models

import (
	"auxpi/auxpiAll"
)

type Role struct {
	Model

	Name        string `gorm:"unique;not null VARCHAR(191)" json:"name"`
	DisplayName string `gorm:"VARCHAR(191)" json:"display_name"`
	Description string `gorm:"VARCHAR(191)" json:"description"`

	Perms []*Permission `gorm:"many2many:role_perms;" json:"perms"`
}

const (
	ROLE_ADMIN        = 1
	ROLE_EDITER       = 2
	ROLE_NORMAL_USER  = 3
	ROLE_BLOCKED_USER = 4
)

func CreateRole(role auxpi.RoleJson) bool {
	err := db.Create(&Role{
		Name:        role.Name,
		DisplayName: role.DisplayName,
		Description: role.Description,
		Perms:       GetPermissionsByIds(role.PIDs),
	}).Error

	return modelsError(auxpi.ErrorToString(err))
}

func createAdminRole() {
	var role = new(Role)
	role.ID = ROLE_ADMIN
	role.Name = "admin"
	role.Description = `可以对用户\用户权限\图片\API\站点 进行管理`
	role.DisplayName = `管理员`
	var ids = make([]uint, 1)
	ids[0] = 1
	role.Perms = GetPermissionsByIds(ids)
	db.Create(role)

	role.ID = ROLE_NORMAL_USER
	role.Name = "normalUser"
	role.Description = `可以上传图片，管理自己的图片`
	role.DisplayName = `普通用户`
	ids = []uint{4, 5}
	role.Perms = GetPermissionsByIds(ids)
	db.Create(role)

	role.ID = ROLE_BLOCKED_USER
	role.Name = "blockUser"
	role.DisplayName = `封禁的用户`
	role.Description = "无法登录，无法进行任何操作,小黑屋用户"
	db.Create(role)
}

func MigrateRole() error {
	if db.HasTable(&Role{}) {
		err := db.DropTable(&Role{}).Error
		err = db.CreateTable(&Role{}).Error
		createAdminRole()
		return err
	} else {
		err := db.CreateTable(&Role{}).Error
		createAdminRole()
		return err
	}
}
