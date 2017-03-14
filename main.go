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
	beego.BConfig.Listen.Graceful = beego.AppConfig.DefaultBool("graceful", false)
	beego.BConfig.WebConfig.Session.SessionOn = beego.AppConfig.DefaultBool("sessionOn", false)
	beego.BConfig.WebConfig.Session.SessionName = beego.AppConfig.DefaultString("sessionName", "Mole")

}

func main() {
	// init default config
	defaultConfig()

	beego.Run()
}
