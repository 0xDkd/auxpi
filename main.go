package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/bootstrap"
	"github.com/auxpi/controllers"
	"github.com/auxpi/models"
	_ "github.com/auxpi/routers"
	"github.com/auxpi/utils"

	_ "net/http/pprof"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	args := os.Args
	p := "\033[32m"
	s := "\033[0m"
	if len(args) <= 1 {
		beego.Run()
	} else {
		switch args[1] {
		case "init":
			if _, err := os.Stat("conf/install.lock"); err == nil {
				fmt.Println(p + "[INFO]:AUXPI have already initialization complete.\r\n[INFO]:Please run \"./auxpi run\"  to start " + s)
				return
			}

			//初始化格式
			formatData()
			//日志文件
			msg := p + `[SUCCESS]:Init completed ! ` + "\r" + `[INFO]:Please edit "conf/siteConfig" to make auxpi can connect to database` + s
			fmt.Println(msg)
		case "run":
			err := logs.SetLogger(logs.AdapterFile, `{"filename":"auxpi.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
			if err != nil {
				beego.Alert(err)
			}
			beego.Run()
		case "migrate":
			var errs = make([]error, 10, 10)
			errs[3] = models.MigrateUsers()
			errs[0] = models.MigrateImages()
			errs[1] = models.MigrateSyncImage()
			errs[2] = models.MigrateStores()
			errs[4] = models.MigrateRole()
			errs[5] = models.MigratePermissions()
			errs[6] = models.MigrateOptions()
			errs[7] = models.MigrateLogs()
			errs[8] = models.MigrateDistribution()
			models.CreateAdminRole()
			for _, err := range errs {
				if err != nil {
					fmt.Println("\033[31m[ERROR]:\033" + auxpi.ErrorToString(err))
					return
				}
			}
			fmt.Println("\033[32m[INFO]:\033[0m Migrate done")
		}
	}

}

func formatData() {
	bootstrap.FormatSoftLink(&bootstrap.SiteConfig.SiteUploadWay.LocalStore.Link)
	bootstrap.FormatStoreLocation(&bootstrap.SiteConfig.SiteUploadWay.LocalStore.StorageLocation)
	//强制重载配置，防止用户填入的数据不规范
	bootstrap.ReGenerate()
}

func logo() {
	b, err := bootstrap.Asset("views/logo/index.tpl")
	if err != nil {
		return
	}
	p := "\033[34m"
	s := "\033[0m"
	fmt.Println(p + string(b) + s)
}

func init() {
	//测试权限
	//models.CreateDB()
	//models.Create()

	logo()
	beego.BConfig.ServerName = "Auxpi-Server-2.0"
	//不可修改静态服务器,请不要修改,否则将影响后台图片显示
	beego.SetStaticPath("/backup", bootstrap.SiteConfig.SiteUploadWay.LocalStore.StorageLocation)
	//自定义静态服务器
	beego.SetStaticPath(bootstrap.SiteConfig.SiteUploadWay.LocalStore.Link, bootstrap.SiteConfig.SiteUploadWay.LocalStore.StorageLocation)
	beego.SetStaticPath("/.well-known", ".well-known")
	//SessionID 重写
	beego.BConfig.WebConfig.Session.SessionName = "PHPSESSIONID"
	//ERROR 自定义
	beego.ErrorController(&controllers.ErrorController{})
	//日志记录
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"auxpi.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	if err != nil {
		beego.Alert(err)
	}

	var (
		mod    string
		user   string
		pass   string
		email  string
		dbName string
		dbUser string
		dbPass string
	)

	flag.StringVar(&user, "name", "auxpi-admin", "Admin UserName")
	flag.StringVar(&pass, "pass", "admin-pass", "Admin PassWord")
	flag.StringVar(&email, "email", "auxpi@0w0.tn", "Admin Email")
	flag.StringVar(&mod, "mod", "", "Choose Module")
	flag.StringVar(&dbName, "dbName", "auxpi", "dataBase Name")
	flag.StringVar(&dbUser, "dbUser", "root", "dataBase UserName")
	flag.StringVar(&dbPass, "dbPass", "root", "dataBase PassWord")
	flag.Parse()
	t := utils.GetSha256CodeWithSalt("auxpiauxpi")

	switch mod {
	case "admin":
		if u, _ := models.GetUserInfoByID(1); u.ID > 0 {
			fmt.Println("\033[31m[ERROR]:\033[0m Admin Is Existed")
			return
		}
		models.RegisterAdmin(user, utils.GetSha256CodeWithSalt(pass), t, email)
		fmt.Println("\033[32m[INFO]:\033[0m Create Admin SUCCESS")
	case "migrate":
		if dbName == "" && dbUser == "" && dbPass == "" {
			fmt.Println("\033[31m[ERROR]:\033" + "dbName,dbUser,dbPass can't be empty")
			return
		}

		options := bootstrap.SiteConfig
		options.DbOption.DbName = dbName
		options.DbOption.DbUser = dbUser
		options.DbOption.DbPass = dbPass
		beego.Alert(dbPass, dbUser, dbName)
		err := bootstrap.ReGenerateByInput(*options)
		if err != nil {
			fmt.Println("\033[31m[ERROR]:\033" + auxpi.ErrorToString(err))
		}
		var errs = make([]error, 10, 10)
		errs[0] = models.MigrateImages()
		if errs[0] != nil {
			fmt.Println("\033[31m[ERROR]:\033" + auxpi.ErrorToString(errs[0]))
			return
		}
		errs[3] = models.MigrateUsers()
		errs[0] = models.MigrateImages()
		errs[1] = models.MigrateSyncImage()
		errs[2] = models.MigrateStores()
		errs[4] = models.MigrateRole()
		errs[5] = models.MigratePermissions()
		errs[6] = models.MigrateOptions()
		errs[7] = models.MigrateLogs()
		errs[8] = models.MigrateDistribution()
		models.CreateAdminRole()
		for _, err := range errs {
			if err != nil {
				fmt.Println("\033[31m[ERROR]:\033" + auxpi.ErrorToString(err))
				return
			}
		}
		fmt.Println("\033[32m[INFO]:\033[0m Migrate done")
	}
}
