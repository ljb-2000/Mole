package controllers

import (
	"github.com/astaxie/beego"
	"github.com/juju/errors"
	"github.com/zssky/Mole/models/k8s"
	"github.com/zssky/log"
	"github.com/zssky/tc"
	"github.com/zssky/tc/http"
)

type PodController struct {
	BaseController
}

func (c *PodController) Get() {
	defer c.ServeJSON()

	namespace := c.GetString(":namespace")

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	list, err := client.ListPods(namespace)
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

func (c *PodController) Post() {
	defer c.ServeJSON()

	namespace := c.GetString(":namespace")
	podyaml := beego.AppConfig.DefaultString("k8s::podyaml", "")

	var params k8s.VttabletPodData

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

	pod, err := client.CreatePod(namespace, podyaml, params)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{
		Code:    0,
		Message: "success",
		Data: map[string]interface{}{
			"Pod": pod,
		},
	}
}

func (c *PodController) Delete() {
	defer c.ServeJSON()

	namespace := c.GetString(":namespace")
	name := c.GetString(":name")

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	err = client.DeletePod(namespace, name, nil)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{Code: 0, Message: "success"}
}

func (c *PodController) Detail() {
	defer c.ServeJSON()

	namespace := c.GetString(":namespace")
	name := c.GetString(":name")

	client, err := c.newK8sClient()
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	pod, err := client.GetPod(namespace, name)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
		c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		return
	}

	c.Data["json"] = http.HttpResponse{
		Code:    0,
		Message: "success",
		Data: map[string]interface{}{
			"Pod": pod,
		},
	}
}
