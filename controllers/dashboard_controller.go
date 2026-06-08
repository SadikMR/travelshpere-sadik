package controllers

import "github.com/SadikMR/travelshpere-sadik/services"

// DashboardController handles the dashboard page.
type DashboardController struct {
	BaseController
}

// Get renders the dashboard page with summary and destinations.
func (c *DashboardController) Get() {
	username := c.GetUsername()

	c.Data["Summary"] = services.GetDashboardSummary(username)
	c.Data["Destinations"] = services.ListWishlists(username)
	c.Data["Title"] = "Travel Dashboard"
	c.Layout = "layouts/base.tpl"
	c.TplName = "dashboard.tpl"
}
