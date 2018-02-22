package controllers

import (
	"hyperapi/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about luaChaincode
type LuaChaincodeController struct {
	beego.Controller
}

// @Title Create
// @Description create luaChaincode
// @Param	body		body 	models.LuaChaincode	true		"The luaChaincode content"
// @Success 200 {string} models.LuaChaincode.Id
// @Failure 403 body is empty
// @router / [post]
func (o *LuaChaincodeController) Post() {
	var ob models.LuaChaincode
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	luaChaincodeid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"LuaChaincodeId": luaChaincodeid}
	o.ServeJSON()
}

// @Title Get
// @Description find luaChaincode by luaChaincodeid
// @Param	luaChaincodeId		path 	string	true		"the luaChaincodeid you want to get"
// @Success 200 {luaChaincode} models.LuaChaincode
// @Failure 403 :luaChaincodeId is empty
// @router /:luaChaincodeId [get]
func (o *LuaChaincodeController) Get() {
	luaChaincodeId := o.Ctx.Input.Param(":luaChaincodeId")
	if luaChaincodeId != "" {
		ob, err := models.GetOne(luaChaincodeId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all luaChaincodes
// @Success 200 {luaChaincode} models.LuaChaincode
// @Failure 403 :luaChaincodeId is empty
// @router / [get]
func (o *LuaChaincodeController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the luaChaincode
// @Param	luaChaincodeId		path 	string	true		"The luaChaincodeid you want to update"
// @Param	body		body 	models.LuaChaincode	true		"The body"
// @Success 200 {luaChaincode} models.LuaChaincode
// @Failure 403 :luaChaincodeId is empty
// @router /:luaChaincodeId [put]
func (o *LuaChaincodeController) Put() {
	luaChaincodeId := o.Ctx.Input.Param(":luaChaincodeId")
	var user models.User
	json.Unmarshal(o.Ctx.Input.RequestBody, &user)

	err := models.Update(luaChaincodeId, user.Id, true)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

