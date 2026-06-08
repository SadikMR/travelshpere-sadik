package routers

import (
	"github.com/SadikMR/travelshpere-sadik/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
