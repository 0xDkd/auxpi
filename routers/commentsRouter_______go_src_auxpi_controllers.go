package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["auxpi/controllers:PagesController"] = append(beego.GlobalControllerRouter["auxpi/controllers:PagesController"],
        beego.ControllerComments{
            Method: "IndexShow",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auxpi/controllers:PagesController"] = append(beego.GlobalControllerRouter["auxpi/controllers:PagesController"],
        beego.ControllerComments{
            Method: "SinaShow",
            Router: `/Sina`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auxpi/controllers:PagesController"] = append(beego.GlobalControllerRouter["auxpi/controllers:PagesController"],
        beego.ControllerComments{
            Method: "SmmsShow",
            Router: `/Smms`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["auxpi/controllers:PagesController"] = append(beego.GlobalControllerRouter["auxpi/controllers:PagesController"],
        beego.ControllerComments{
            Method: "AboutShow",
            Router: `/about/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
