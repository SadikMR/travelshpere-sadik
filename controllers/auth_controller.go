package controllers

import (
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

// AuthController handles authentication-related requests.
type AuthController struct {
	beego.Controller
}

// LoginPage renders the login page.
func (c *AuthController) LoginPage() {
	// Don't use layout for login page
	c.Layout = ""
	c.TplName = "auth/login.tpl"
}

// Login creates a session for the provided username.
func (c *AuthController) Login() {
	username := strings.TrimSpace(c.GetString("username"))

	if username == "" {
		if c.Data == nil {
			c.Data = make(map[interface{}]interface{})
		}
		c.Data["Error"] = "Username is required"
		c.Layout = ""
		c.TplName = "auth/login.tpl"
		return
	}

	// Set cookie to store username
	c.Ctx.SetCookie("username", username, 3600, "/")

	// Redirect to home
	c.Redirect("/", 302)
}

// Logout removes the current session.
func (c *AuthController) Logout() {
	// Delete the username cookie
	c.Ctx.SetCookie("username", "", -1, "/")
	c.Redirect("/login", 302)
}
