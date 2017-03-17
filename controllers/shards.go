package controllers

import (
	"github.com/astaxie/beego"
	"github.com/juju/errors"
	"github.com/zssky/Mole/models/vtctl"
	"github.com/zssky/log"
	"github.com/zssky/tc/http"
)

type ShardsController struct {
	BaseController
}

func (c *ShardsController) Get() {
	defer c.ServeJSON()

	keyspace := c.GetString(":keyspace")

	list, err := vtctl.Shards(beego.AppConfig.String("vtctl::server"), keyspace)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{Code: 0, Message: "success", Data: map[string]interface{}{
		"List": list,
	}}

}
