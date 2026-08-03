package v1

import (
	"ship/internal/handler"

	"github.com/gin-gonic/gin"
)

func mountHealthRouter(
	routerGrp *gin.RouterGroup,
	handler *handler.HealthHandler,
) {
	routerGrp.GET("/health", handler.HealthCheck)
}
