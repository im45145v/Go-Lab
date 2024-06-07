package main

import (
	"coffeeshop/coffee"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the coffeeshop!",
		})
	})
	r.GET("/home", coffeelist)
	r.GET("/coffee", getcoffee)
	r.Run(":8081")
}
func getcoffee(c *gin.Context) {
	coffeelist, _ := coffee.GetCoffees()
	c.String(http.StatusOK, "%s", coffeelist)
}
func coffeelist(c *gin.Context) {
	coffee, _ := coffee.GetCoffees()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"list": coffee.List,
	})
}
