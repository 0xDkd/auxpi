package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["auxpi/controllers/upload:UpLoadController"] = append(beego.GlobalControllerRouter["auxpi/controllers/upload:UpLoadController"],
        beego.ControllerComments{
            Method: "AuthUpLoadHandle",
            Router: `/api/v1/auth/upload/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
