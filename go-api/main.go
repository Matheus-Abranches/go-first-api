package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// camada repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// camada usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	//camada controller
	productController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/products", productController.CreateProduct)
	server.GET("/products/:productId", productController.GetProductById)

	server.Run(":8080")

}
