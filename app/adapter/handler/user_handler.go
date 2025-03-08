package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type implUserHandler struct{}

func NewUserHandler() UserHandler {
	return &implUserHandler{}
}

func (h *implUserHandler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "Bob", "age": 20})
}

func (h *implUserHandler) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

func (h *implUserHandler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (h *implUserHandler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
