package controllers

type KeyspacesController struct {
	BaseController
}

func (c *KeyspacesController) Get() {
	c.Data["json"] = "keyspaces"
	c.ServeJSON()
}
