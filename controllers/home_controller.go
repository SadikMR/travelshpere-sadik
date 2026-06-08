package controllers

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	c.Layout = "layouts/base.tpl"
	c.TplName = "home/index.tpl"

	c.Data["Title"] = "Home"
}
