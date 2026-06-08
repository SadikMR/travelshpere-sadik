package api

import (
	beego "github.com/beego/beego/v2/server/web"
)

// BaseAPIController is the shared base for all API controllers.
// Extracts Username from session in Prepare(). Auth is enforced
// by the APIAuthRequired filter, so Username is guaranteed to be set.
type BaseAPIController struct {
	beego.Controller
	Username string
}

// Prepare extracts the authenticated username from the session.
func (c *BaseAPIController) Prepare() {
	if u, ok := c.GetSession("username").(string); ok {
		c.Username = u
	}
}
