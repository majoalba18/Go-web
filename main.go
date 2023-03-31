package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productos struct{
	Id 				int 	`json:"id"`
	Name 			string 	`json:"name"`
	Quantity 		int 	`json:"quantity"`
	Code_value 		string 	`json:"code_value"`
	Is_published 	bool 	`json:"is_published"`
	Expiration 		string 	`json:"expiration"`
	Price 			float64	`json:"price"`
}

func main() {

	file, err := ioutil.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	var producto []productos

	err = json.Unmarshal(file, &producto)
	if err != nil {
		log.Fatal(err)
		return
	}

	router := gin.Default()
	router.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, producto)
	})
	router.GET("/product/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid product ID",
			})
			return
		}

		for _, p := range producto {
			if p.Id == id {
				ctx.JSON(200, p)
				return
			}
		}

		ctx.JSON(404, gin.H{
			"error": "Product not found",
		})
	})
	router.POST("/products/search", func(ctx *gin.Context) {
		var searchProduct struct {
			Price float64 `json:"price"`
		}

		err := ctx.BindJSON(&searchProduct)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid search query",
			})
			return
		}

		var result []productos

		for _, p := range producto {
			if p.Price >= searchProduct.Price {
				result = append(result, p)
			}
		}
		ctx.JSON(200,result)
	})


	router.Run(":8080")
}


