package handler

import (
	"github.com/labstack/echo"
)

func MyErrorHandler(err error, context echo.Context) {
	//code := http.StatusInternalServerError
	context.Logger().Error(err)
}
