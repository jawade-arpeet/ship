package middleware

import (
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type RequestMiddleware struct{}

func newRequestMiddleware() *RequestMiddleware {
	return &RequestMiddleware{}
}

func (m *RequestMiddleware) SetRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqID := ctx.GetHeader("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
			ctx.Request.Header.Set("X-Request-ID", reqID)
		}

		ctx.Header("X-Request-ID", reqID)
		ctx.Set("request_id", reqID)

		ctx.Next()
	}
}
