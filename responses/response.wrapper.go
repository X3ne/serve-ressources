package responses

import (
	"github.com/labstack/echo/v4"
)

type Error struct {
	Code		int		`json:"code"`
	Message	string	`json:"message"`
}

type Success struct {
	Success	bool		`json:"success"`
	Data		interface{}	`json:"data"`
}

func Response(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, data)
}

func ErrorResponse(c echo.Context, status int, err error) error {
	return c.JSON(status, Error{
		Code:		status,
		Message:	err.Error(),
	})
}

func SuccessResponse(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, Success{
		Success: true,
		Data: data,
	})
}
