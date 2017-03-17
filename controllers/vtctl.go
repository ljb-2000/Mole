package controllers

import (
	"github.com/astaxie/beego"
	"github.com/juju/errors"
	"github.com/zssky/Mole/models/vtctl"
	"github.com/zssky/log"
	"github.com/zssky/tc"
	"github.com/zssky/tc/http"
)

type VtctlController struct {
	BaseController
}

func (c *VtctlController) Post() {
	defer c.ServeJSON()

	var params []string

	if err := tc.DecodeJSON(c.Ctx.Input.RequestBody, &params); err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	log.Debugf("params:%v", params)

	resp, err := vtctl.Vtctl(beego.AppConfig.String("vtctl::server"), &params)
	log.Debugf("resp:%v err:%v", resp, err)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	if resp != nil && len(resp.Error) > 0 {
		log.Errorf("%v", resp.Error)
		c.Data["json"] = http.HttpResponse{Code: 1, Message: resp.Error}
	}

	c.Data["json"] = http.HttpResponse{Code: 0, Message: "success", Data: map[string]interface{}{
		"Result": resp.Output,
	}}
}
