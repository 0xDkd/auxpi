package models

import (
	"auxpi/bootstrap"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

//定义基础的 Model 实例
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

//初始化链接数据库
func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = bootstrap.SiteConfig.DbOption.DbType
	dbName = bootstrap.SiteConfig.DbOption.DbName
	user = bootstrap.SiteConfig.DbOption.DbUser
	password = bootstrap.SiteConfig.DbOption.DblPass
	host = bootstrap.SiteConfig.DbOption.DbHost
	tablePrefix = bootstrap.SiteConfig.DbOption.TablePrefix

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(db)
		log.Println(err)
	}

	//获取表名称
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName;
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
