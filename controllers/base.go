package controllers

import "github.com/beego/beego/v2/server/web"

type BaseController struct {
	web.Controller
}

func (c *BaseController) Prepare() {
	// Initialize Data if nil
	if c.Data == nil {
		c.Data = make(map[interface{}]interface{})
	}

	var isLoggedIn bool

	// Check for username in cookie
	usernameCookie := c.Ctx.GetCookie("username")
	if usernameCookie != "" {
		isLoggedIn = true
	}

	c.Data["IsLoggedIn"] = isLoggedIn
}
