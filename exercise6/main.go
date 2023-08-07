package main

import (
	cartService "github.com/TinTran96/go23/exercise6/CartService"
	productService "github.com/TinTran96/go23/exercise6/ProductService"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/products")
	{
		client.GET("/", productService.GetAll)
		client.POST("/create", productService.Create)
		client.PATCH("/:id", productService.Update)
		client.DELETE("/:id", productService.Delete)
	}

	clientt := r.Group("/cart")
	{
		clientt.POST("/add", cartService.Add)
		clientt.DELETE("/remove", cartService.Delete)
		clientt.POST("/checkout", cartService.CheckOut)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":9888") // listen and serve on 0.0.0.0:8080
}
