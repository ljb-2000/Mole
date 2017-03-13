package controllers

type SchemaController struct {
	BaseController
}

func (c *SchemaController) Get() {
	c.Data["json"] = "schema"
	c.ServeJSON()
}
