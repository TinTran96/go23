package CartService

import (
	"fmt"

	productService "github.com/TinTran96/go23/exercise6/ProductService"
	"github.com/gin-gonic/gin"
)

type Item struct {
	Product  productService.Product `json:"product"`
	Quantity int                    `json:"quantity"`
	Price    float32                `json:"price"`
}

var carts []Item

func RemoveIndex(s []Item, index int) []Item {
	return append(s[:index], s[index+1:]...)
}

func Add(c *gin.Context) {
	var products = productService.GetProducts()

	type AddParam struct {
		Id       int `form:"id" json:"id" binding:"required"`
		Quantity int `form:"quantity" json:"quantity" binding:"required"`
	}

	var json AddParam

	if err := c.ShouldBindJSON(&json); err == nil {
		for i := 0; i < len(products); i++ {
			if json.Id == products[i].ID {
				var item Item
				item.Product = products[i]
				item.Quantity = json.Quantity
				item.Price = products[i].Price * float32(json.Quantity)
				carts = append(carts, item)
			}
		}
		c.JSON(200, gin.H{
			"messages": "Added",
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}

}

func CheckOut(c *gin.Context) {
	var totalPrice float32 = 0
	for i := 0; i < len(carts); i++ {
		totalPrice += carts[i].Product.Price
	}
	s := fmt.Sprintf("Check out complete, Total Price: %.2f", totalPrice)
	c.JSON(200, gin.H{
		"messages": s,
	})
	carts = nil

}

func Delete(c *gin.Context) {
	type DeleteParam struct {
		Id int `form:"id" json:"id" binding:"required"`
	}

	var json DeleteParam

	if err := c.ShouldBindJSON(&json); err == nil {
		for i := 0; i < len(carts); i++ {
			if carts[i].Product.ID == json.Id {
				carts = RemoveIndex(carts, i)
				break
			}
		}
		c.JSON(200, gin.H{
			"messages": "Deleted",
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}
