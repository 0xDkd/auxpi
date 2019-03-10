package models

import (
	"auxpi/auxpiAll"

	"github.com/astaxie/beego"
)

type User struct {
	Model

	Username string `gorm:"size:32;UNIQUE" json:"username" `
	Password string `gorm:"size:64" json:"password"`
	IsAdmin  bool   `json:"is_admin"`
	Email    string `gorm:"size:100;UNIQUE" json:"email"`
	Status   uint   `json:"status"`
	Token    string `gorm:"UNIQUE" json:"token"`
	Version  uint   `json:"version"`

	RoleID uint `gorm:"UNIQUE_INDEX;" json:"role_id"`
	Role   Role `json:"role"`

	Image []Image `json:"images" json:"image"`
}

func CheckAdminAuth(username, password string, ) bool {
	var auth User
	err := db.Model(&User{}).
		Where(User{Username: username, Password: password}).
		First(&auth).Error

	modelsError(auxpi.ErrorToString(err))

	if auth.ID > 0 && auth.RoleID == 1 {
		return true
	}
	return false
}

func CheckAndGetUser(email, password string) (user User, status bool) {
	err := db.Model(&User{}).Where(User{Email: email, Password: password}).
		First(&user).
		Error

	if user.ID > 0 && modelsError(auxpi.ErrorToString(err)) {
		status = true
		return
	}

	status = false

	return

}

func CanUserRegister(username, email string) bool {
	var user User
	//首先查库看看能否注册
	db.Model(&User{}).
		Where(&User{Username: username}).
		First(&user)
	if user.ID > 0 {
		return false
	}
	db.Model(&User{}).
		Where(&User{Email: email}).
		First(&user)
	if user.ID > 0 {
		return false
	}
	return true
}

func RegisterUser(username, email, password, token string) bool {
	//验证一波先
	if CanUserRegister(username, email) {

		err := db.Create(&User{
			Username: username,
			Email:    email,
			Password: password,
			Status:   1,
			RoleID:   ROLE_NORMAL_USER,
			Token:    token,
		}).Error

		return modelsError(auxpi.ErrorToString(err))
	}
	return false

}

func ActiveUser(token string) bool {
	err := db.Model(&User{}).
		Where(&User{Token: token}).
		Update("token", token).
		Error

	return modelsError(auxpi.ErrorToString(err))
}

func GetUserInfo(username string) (info User) {
	err := db.Preload("Role").
		Model(&User{}).
		Where("username = ?", username).
		First(&info).Error

	modelsError(auxpi.ErrorToString(err))
	return info
}

func GetUsers(page, size int) (u []User, c int) {
	err := db.Preload("Role").
		Model(&User{}).
		Count(&c).
		Offset(page).
		Limit(size).
		Find(&u).
		Error
	modelsError(auxpi.ErrorToString(err))
	return
}

func ResetUserPass(username, password string) bool {
	err := db.Model(&User{}).
		Where("username = ?", username).
		Update("password", password).Error

	return modelsError(auxpi.ErrorToString(err))
}

func ResetUserPassWithOld(id int, oldPass, newPass string) bool {
	var user = User{}
	err := db.Model(&User{}).Where(&User{
		Model:    Model{ID: id},
		Password: oldPass,
	}).First(&user).Error
	modelsError(auxpi.ErrorToString(err))
	if user.ID > 0 {
		err = db.Model(&User{}).
			Where("id = ?", id).
			Update("password", newPass).Error

		return modelsError(auxpi.ErrorToString(err))
	}

	return false

}

func ResetUserPassByEmail(email, password string) bool {
	err := db.Model(&User{}).
		Where("username = ?", email).
		Update("password", password).Error

	return modelsError(auxpi.ErrorToString(err))
}

func GetUserEmail(username string) (string, uint) {
	var email User
	err := db.Select("email").
		Where(&User{Username: username}).
		First(&email).Error

	modelsError(auxpi.ErrorToString(err))

	if email.Email != "" {
		return email.Email, email.Version
	}
	return "", 0
}

func GetUserRegisterSevenDayReport() (report []Report) {
	err := db.Model(&User{}).
		Select("COUNT(*) AS `number` , created_day AS `date`").
		Order("created_day ASC").
		Group("created_day").
		Limit(7).
		Scan(&report).Error

	modelsError(auxpi.ErrorToString(err))

	return
}

func GetUserImagesByUserName(username string, size, offset int) (user User, images []Image, count int) {
	err := db.
		Model(&User{}).
		Where(&User{Username: username}).
		Find(&user).
		Error
	modelsError(auxpi.ErrorToString(err))

	err = db.Model(&Image{}).
		Where(&Image{UserID: user.ID}).
		Count(&count).
		Offset(offset).
		Limit(size).
		Find(&images).
		Error

	modelsError(auxpi.ErrorToString(err))
	return
}

func GetUserInfoByToken(token string) (user User, status bool) {
	err := db.Model(&user).
		Where(&User{Token: token}).
		First(&user).
		Error

	modelsError(auxpi.ErrorToString(err))

	status = false
	if user.ID > 0 {
		status = true
	}

	return

}

func GetUserInfoByEmail(email string) (user User, status bool) {
	err := db.
		Model(&user).
		Where(&User{Email: email}).
		First(&user).Error
	modelsError(auxpi.ErrorToString(err))

	if user.ID > 0 {
		status = true
		return
	}

	status = false
	return

}

func GetUserInfoByID(id int) (user User, status bool) {
	err := db.Model(&user).
		Preload("Role").
		Where("id=?", id).
		First(&user).Error
	modelsError(auxpi.ErrorToString(err))

	if user.ID > 0 {
		status = true
		return
	}

	status = false
	return
}

func DeleteUserById(id int) bool {
	//任何情况都不能删除管理员
	user, status := GetUserInfoByID(id)
	if !status {
		return false
	}

	if user.RoleID == 1 {
		return false
	}

	err := db.Model(&User{}).
		Where("id=?", id).
		Delete(&User{}).Error

	return modelsError(auxpi.ErrorToString(err))
}

func MigrateUsers() error {
	if db.HasTable(&User{}) {
		err := db.DropTable(&User{}).Error
		err = db.CreateTable(&User{}).Error
		return err
	} else {
		err := db.CreateTable(&User{}).Error
		return err
	}
}

func RegisterAdmin(u, pass, token, email string) {
	if beego.BConfig.RunMode == "dev" {
		user := &User{}
		user.ID = 1
		user.Password = pass
		user.Username = u
		user.Token = token
		user.Status = 1
		user.IsAdmin = true
		user.Email = email
		user.Version = 1
		user.RoleID = 1
		db.Create(user)
	}
}
