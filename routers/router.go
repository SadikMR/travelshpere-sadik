package routers

import (
	"github.com/SadikMR/travelshpere-sadik/controllers"
	"github.com/SadikMR/travelshpere-sadik/controllers/api"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// Home
	beego.Router("/", &controllers.HomeController{})

	// Authentication
	beego.Router(
		"/login",
		&controllers.AuthController{},
		"get:LoginPage;post:Login",
	)

	beego.Router(
		"/logout",
		&controllers.AuthController{},
		"get:Logout;post:Logout",
	)

	// Countries
	beego.Router(
		"/countries",
		&controllers.CountryController{},
		"get:Get",
	)

	beego.Router(
		"/countries/:slug",
		&controllers.CountryController{},
		"get:Details",
	)

	// Wishlist SSR
	beego.Router(
		"/wishlist",
		&controllers.WishlistController{},
		"get:Get",
	)

	beego.Router(
		"/wishlist/rows",
		&controllers.WishlistController{},
		"get:Rows",
	)

	// Wishlist API
	beego.Router(
		"/api/wishlist",
		&api.WishlistController{},
		"get:Get;post:Post",
	)

	beego.Router(
		"/api/wishlist/:id",
		&api.WishlistController{},
		"put:Put;delete:Delete",
	)
}