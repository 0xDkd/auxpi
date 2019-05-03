package v1

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/auxpiAll/e"
	"github.com/auxpi/bootstrap"
	"github.com/auxpi/controllers/api/base"
	"github.com/auxpi/log"
	"github.com/auxpi/utils"

	"github.com/astaxie/beego/toolbox"
)

type AuxpiInfo struct {
	base.ApiController
}

//获取应用的各种信息
func GetAuxpiNowInfo() {

}

//统计各种系统信息(仅仅调试的时候使用)
func (a *AuxpiInfo) GetAuxpiSystemInfo() {
	gr := new(bytes.Buffer)
	hp := new(bytes.Buffer)
	tc := new(bytes.Buffer)
	gs := new(bytes.Buffer)
	toolbox.ProcessInput("lookup goroutine", gr)
	toolbox.ProcessInput("lookup heap", hp)
	toolbox.ProcessInput("lookup threadcreate", tc)
	toolbox.ProcessInput("gc summary", gs)
	data := make(map[string]string)
	data["goroutines"] = gr.String()
	//data["goroutines"] = strings.Replace(data["goroutines"], "\t", `&nbsp;&nbsp;&nbsp;&nbsp;`, -1)
	//data["goroutines"] = strings.Replace(data["goroutines"], "\n", `<br>`, -1)
	data["heap"] = hp.String()
	data["thread"] = tc.String()
	data["gc"] = gs.String()

	a.Data["json"] = &auxpi.RespJson{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
		Data: data,
	}
	a.ServeJSON()
}

//获取 ras key 的
func (a *AuxpiInfo) GetRsaKey() {
	pb, pv := utils.GetRsaKey()
	data := make(map[string]string)
	data["private_key"] = string(pv)
	data["public_key"] = string(pb)
	a.Data["json"] = &auxpi.RespJson{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
		Data: data,
	}
	a.ServeJSON()
}

//获取配置
func (a *AuxpiInfo) GetSiteConf() {
	a.Data["json"] = &auxpi.RespJson{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
		Data: bootstrap.SiteConfig,
	}
	a.ServeJSON()
}

//更新并且重载配置
func (a *AuxpiInfo) ResetSiteConf() {
	config := auxpi.SiteConfig{}
	err := config.UnmarshalJSON(a.Ctx.Input.RequestBody)
	if err != nil {
		a.Data["json"] = auxpi.RespJson{
			Code: 500,
			Msg:  e.GetMsg(500),
		}
		auxpiLog.SetAnErrorLog("CONFIG_RESET", err)
		a.ServeJSON()
		return
	}
	err = bootstrap.ReGenerateByInput(config)
	if err != nil {
		a.Data["json"] = auxpi.RespJson{
			Code: 500,
			Msg:  e.GetMsg(500),
		}
		auxpiLog.SetAnErrorLog("CONFIG_RESET", err)
		a.ServeJSON()
		return
	}
	a.Data["json"] = auxpi.RespJson{
		Code: 200,
		Msg:  "ok",
	}
	a.ServeJSON()

}

//检查更新
func (a *AuxpiInfo) CheckUpdate() {
	resp, err := http.Get("https://www.0w0.tn/auxpi/version")
	if err != nil {
		auxpiLog.SetAnErrorLog("CHECK_UPDATE_ERROR", err)
		a.Data["json"] = auxpi.RespJson{
			Code: e.ERROR,
			Msg:  e.GetMsg(e.ERROR),
		}
		a.ServeJSON()
		return
	}
	data := make(map[string]interface{})
	b, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &data)
	if err != nil {
		auxpiLog.SetAnErrorLog("CHECK_UPDATE_ERROR", err)
		a.Data["json"] = auxpi.RespJson{
			Code: e.ERROR,
			Msg:  "json 解析失败",
		}
		a.ServeJSON()
		return
	}

	a.Data["json"] = auxpi.RespJson{
		Code: e.SUCCESS,
		Msg:  "ok",
		Data: data,
	}
	a.ServeJSON()
}
