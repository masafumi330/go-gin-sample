package dispatcher

import (
	"go-gin-sample/app/adapter/handler"

	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine, h *handler.Handler) {
	user := e.Group("/user")
	{
		user.GET("/:id", h.Get)
		user.POST("", h.Create)
		user.PUT("/:id", h.Update)
		user.DELETE("/:id", h.Delete)
	}

	// add more endpoints below
}
