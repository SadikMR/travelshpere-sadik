package responses

import (
	beego "github.com/beego/beego/v2/server/web"
)

// ErrorResponse represents a standard API error response.
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// WriteError sends a JSON error response.
func WriteError(
	c *beego.Controller,
	status int,
	message string,
) {
	c.Ctx.Output.SetStatus(status)

	c.Data["json"] = ErrorResponse{
		Status:  status,
		Message: message,
	}

	c.ServeJSON()
}
