package router

import (
	"ship/internal/constants"
	"ship/internal/handler"
	"ship/internal/middleware"
	v1 "ship/internal/router/v1"

	"github.com/gin-gonic/gin"
)

func New(
	runEnv constants.Env,
	middleware *middleware.Middleware,
	handler *handler.Handler,
) *gin.Engine {
	if runEnv == constants.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Request.SetRequestID())

	apiGrp := router.Group("/api")

	v1.MountV1Router(apiGrp, middleware, handler)

	return router
}
