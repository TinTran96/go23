package ProductService

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

var products []Product

func RemoveIndex(s []Product, index int) []Product {
	return append(s[:index], s[index+1:]...)
}

func Create(c *gin.Context) {

	type CreateProduct struct {
		Name        string  `form:"name" json:"name" binding:"required"`
		Description string  `form:"description" json:"description" binding:"required"`
		Price       float32 `form:"price" json:"price" binding:"required"`
	}

	var json CreateProduct

	if err := c.ShouldBindJSON(&json); err == nil {
		var index = len(products) + 1
		var newProduct Product
		newProduct.ID = index
		newProduct.Name = json.Name
		newProduct.Description = json.Description
		newProduct.Price = json.Price
		products = append(products, newProduct)
		s := fmt.Sprintf("%s - %s - %f | Inserted", newProduct.Name, newProduct.Description, newProduct.Price)
		c.JSON(200, gin.H{
			"messages": s,
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}

}

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		type UpdateProduct struct {
			Name        string  `form:"name" json:"name" binding:"required"`
			Description string  `form:"description" json:"description" binding:"required"`
			Price       float32 `form:"price" json:"price" binding:"required"`
		}

		var json UpdateProduct
		if err := c.ShouldBindJSON(&json); err == nil {
			for i := 0; i < len(products); i++ {
				if products[i].ID == int(id) {
					products[i].Name = json.Name
					products[i].Description = json.Description
					products[i].Price = json.Price
					break
				}
			}
			s := fmt.Sprintf("%s - %s - %f | Updated", json.Name, json.Description, json.Price)
			c.JSON(200, gin.H{
				"messages": s,
			})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(500, gin.H{"error": "Wrong ID Format"})
	}

}

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		for i := 0; i < len(products); i++ {
			if products[i].ID == id {
				products = RemoveIndex(products, i)
				break
			}
		}

		c.JSON(200, gin.H{
			"messages": "Deleted",
		})
	} else {
		c.JSON(500, gin.H{"error": "Wrong ID Format"})
	}
}

func GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": GetProducts(),
	})
}

func GetProducts() []Product {
	return products
}
