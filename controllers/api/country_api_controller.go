package api

import (
	"net/http"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/SadikMR/travelshpere-sadik/services"
	"github.com/SadikMR/travelshpere-sadik/utils/responses"
	"github.com/SadikMR/travelshpere-sadik/utils/validators"
)

// CountryController handles country API requests.
type CountryController struct {
	beego.Controller
}

var countryService = services.CountryService{}

// Get returns countries filtered by search and region.
func (c *CountryController) Get() {
	search := c.GetString("search")
	region := c.GetString("region")

	if err := validators.ValidateCountrySearch(search, region); err != nil {
		responses.WriteError(
			&c.Controller,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	countries, err := countryService.GetCountries(
		search,
		region,
	)
	if err != nil {
		responses.WriteError(
			&c.Controller,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	c.Data["json"] = countries
	c.ServeJSON()
}

// Detail returns JSON detail for a single country by slug.
func (c *CountryController) Detail() {
	slug := c.Ctx.Input.Param(":slug")

	country, err := countryService.GetCountryBySlug(slug)
	if err != nil {
		responses.WriteError(
			&c.Controller,
			http.StatusNotFound,
			"country not found",
		)
		return
	}

	c.Data["json"] = country
	c.ServeJSON()
}
