package api

import (
	"net/http"

	"github.com/SadikMR/travelshpere-sadik/services"
	"github.com/SadikMR/travelshpere-sadik/utils/responses"

	beego "github.com/beego/beego/v2/server/web"
)

// AttractionController handles attraction API requests.
type AttractionController struct {
	beego.Controller
}

// Get returns attractions by lat/lon coordinates.
func (c *AttractionController) Get() {
	lat, errLat := c.GetFloat("lat")
	lon, errLon := c.GetFloat("lon")

	if errLat != nil || errLon != nil {
		responses.WriteError(
			&c.Controller,
			http.StatusBadRequest,
			"lat and lon query parameters are required",
		)
		return
	}

	attractions, err := services.GetAttractions(lat, lon)
	if err != nil {
		responses.WriteError(
			&c.Controller,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	c.Data["json"] = attractions
	c.ServeJSON()
}
