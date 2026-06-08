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
	c.TplName = "auth/login.tpl"
}

// Login creates a session for the provided username.
func (c *AuthController) Login() {
	username := strings.TrimSpace(c.GetString("username"))

	if username == "" {
		c.Data["Error"] = "Username is required"
		c.TplName = "auth/login.tpl"
		return
	}

	c.SetSession("username", username)

	c.Redirect("/", 302)
}

// Logout removes the current session.
func (c *AuthController) Logout() {
	c.DestroySession()
	c.Redirect("/login", 302)
}
