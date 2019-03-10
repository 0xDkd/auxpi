package main

import (
	"auxpi/auxpiAll"
	"auxpi/bootstrap"
	"auxpi/controllers"
	"auxpi/models"
	_ "auxpi/routers"
	"auxpi/utils"
	"flag"
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	args := os.Args
	p := "\033[32m"
	s := "\033[0m"
	//pr := "\033[31m"
	beego.Alert(args)
	if len(args) <= 1 {
		err := logs.SetLogger(logs.AdapterFile, `{"filename":"auxpi.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
		if err != nil {
			beego.Alert(err)
		}
		beego.BConfig.ServerName = "Auxpi-Server-2.0"
		//不可修改静态服务器,请不要修改,否则将影响后台图片显示
		beego.SetStaticPath("/backup", bootstrap.SiteConfig.SiteUploadWay.LocalStore.StorageLocation)
		//自定义静态服务器
		beego.SetStaticPath(bootstrap.SiteConfig.SiteUploadWay.LocalStore.Link, bootstrap.SiteConfig.SiteUploadWay.LocalStore.StorageLocation)
		//SessionID 重写
		beego.BConfig.WebConfig.Session.SessionName = "PHPSESSION"
		//ERROR 自定义
		beego.ErrorController(&controllers.ErrorController{})
		beego.Run()
	}else {
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
			beego.BConfig.ServerName = "Auxpi-Server-2.0"
			//不可修改静态服务器,请不要修改,否则将影响后台图片显示
			beego.SetStaticPath("/backup", bootstrap.SiteConfig.SiteUploadWay.LocalStore.StorageLocation)
			//自定义静态服务器
			beego.SetStaticPath(bootstrap.SiteConfig.SiteUploadWay.LocalStore.Link, bootstrap.SiteConfig.SiteUploadWay.LocalStore.StorageLocation)
			//SessionID 重写
			beego.BConfig.WebConfig.Session.SessionName = "PHPSESSION"
			//ERROR 自定义
			beego.ErrorController(&controllers.ErrorController{})
			beego.Run()
		case "migrate":
			var errs = make([]error, 8, 10)
			errs[0] = models.MigrateImages()
			errs[1] = models.MigrateSyncImage()
			errs[2] = models.MigrateStores()
			errs[3] = models.MigrateUsers()
			errs[4] = models.MigrateRole()
			errs[5] = models.MigratePermissions()
			errs[6] = models.MigrateOptions()
			errs[7] = models.MigrateLogs()
			for _, err := range errs {
				if err != nil {
					fmt.Println("[ERROR]:" + auxpi.ErrorToString(err))
					return
				}
			}
			fmt.Println(p + "[SUCCESS]: Database migrate Done" + s)
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
	logo()
	mod := flag.String("mod", "", "Choose Module")
	user := flag.String("name", "auxpi-admin", "Admin UserName")
	password := flag.String("pass", "admin", "Admin Pass")
	email := flag.String("email", "test@0w0.tn", "Admin Email")
	p := utils.GetSha256CodeWithSalt(*password)
	token := flag.String("token", "sakdjo9wasd", "User API Token")
	t := utils.GetSha256CodeWithSalt(*token)
	flag.Parse()
	if *mod != "" {
		if u, _ := models.GetUserInfoByID(1); u.ID > 0 {
			fmt.Println("\033[31m[ERROR]:Admin Is Existed\033[0m")
			return
		}
		models.RegisterAdmin(*user, p, t, *email)
		fmt.Println("\033[32m[SUCCESS]:Create Admin SUCCESS\033[0m")
	}
}
