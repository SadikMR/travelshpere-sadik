package controllers

import (
	"github.com/SadikMR/travelshpere-sadik/services"
)

// WishlistController handles wishlist pages.
type WishlistController struct {
	BaseController
}

// Get renders the wishlist page.
func (c *WishlistController) Get() {
	c.Data["Wishlists"] = services.ListWishlists(c.GetUsername())
	c.Data["Title"] = "My Wishlist"
	c.Layout = "layouts/base.tpl"
	c.TplName = "wishlist/index.tpl"
}

// Rows renders the wishlist rows partial.
func (c *WishlistController) Rows() {
	c.Data["Wishlists"] = services.ListWishlists(c.GetUsername())
	c.TplName = "wishlist/rows.tpl"
}