package controllers

import (
	"github.com/astaxie/beego"
	"github.com/juju/errors"
	"github.com/zssky/Mole/models/vtctl"
	"github.com/zssky/log"
	"github.com/zssky/tc"
	"github.com/zssky/tc/http"
)

type SchemaController struct {
	BaseController
}

func (c *SchemaController) Apply() {
	defer c.ServeJSON()

	var params struct {
		Keyspace string `json:"Keyspace"`
		SQL      string `json:"SQL"`
	}

	if err := tc.DecodeJSON(c.Ctx.Input.RequestBody, &params); err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	log.Debugf("params:%v", params)

	resp, err := vtctl.SchemaApply(beego.AppConfig.String("vtctl::server"), params.Keyspace, params.SQL)
	log.Debugf("resp:%v err:%v", resp, err)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{Code: 0, Message: "success", Data: map[string]interface{}{
		"Result": resp.Output,
	}}

}
