package router

import (
	"ship/internal/handler"
	"ship/internal/middleware"
	v1 "ship/internal/router/v1"

	"github.com/labstack/echo/v5"
)

func New(
	mw *middleware.Middleware,
	hdlr *handler.Handler,
) *echo.Echo {
	e := echo.New()

	apiGrp := e.Group("/api")

	e.Use(mw.Request.SetRequestID)

	v1.MountV1Router(apiGrp, hdlr)

	return e
}
