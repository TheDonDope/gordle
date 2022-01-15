package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthy returns HTTP 200 on hit of the root route
func Healthy(c *gin.Context) {
	c.String(http.StatusOK, "%s", "Hello")
}
