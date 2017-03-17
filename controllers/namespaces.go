package controllers

import (
	"github.com/astaxie/beego"
	"github.com/juju/errors"
	"github.com/zssky/Mole/models/k8s"
	"github.com/zssky/log"
	"github.com/zssky/tc"
	"github.com/zssky/tc/http"
)

type NamespacesController struct {
	BaseController
}

func (c *NamespacesController) Get() {
	defer c.ServeJSON()

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	list, err := client.ListNamespaces()
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

func (c *NamespacesController) Post() {
	defer c.ServeJSON()

	nsyaml := beego.AppConfig.DefaultString("k8s::nsyaml", "")

	var params k8s.NamespaceData
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

	svc, err := client.CreateNamespace(nsyaml, params)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{
		Code:    0,
		Message: "success",
		Data: map[string]interface{}{
			"Namespace": svc,
		},
	}
}

func (c *NamespacesController) Delete() {
	defer c.ServeJSON()

	name := c.GetString(":name")

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	err = client.DeleteNamespace(name, nil)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{Code: 0, Message: "success"}
}

func (c *NamespacesController) Detail() {
	defer c.ServeJSON()

	name := c.GetString(":name")

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	ns, err := client.GetNamespace(name)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{
		Code:    0,
		Message: "success",
		Data: map[string]interface{}{
			"Namespace": ns,
		},
	}
}
