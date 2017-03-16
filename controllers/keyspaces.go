package controllers

import (
	"github.com/astaxie/beego"
	"github.com/zssky/Mole/models/vtctl"
	"github.com/zssky/tc/http"
)

type KeyspacesController struct {
	BaseController
}

func (c *KeyspacesController) Get() {

	defer c.ServeJSON()

	list, err := vtctl.KeyspacesList(beego.AppConfig.String("vtctl::server"))
	if err != nil {
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{Code: 0, Message: "success", Data: map[string]interface{}{
		"List": list,
	}}
}
