package api

import (
	"github.com/SadikMR/travelshpere-sadik/services"
)

// DashboardController handles dashboard API requests.
type DashboardController struct {
	BaseAPIController
}

// Summary returns dashboard counters as JSON for AJAX refresh.
func (c *DashboardController) Summary() {
	c.Data["json"] = services.GetDashboardSummary(c.Username)
	c.ServeJSON()
}