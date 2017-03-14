package controllers

type ServicesController struct {
	BaseController
}

func (c *ServicesController) Get() {
	c.Data["json"] = "services"
	c.ServeJSON()
}
