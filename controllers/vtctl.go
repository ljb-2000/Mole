package controllers

type VtctlController struct {
	BaseController
}

func (c *VtctlController) Get() {
	c.Data["json"] = "vtctl"
	c.ServeJSON()
}
