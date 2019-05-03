package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/auxpi/bootstrap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

//定义基础的 Model 实例
type Model struct {
	ID         int    `json:"id" gorm:"primary_key" `
	CreatedOn  int    `json:"created_on" `
	ModifiedOn int    `json:"modified_on" `
	DeletedOn  int    `json:"deleted_on" `
	CreatedDay string `json:"created_day" gorm:"size:32"`
}

//首页返回公用 Report
type Report struct {
	Date   string `json:"date"`
	Number int    `json:"number"`
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
	password = bootstrap.SiteConfig.DbOption.DbPass
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
		return tablePrefix + defaultTableName
	}

	//注册回调
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	if beego.BConfig.RunMode == "dev" {
		db.LogMode(true)
	}

}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		nowDay := beego.Date(time.Now(), "Y/m/d")
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime.Unix())
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime.Unix())
			}
		}

		if createDay, ok := scope.FieldByName("CreatedDay"); ok {
			if createDay.IsBlank {
				createDay.Set(nowDay)

			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifyTime` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

func BatchInsert(db *gorm.DB, objArr []SyncImage) error {
	// If there is no data, nothing to do.
	if len(objArr) == 0 {
		return nil
	}

	mainObj := objArr[0]
	mainScope := db.NewScope(mainObj)
	mainFields := mainScope.Fields()
	quoted := make([]string, 0, len(mainFields))
	for i := range mainFields {
		// If primary key has blank value (0 for int, "" for string, nil for interface ...), skip it.
		// If field is ignore field, skip it.
		if (mainFields[i].IsPrimaryKey && mainFields[i].IsBlank) || (mainFields[i].IsIgnored) {
			continue
		}
		quoted = append(quoted, mainScope.Quote(mainFields[i].DBName))
	}

	placeholdersArr := make([]string, 0, len(objArr))

	for _, obj := range objArr {
		scope := db.NewScope(obj)
		fields := scope.Fields()
		placeholders := make([]string, 0, len(fields))
		for i := range fields {
			if (fields[i].IsPrimaryKey && fields[i].IsBlank) || (fields[i].IsIgnored) {
				continue
			}
			placeholders = append(placeholders, scope.AddToVars(fields[i].Field.Interface()))
		}
		placeholdersStr := "(" + strings.Join(placeholders, ", ") + ")"
		placeholdersArr = append(placeholdersArr, placeholdersStr)
		// add real variables for the replacement of placeholders' '?' letter later.
		mainScope.SQLVars = append(mainScope.SQLVars, scope.SQLVars...)
	}

	mainScope.Raw(fmt.Sprintf("INSERT INTO %s (%s) VALUES %s",
		mainScope.QuotedTableName(),
		strings.Join(quoted, ", "),
		strings.Join(placeholdersArr, ", "),
	))

	if _, err := mainScope.SQLDB().Exec(mainScope.SQL, mainScope.SQLVars...); err != nil {
		return err
	}
	return nil
}

func CloseDB() {
	defer db.Close()
}

func modelsError(err string) bool {
	if err != "" {
		AddLog("MODEL", err, "SYSTEM", "ERROR")
		fmt.Println("[Models Error]: ", err)
		return false
	}
	return true
}

func CreateDB() {
	db.Exec("CREATE DATABASE IF NOT EXISTS test_for")
	MigrateUsers()
	MigrateImages()
	MigrateSyncImage()
	MigrateStores()
	MigrateRole()
	MigratePermissions()
	MigrateOptions()
	MigrateLogs()
	MigrateDistribution()
}
