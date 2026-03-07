package views

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func SendSuccess(c *gin.Context, code int, data any) {
	c.JSON(code, Response{
		Success: true,
		Data:    data,
	})
}

func SendError(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Success: false,
		Error:   message,
	})
}
