package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type CalculateRequest struct {
	Num1     float64 `json:"num1"`
	Num2     float64 `json:"num2"`
	Operator string  `json:"operator"`
}

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the Calculator API")
	})
	

	r.POST("/calculate", func(c *gin.Context) {
		var req CalculateRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		var result float64

		switch req.Operator {
		case "+":
			result = req.Num1 + req.Num2
		case "-":
			result = req.Num1 - req.Num2
		case "*":
			result = req.Num1 * req.Num2
		case "/":
			if req.Num2 == 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Division by zero not allowed",
				})
				return
			}
			result = req.Num1 / req.Num2
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid operator",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	r.Run(":4000")
}

	