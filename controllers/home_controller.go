package controllers

import (
	"strings"

	"github.com/SadikMR/travelshpere-sadik/services"
)

type HomeController struct {
	BaseController
}

// Get — SSR homepage. Countries + attractions from real services.
func (c *HomeController) Get() {
	c.Layout = "layouts/base.tpl"
	c.TplName = "home/index.tpl"
	c.Data["Title"] = "Home"
	c.Data["ActivePage"] = "home"

	// Featured countries — first 6 alphabetically
	svc := &services.CountryService{}
	countries, err := svc.GetCountries("", "")
	if err == nil {
		if len(countries) > 6 {
			countries = countries[:6]
		}
		c.Data["FeaturedCountries"] = countries
	} else {
		c.Data["FeaturedCountries"] = nil
	}

	// Popular attractions — Paris as global default lat/lon
	attractions, err := services.GetAttractions(48.8566, 2.3522)
	if err == nil {
		if len(attractions) > 4 {
			attractions = attractions[:4]
		}
		c.Data["PopularAttractions"] = attractions
	} else {
		c.Data["PopularAttractions"] = nil
	}
}

// Search — AJAX only. GET /api/countries/search?q=
// Returns JSON [{name, capital, slug}], never reloads the page.
func (c *HomeController) Search() {
	q := strings.TrimSpace(c.GetString("q"))

	svc := &services.CountryService{}
	countries, err := svc.GetCountries(q, "")
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	type Suggestion struct {
		Name    string `json:"name"`
		Capital string `json:"capital"`
		Slug    string `json:"slug"`
	}

	limit := 8
	if len(countries) < limit {
		limit = len(countries)
	}

	out := make([]Suggestion, limit)
	for i, co := range countries[:limit] {
		out[i] = Suggestion{
			Name:    co.Name,
			Capital: co.Capital,
			Slug:    strings.ToLower(strings.ReplaceAll(co.Name, " ", "-")),
		}
	}

	c.Data["json"] = out
	c.ServeJSON()
}
