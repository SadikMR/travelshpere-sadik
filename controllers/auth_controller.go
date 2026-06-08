package controllers

import "strings"

type AuthController struct {
	BaseController
}

// Override Prepare to skip session check for auth routes
func (c *AuthController) Prepare() {
	if c.Data == nil {
		c.Data = make(map[interface{}]interface{})
	}
	// Don't call BaseController.Prepare() here
	// No session needed for login/logout pages
	c.Data["IsLoggedIn"] = false
	c.Data["Username"] = ""
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
