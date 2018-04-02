package controllers

import (
	"github.com/astaxie/beego"
	"hyperapi/models"
)

// Operations about Users
type CurrentuserController struct {
	beego.Controller
}

// @Title Get
// @Description string models.User
// @Success 200 {string} string
// @router / [get]
func (u *CurrentuserController) Get() {
	u.Data["json"] = models.GetCurrentUser()
	u.ServeJSON()
}
