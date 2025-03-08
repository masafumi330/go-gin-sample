package main

import (
	"go-gin-sample/app/adapter/dispatcher"
	"go-gin-sample/app/adapter/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	handler := handler.NewHandler()
	dispatcher.Register(e, handler)

	e.Run(":8080")
}
