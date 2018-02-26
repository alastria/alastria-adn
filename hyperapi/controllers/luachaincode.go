package controllers

import (
	"hyperapi/models"
	"encoding/json"
    "fmt"

	"github.com/astaxie/beego"
)

// Operations about luaChaincode
type LuaChaincodeController struct {
	beego.Controller
}

// @Title Create
// @Description create LuaChaincodePost	
// @Param	body		body 	models.LuaChaincodePost	true		"The luaChaincode content"
// @Success 200 {string} models.LuaChaincode.Id
// @Failure 403 body is empty
// @router / [post]
func (o *LuaChaincodeController) Post() {
	var ob models.LuaChaincodePost
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
    fmt.Println(o.Ctx.Input.RequestBody)
    fmt.Println(ob.Name)
    fmt.Println(ob.SourceCode)
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
// @Description update the luaChaincode, validate chaincode
// @Param	luaChaincodeId		path 	string	true		"The luaChaincodeid you want to validate"
// @Success 200 {luaChaincode} models.LuaChaincode
// @Failure 403 :luaChaincodeId is empty
// @router /:luaChaincodeId [put]
func (o *LuaChaincodeController) Put() {
	luaChaincodeId := o.Ctx.Input.Param(":luaChaincodeId")
	err := models.Update(luaChaincodeId)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Execute
// @Description excute the lua luaChaincode
// @Param	luaChaincodeId		path 	string	true		"The luaChaincodeid you want to execute"
// @Success 200 {luaChaincode} execution return 
// @Failure 403 :luaChaincodeId is empty
// @router /:luaChaincodeId [patch]
func (o *LuaChaincodeController) Patch() {
	luaChaincodeId := o.Ctx.Input.Param(":luaChaincodeId")
	result := models.Execute(luaChaincodeId)
    o.Data["json"] = result 
	o.ServeJSON()
}


