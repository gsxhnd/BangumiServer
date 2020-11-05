package middleware

import "github.com/gin-gonic/gin"

func NotFoundMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() == 404 {
			c.JSON(404, struct {
				Message string `json:"message"`
			}{
				Message: "The incorrect API route.",
			})
		}
	}
}
