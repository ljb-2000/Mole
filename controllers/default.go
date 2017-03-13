package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {

	c.Data["json"] = "default"
	c.ServeJSON()
}
