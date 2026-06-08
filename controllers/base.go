package controllers

import "github.com/beego/beego/v2/server/web"

type BaseController struct {
	web.Controller
}

func (c *BaseController) Prepare() {
	if c.Data == nil {
		c.Data = make(map[interface{}]interface{})
	}

	// Guard against uninitialized session
	defer func() {
		if r := recover(); r != nil {
			c.Data["IsLoggedIn"] = false
			c.Data["Username"] = ""
		}
	}()

	username := c.GetSession("username")
	if username != nil {
		c.Data["IsLoggedIn"] = true
		c.Data["Username"] = username.(string)
	} else {
		c.Data["IsLoggedIn"] = false
		c.Data["Username"] = ""
	}
}
