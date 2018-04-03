package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["apoyoalimentario_MID_API/controllers:TipoApoyoController"] = append(beego.GlobalControllerRouter["apoyoalimentario_MID_API/controllers:TipoApoyoController"],
		beego.ControllerComments{
			Method: "CalculateScore",
			Router: `/:smlv/:code`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

}
