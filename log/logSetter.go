package auxpiLog

import (
	"auxpi/models"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func SetAnErrorLog(t string, err error) bool {
	if err != nil {
		str := fmt.Sprintf("%v", err)
		logs.Error(err)
		models.AddLog(t, str, "SYSTEM", "ERROR")
		return true
	}
	return false
}

func SetAWarningLog(t string, err error) bool {
	if err != nil {
		str := fmt.Sprintf("%v", err)
		logs.Warning(err)
		models.AddLog(t, str, "SYSTEM", "WARNING")
		return true
	}
	return false

}

func SetADebugLog(t string, err error) bool {
	if err != nil {
		str := fmt.Sprintf("%v", err)
		logs.Debug(err)
		models.AddLog(t, str, "SYSTEM", "DEBUG")
		return true
	}
	return false

}

func SetAImageLog(err error) bool {
	if err != nil {
		str := fmt.Sprintf("%v", err)
		logs.Error("[IMAGE ERROR]: ", str)
		models.AddLog("IMAGE", str, "SYSTEM", "ERROR")
		return true
	}
	return false
}

func SetAOptionLog(err error) bool {
	if err != nil {
		str := fmt.Sprintf("%v", err)
		logs.Error(err)
		models.AddLog("OPTION", str, "SYSTEM", "ERROR")
		return true
	}
	return false
}

func SetALog(t string, err error, part string) bool {
	if err != nil {
		str := fmt.Sprintf("%v", err)
		logs.Error(err)
		models.AddLog(t, str, part, "ERROR")
		return true
	}
	return false
}

func SetUserLogin(str string)  {
	models.AddLog("USER_LOGIN",str,"SYSTEM","INFO")
}
