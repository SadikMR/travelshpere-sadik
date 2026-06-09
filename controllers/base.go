package controllers

import (
	"context"

	"github.com/beego/beego/v2/server/web"
)

// BaseController is the shared base for all SSR controllers.
// It populates session-based template data in Prepare().
type BaseController struct {
	web.Controller
}

// Prepare runs before every handler. Sets IsLoggedIn and Username for templates.
func (c *BaseController) Prepare() {
	if c.Data == nil {
		c.Data = make(map[interface{}]interface{})
	}

	var username interface{}
	if c.Ctx != nil && c.Ctx.Input != nil && c.Ctx.Input.CruSession != nil {
		username = c.Ctx.Input.CruSession.Get(context.Background(), "username")
	}

	if u, ok := username.(string); ok && u != "" {
		c.Data["IsLoggedIn"] = true
		c.Data["Username"] = u
	} else {
		c.Data["IsLoggedIn"] = false
		c.Data["Username"] = ""
	}
}

// GetUsername returns the logged-in username, or empty string if not authenticated.
func (c *BaseController) GetUsername() string {
	if u, ok := c.Data["Username"].(string); ok {
		return u
	}
	return ""
}
