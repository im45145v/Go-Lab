package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html",nil)
	})
	fmt.Println("Starting the server...")
	r.Run(":8081")
}
