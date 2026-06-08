package controllers

import "strings"

// AuthController handles login and logout.
type AuthController struct {
	BaseController
}

func (c *AuthController) LoginPage() {
	c.Layout = ""
	c.TplName = "auth/login.tpl"
}

func (c *AuthController) Login() {
	username := strings.TrimSpace(c.GetString("username"))

	if username == "" {
		c.Data["Error"] = "Username is required"
		c.Layout = ""
		c.TplName = "auth/login.tpl"
		return
	}

	c.SetSession("username", username)
	c.Redirect("/", 302)
}

func (c *AuthController) Logout() {
	c.DelSession("username")
	c.Redirect("/login", 302)
}
