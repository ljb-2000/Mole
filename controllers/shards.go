package controllers

type ShardsController struct {
	BaseController
}

func (c *ShardsController) Get() {
	c.Data["json"] = "shards"
	c.ServeJSON()
}
