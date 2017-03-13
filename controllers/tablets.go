package controllers

type TabletsController struct {
	BaseController
}

func (c *TabletsController) Get() {
	c.Data["json"] = "tablets"
	c.ServeJSON()
}
