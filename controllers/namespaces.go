package controllers

type NamespacesController struct {
	BaseController
}

func (c *NamespacesController) Get() {
	c.Data["json"] = "namespaces"
	c.ServeJSON()
}
