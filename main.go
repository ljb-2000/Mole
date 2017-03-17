package main

import (
	"github.com/astaxie/beego"
	_ "github.com/zssky/Mole/routers"
)

func defaultConfig() {
	beego.BConfig.AppName = beego.AppConfig.DefaultString("default::appname", "Mole")
	beego.BConfig.RunMode = beego.AppConfig.DefaultString("default::runmode", "prod")
	beego.BConfig.Listen.EnableAdmin = beego.AppConfig.DefaultBool("default::enableAdmin", false)
	beego.BConfig.WebConfig.AutoRender = beego.AppConfig.DefaultBool("default::autoRender", true)
	beego.BConfig.WebConfig.DirectoryIndex = beego.AppConfig.DefaultBool("default::directoryIndex", false)
	beego.BConfig.CopyRequestBody = beego.AppConfig.DefaultBool("default::copyRequestBody", true)
	beego.BConfig.Listen.Graceful = beego.AppConfig.DefaultBool("default::graceful", false)
	beego.BConfig.WebConfig.Session.SessionOn = beego.AppConfig.DefaultBool("default::sessionOn", false)
	beego.BConfig.WebConfig.Session.SessionName = beego.AppConfig.DefaultString("default::sessionName", "Mole")
	beego.BConfig.RecoverPanic = beego.AppConfig.DefaultBool("default::recoverPanic", true)

}

func main() {
	// init default config
	defaultConfig()

	beego.Run()
}
