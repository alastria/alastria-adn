package main


import (
	_ "hyperapi/routers"
	"github.com/astaxie/beego"
)

var initArgs = [][]byte{[]byte("init")}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

