package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/juju/errors"
	"github.com/zssky/Mole/models/k8s"
	"github.com/zssky/log"
	"github.com/zssky/tc"
	"github.com/zssky/tc/http"
)

type ServicesController struct {
	BaseController
}

func (c *ServicesController) Get() {
	defer c.ServeJSON()

	namespace := c.GetString(":namespace")

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	list, err := client.ListServices(namespace)
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

func (c *ServicesController) Post() {
	defer c.ServeJSON()

	var svcyaml string
	namespace := c.GetString(":namespace")

	var params k8s.VtgateServiceData
	if err := tc.DecodeJSON(c.Ctx.Input.RequestBody, &params); err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	reqType := strings.ToUpper(params.Type)
	if reqType == "ETCDLB" {
		svcyaml = beego.AppConfig.DefaultString("k8s::etcdlbsvcyaml", "")
	} else if reqType == "ETCDCB" {
		svcyaml = beego.AppConfig.DefaultString("k8s::etcdcbsvcyaml", "")
	} else if reqType == "VTCTLD" {
		svcyaml = beego.AppConfig.DefaultString("k8s::vtctldsvcyaml", "")
	} else if reqType == "VTGATE" {
		svcyaml = beego.AppConfig.DefaultString("k8s::vtgatesvcyaml", "")
	} else {
		c.Data["json"] = http.HttpResponse{Code: 1, Message: errors.Errorf("unknow request type").Error()}
		return
	}

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	svc, err := client.CreateService(namespace, svcyaml, params)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{
		Code:    0,
		Message: "success",
		Data: map[string]interface{}{
			"Service": svc,
		},
	}
}

func (c *ServicesController) Delete() {
	defer c.ServeJSON()

	namespace := c.GetString(":namespace")
	name := c.GetString(":name")

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	err = client.DeleteService(namespace, name, nil)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{Code: 0, Message: "success"}
}

func (c *ServicesController) Detail() {
	defer c.ServeJSON()

	namespace := c.GetString(":namespace")
	name := c.GetString(":name")

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	svc, err := client.GetService(namespace, name)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{
		Code:    0,
		Message: "success",
		Data: map[string]interface{}{
			"Service": svc,
		},
	}
}
