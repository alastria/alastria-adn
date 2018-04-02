package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["hyperapi/controllers:CurrentuserController"] = append(beego.GlobalControllerRouter["hyperapi/controllers:CurrentuserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hyperapi/controllers:LuaChaincodeController"] = append(beego.GlobalControllerRouter["hyperapi/controllers:LuaChaincodeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hyperapi/controllers:LuaChaincodeController"] = append(beego.GlobalControllerRouter["hyperapi/controllers:LuaChaincodeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hyperapi/controllers:LuaChaincodeController"] = append(beego.GlobalControllerRouter["hyperapi/controllers:LuaChaincodeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:luaChaincodeId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hyperapi/controllers:LuaChaincodeController"] = append(beego.GlobalControllerRouter["hyperapi/controllers:LuaChaincodeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:luaChaincodeId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hyperapi/controllers:LuaChaincodeController"] = append(beego.GlobalControllerRouter["hyperapi/controllers:LuaChaincodeController"],
		beego.ControllerComments{
			Method: "Patch",
			Router: `/:luaChaincodeId`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hyperapi/controllers:OrgnameController"] = append(beego.GlobalControllerRouter["hyperapi/controllers:OrgnameController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hyperapi/controllers:UserController"] = append(beego.GlobalControllerRouter["hyperapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hyperapi/controllers:UserController"] = append(beego.GlobalControllerRouter["hyperapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hyperapi/controllers:UserController"] = append(beego.GlobalControllerRouter["hyperapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hyperapi/controllers:UserController"] = append(beego.GlobalControllerRouter["hyperapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hyperapi/controllers:UserController"] = append(beego.GlobalControllerRouter["hyperapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
