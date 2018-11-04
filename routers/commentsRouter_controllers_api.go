package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["auxpi/controllers/api:ApiUpLoadController"] = append(beego.GlobalControllerRouter["auxpi/controllers/api:ApiUpLoadController"],
        beego.ControllerComments{
            Method: "AuthUpLoadHandle",
            Router: `/api/v1/upload/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
