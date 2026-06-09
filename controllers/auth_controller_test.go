package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func newAuthContext(req *http.Request, w *httptest.ResponseRecorder) *context.Context {
	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Output.Reset(ctx)
	return ctx
}

func TestAuthControllerLoginShowsErrorWhenUsernameIsMissing(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("username="))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	ctx := newAuthContext(req, w)

	ctrl := AuthController{BaseController: BaseController{Controller: beego.Controller{Ctx: ctx}}}
	ctrl.Data = make(map[interface{}]interface{})

	ctrl.Login()

	assert.Equal(t, "auth/login.tpl", ctrl.TplName)
	assert.Equal(t, "", ctrl.Layout)
	assert.Equal(t, "Username is required", ctrl.Data["Error"])
}

