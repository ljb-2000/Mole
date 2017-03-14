package controllers

type RcController struct {
	BaseController
}

func (c *RcController) Get() {
	c.Data["json"] = "rc"
	c.ServeJSON()
}
