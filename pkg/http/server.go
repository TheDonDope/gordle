package http

import (
	"github.com/gin-gonic/gin"
)

var (
	audience string
	domain   string
)

// CORSMiddleware returns the Handler Function for the CORS configuration.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// NewServer ...
func NewServer() *gin.Engine {
	setAuthVariables()
	r := gin.Default()
	r.Use(CORSMiddleware())

	baseURL := r.Group("/gordle")
	// authorized.Use(AuthRequired())
	baseURL.GET("healthy", Healthy)

	return r
}

// SetAuthVariables reads the environment variables
func setAuthVariables() {
	audience = "http://gordle-api" // FIXME: os.Getenv("AUTH0_API_IDENTIFIER")
	domain = "localhost:4242"      // FIXME: os.Getenv("AUTH0_DOMAIN")
}
