package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func newControllerContext(req *http.Request, w *httptest.ResponseRecorder) *context.Context {
	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Output.Reset(ctx)
	return ctx
}

func TestBaseControllerPrepareSetsLoggedOutWhenNoSession(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := newControllerContext(req, w)

	ctrl := BaseController{Controller: beego.Controller{Ctx: ctx}}
	ctrl.Data = make(map[interface{}]interface{})

	ctrl.Prepare()

	assert.False(t, ctrl.Data["IsLoggedIn"].(bool))
	assert.Equal(t, "", ctrl.Data["Username"])
}

