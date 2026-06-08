package api

import (
	"encoding/json"
	"net/http"

	"github.com/SadikMR/travelshpere-sadik/services"

	beego "github.com/beego/beego/v2/server/web"
)

// WishlistController handles wishlist API requests.
type WishlistController struct {
	beego.Controller
}

// getUserID safely extracts the username from the session.
// Returns empty string and aborts with 401 if not authenticated.
func (c *WishlistController) getUserID() (string, bool) {
	session := c.GetSession("username")
	if session == nil {
		c.CustomAbort(http.StatusUnauthorized, "login required")
		return "", false
	}

	userID, ok := session.(string)
	if !ok || userID == "" {
		c.CustomAbort(http.StatusUnauthorized, "login required")
		return "", false
	}

	return userID, true
}

// Get returns the authenticated user's wishlist.
func (c *WishlistController) Get() {
	userID, ok := c.getUserID()
	if !ok {
		return
	}

	c.Data["json"] = services.ListWishlists(userID)
	c.ServeJSON()
}

// Post creates a wishlist entry.
func (c *WishlistController) Post() {
	userID, ok := c.getUserID()
	if !ok {
		return
	}

	var payload struct {
		CountryName string `json:"country_name"`
		Note        string `json:"note"`
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload); err != nil {
		c.CustomAbort(http.StatusBadRequest, "invalid payload")
		return
	}

	if payload.CountryName == "" {
		c.CustomAbort(http.StatusBadRequest, "country name is required")
		return
	}

	wishlist := services.CreateWishlist(
		userID,
		payload.CountryName,
		payload.Note,
	)

	c.Ctx.Output.SetStatus(http.StatusCreated)

	c.Data["json"] = wishlist
	c.ServeJSON()
}

// Put updates a wishlist entry.
func (c *WishlistController) Put() {
	userID, ok := c.getUserID()
	if !ok {
		return
	}

	id, err := c.GetInt(":id")
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, "invalid id")
		return
	}

	var payload struct {
		Note   string `json:"note"`
		Status string `json:"status"`
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload); err != nil {
		c.CustomAbort(http.StatusBadRequest, "invalid payload")
		return
	}

	wishlist, err := services.UpdateWishlist(
		userID,
		id,
		payload.Note,
		payload.Status,
	)

	if err != nil {
		switch err {
		case services.ErrWishlistNotFound:
			c.CustomAbort(http.StatusNotFound, err.Error())

		case services.ErrForbidden:
			c.CustomAbort(http.StatusForbidden, err.Error())

		case services.ErrInvalidStatus:
			c.CustomAbort(http.StatusBadRequest, err.Error())
		}

		return
	}

	c.Data["json"] = wishlist
	c.ServeJSON()
}

// Delete removes a wishlist entry.
func (c *WishlistController) Delete() {
	userID, ok := c.getUserID()
	if !ok {
		return
	}

	id, err := c.GetInt(":id")
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, "invalid id")
		return
	}

	if err := services.DeleteWishlist(userID, id); err != nil {
		switch err {
		case services.ErrWishlistNotFound:
			c.CustomAbort(http.StatusNotFound, err.Error())

		case services.ErrForbidden:
			c.CustomAbort(http.StatusForbidden, err.Error())
		}

		return
	}

	c.Ctx.Output.SetStatus(http.StatusNoContent)
}