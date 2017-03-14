package controllers

type PodController struct {
	BaseController
}

func (c *PodController) Get() {
	c.Data["json"] = "Pod"
	c.ServeJSON()
}
