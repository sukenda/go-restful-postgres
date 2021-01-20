package exception

import (
	"github.com/labstack/echo/v4"
	"github.com/sukenda/go-restful-postgres/model"
	"net/http"
)

func ErrorHandler(ctx echo.Context, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return ctx.JSON(http.StatusInternalServerError, model.WebResponse{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
