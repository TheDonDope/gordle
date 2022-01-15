package http

import (
	"github.com/gin-gonic/gin"
)

// NewServer ...
func NewServer() *gin.Engine {
	s := gin.Default()
	baseURL := s.Group("/gordle")
	baseURL.GET("healthy", Healthy)

	return s
}
