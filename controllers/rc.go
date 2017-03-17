package controllers

import (
	"github.com/astaxie/beego"
	"github.com/juju/errors"
	"github.com/zssky/Mole/models/k8s"
	"github.com/zssky/log"
	"github.com/zssky/tc"
	"github.com/zssky/tc/http"
)

type RcController struct {
	BaseController
}

func (c *RcController) Get() {
	defer c.ServeJSON()

	namespace := c.GetString(":namespace")

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	list, err := client.ListRC(namespace)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{
		Code:    0,
		Message: "success",
		Data:    list,
	}
}

func (c *RcController) Post() {
	defer c.ServeJSON()

	namespace := c.GetString(":namespace")
	rcyaml := beego.AppConfig.DefaultString("k8s::rcyaml", "")

	var params k8s.VtgateRCData
	if err := tc.DecodeJSON(c.Ctx.Input.RequestBody, &params); err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	rc, err := client.CreateRC(namespace, rcyaml, params)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{
		Code:    0,
		Message: "success",
		Data: map[string]interface{}{
			"ReplicationController": rc,
		},
	}
}

func (c *RcController) Delete() {
	defer c.ServeJSON()

	namespace := c.GetString(":namespace")
	name := c.GetString(":name")

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	err = client.DeleteRC(namespace, name, nil)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{Code: 0, Message: "success"}
}

func (c *RcController) Detail() {
	defer c.ServeJSON()

	namespace := c.GetString(":namespace")
	name := c.GetString(":name")

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	rc, err := client.GetRC(namespace, name)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{
		Code:    0,
		Message: "success",
		Data: map[string]interface{}{
			"ReplicationController": rc,
		},
	}
}
