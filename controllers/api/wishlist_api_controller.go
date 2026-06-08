package api

import (
	"encoding/json"
	"net/http"

	"github.com/SadikMR/travelshpere-sadik/services"
	"github.com/SadikMR/travelshpere-sadik/utils/responses"
)

// WishlistController handles wishlist API requests.
type WishlistController struct {
	BaseAPIController
}

// Get returns the authenticated user's wishlist.
func (c *WishlistController) Get() {
	c.Data["json"] = services.ListWishlists(c.Username)
	c.ServeJSON()
}

// Post creates a wishlist entry.
func (c *WishlistController) Post() {
	var payload struct {
		CountryName string `json:"country_name"`
		Note        string `json:"note"`
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload); err != nil {
		responses.WriteError(&c.Controller, http.StatusBadRequest, "invalid payload")
		return
	}

	wishlist, err := services.CreateWishlist(
		c.Username,
		payload.CountryName,
		payload.Note,
	)
	if err != nil {
		responses.WriteError(&c.Controller, http.StatusBadRequest, err.Error())
		return
	}

	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = wishlist
	c.ServeJSON()
}

// Put updates a wishlist entry.
func (c *WishlistController) Put() {
	id, err := c.GetInt(":id")
	if err != nil {
		responses.WriteError(&c.Controller, http.StatusBadRequest, "invalid id")
		return
	}

	var payload struct {
		Note   string `json:"note"`
		Status string `json:"status"`
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload); err != nil {
		responses.WriteError(&c.Controller, http.StatusBadRequest, "invalid payload")
		return
	}

	wishlist, err := services.UpdateWishlist(
		c.Username,
		id,
		payload.Note,
		payload.Status,
	)

	if err != nil {
		switch err {
		case services.ErrWishlistNotFound:
			responses.WriteError(&c.Controller, http.StatusNotFound, err.Error())
		case services.ErrForbidden:
			responses.WriteError(&c.Controller, http.StatusForbidden, err.Error())
		case services.ErrInvalidStatus:
			responses.WriteError(&c.Controller, http.StatusBadRequest, err.Error())
		default:
			responses.WriteError(&c.Controller, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.Data["json"] = wishlist
	c.ServeJSON()
}

// Delete removes a wishlist entry.
func (c *WishlistController) Delete() {
	id, err := c.GetInt(":id")
	if err != nil {
		responses.WriteError(&c.Controller, http.StatusBadRequest, "invalid id")
		return
	}

	if err := services.DeleteWishlist(c.Username, id); err != nil {
		switch err {
		case services.ErrWishlistNotFound:
			responses.WriteError(&c.Controller, http.StatusNotFound, err.Error())
		case services.ErrForbidden:
			responses.WriteError(&c.Controller, http.StatusForbidden, err.Error())
		default:
			responses.WriteError(&c.Controller, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.Ctx.Output.SetStatus(http.StatusNoContent)
}