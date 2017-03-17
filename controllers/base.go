package controllers

import (
	"github.com/astaxie/beego"
	"github.com/juju/errors"
	"github.com/zssky/Mole/models/k8s"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) newK8sClient() (*k8s.K8sClient, error) {
	client, err := k8s.NewK8sClient(beego.AppConfig.DefaultString("k8s::config", ""))
	if err != nil {
		return nil, errors.Trace(err)
	}

	return client, nil
}

// TODO if needed, we can implement Init(), Prepare(), Finish()
