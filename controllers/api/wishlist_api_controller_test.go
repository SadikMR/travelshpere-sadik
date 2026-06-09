
package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func newAPIContext(req *http.Request, w *httptest.ResponseRecorder) *context.Context {
	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Output.Reset(ctx)
	return ctx
}

func TestWishlistControllerPostReturnsBadRequestOnInvalidPayload(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/wishlist", strings.NewReader("{invalid-json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx := newAPIContext(req, w)

	ctrl := WishlistController{BaseAPIController: BaseAPIController{Controller: beego.Controller{Ctx: ctx}, Username: "sadik"}}
	ctrl.Data = make(map[interface{}]interface{})
	ctrl.Post()

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid payload")
}
