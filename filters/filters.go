package filters

import (
	"net/http"

	"github.com/beego/beego/v2/server/web/context"
)

// AuthRequired redirects unauthenticated users to /login.
// Use for SSR routes that require a logged-in user.
func AuthRequired(ctx *context.Context) {
	username := ctx.Input.Session("username")
	if username == nil || username == "" {
		ctx.Redirect(http.StatusFound, "/login")
	}
}

// APIAuthRequired returns 401 JSON for unauthenticated API requests.
// Use for API routes that require a logged-in user.
func APIAuthRequired(ctx *context.Context) {
	username := ctx.Input.Session("username")
	if username == nil || username == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]string{
			"error": "login required",
		}, false, false)
	}
}
