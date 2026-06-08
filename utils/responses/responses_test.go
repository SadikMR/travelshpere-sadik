package responses

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"

	beego "github.com/beego/beego/v2/server/web"
)

func TestWriteError(t *testing.T) {
	// Create a fake HTTP request/response pair
	r := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Build a beego context and controller
	ctx := context.NewContext()
	ctx.Reset(w, r)
	ctx.Output.Reset(ctx)

	ctrl := beego.Controller{}
	ctrl.Ctx = ctx
	ctrl.Data = make(map[interface{}]interface{})

	WriteError(&ctrl, http.StatusBadRequest, "something went wrong")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "something went wrong")
	assert.Contains(t, w.Body.String(), "400")
}
