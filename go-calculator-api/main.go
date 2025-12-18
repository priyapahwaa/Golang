package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/add", func(c *gin.Context) {
		a, b, ok := getNumbers(c)
		if !ok {
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": a + b})
	})

	r.GET("/sub", func(c *gin.Context) {
		a, b, ok := getNumbers(c)
		if !ok {
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": a - b})
	})

	r.GET("/mul", func(c *gin.Context) {
		a, b, ok := getNumbers(c)
		if !ok {
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": a * b})
	})

	r.GET("/div", func(c *gin.Context) {
		a, b, ok := getNumbers(c)
		if !ok {
			return
		}
		if b == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "division by zero"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": a / b})
	})

	r.Run(":8080")
}

func getNumbers(c *gin.Context) (float64, float64, bool) {
	aStr := c.Query("a")
	bStr := c.Query("b")

	a, err1 := strconv.ParseFloat(aStr, 64)
	b, err2 := strconv.ParseFloat(bStr, 64)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid numbers"})
		return 0, 0, false
	}

	return a, b, true
}

type Student struct {
    name string
    age  int
    cgpa float64
}
