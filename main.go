package main

import (
	"strings"

	_ "github.com/SadikMR/travelshpere-sadik/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// Fallback in case app.conf is not loaded
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.BConfig.WebConfig.Session.SessionName = "travelsphere_sess"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600
	beego.BConfig.WebConfig.ViewsPath = "views"

	// Register template helper: {{.Name | slugify}}
	beego.AddFuncMap("slugify", func(s string) string {
		return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
	})

	beego.Run()
}
