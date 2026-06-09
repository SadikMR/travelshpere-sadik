package routers

import (
	"github.com/SadikMR/travelshpere-sadik/controllers"
	"github.com/SadikMR/travelshpere-sadik/controllers/api"
	"github.com/SadikMR/travelshpere-sadik/filters"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// ── Filters ─────────────────────────────────────────────
	// SSR auth: redirect to /login
	beego.InsertFilter("/wishlist", beego.BeforeRouter, filters.AuthRequired)
	beego.InsertFilter("/wishlist/*", beego.BeforeRouter, filters.AuthRequired)

	// API auth: return 401 JSON
	beego.InsertFilter("/api/wishlist", beego.BeforeRouter, filters.APIAuthRequired)
	beego.InsertFilter("/api/wishlist/*", beego.BeforeRouter, filters.APIAuthRequired)

	// ── Routes ──────────────────────────────────────────────

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

	// Country API
	beego.Router(
		"/api/countries",
		&api.CountryController{},
		"get:Get",
	)

	beego.Router(
		"/api/countries/search",
		&controllers.HomeController{},
		"get:Search",
	)

	beego.Router(
		"/api/countries/:slug",
		&api.CountryController{},
		"get:Detail",
	)

	// Attractions API
	beego.Router(
		"/api/attractions",
		&api.AttractionController{},
		"get:Get",
	)

	// Wishlist SSR (protected by AuthRequired filter)
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

	// Wishlist API (protected by APIAuthRequired filter)
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

	// Dashboard SSR (protected by AuthRequired filter)
	beego.InsertFilter("/dashboard", beego.BeforeRouter, filters.AuthRequired)
	beego.Router(
		"/dashboard",
		&controllers.DashboardController{},
		"get:Get",
	)

	// Dashboard API (protected by APIAuthRequired filter)
	beego.InsertFilter("/api/dashboard/*", beego.BeforeRouter, filters.APIAuthRequired)
	beego.Router(
		"/api/dashboard/summary",
		&api.DashboardController{},
		"get:Summary",
	)
}
