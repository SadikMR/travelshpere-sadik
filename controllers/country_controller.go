package controllers

import (
	"github.com/SadikMR/travelshpere-sadik/services"
	"github.com/SadikMR/travelshpere-sadik/utils/validators"
)

// CountryController handles country pages.
type CountryController struct {
	BaseController
}

var countryService = services.CountryService{}

// Get renders the country explorer page.
func (c *CountryController) Get() {
	search := c.GetString("search")
	region := c.GetString("region")

	if err := validators.ValidateCountrySearch(search, region); err != nil {
		c.CustomAbort(400, err.Error())
		return
	}

	countries, err := countryService.GetCountries(search, region)
	if err != nil {
		c.CustomAbort(500, err.Error())
		return
	}

	c.Data["Countries"] = countries
	c.Data["Title"] = "Country Explorer"
	c.Layout = "layouts/base.tpl"
	c.TplName = "countries.tpl"
}

func (c *CountryController) Details() {
	slug := c.Ctx.Input.Param(":slug")

	country, err := countryService.GetCountryBySlug(slug)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["Message"] = "Country not found"
		c.Layout = "layouts/base.tpl"
		c.TplName = "404.tpl"
		return
	}

	attractions, err := services.GetAttractions(
		country.Latitude,
		country.Longitude,
	)
	if err != nil {
		attractions = nil
	}

	c.Data["Country"] = country
	c.Data["Attractions"] = attractions
	c.Data["IsWishlisted"] = false // update when wishlist service is ready
	c.Data["Title"] = country.Name
	c.Layout = "layouts/base.tpl"
	c.TplName = "destination.tpl"
}
