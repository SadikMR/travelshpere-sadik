package controllers

import "github.com/beego/beego/v2/server/web"

// BaseController is the shared base for all SSR controllers.
// It populates session-based template data in Prepare().
type BaseController struct {
	web.Controller
}

// Prepare runs before every handler. Sets IsLoggedIn and Username for templates.
func (c *BaseController) Prepare() {
	username := c.GetSession("username")
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
