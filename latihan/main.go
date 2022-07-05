package main

import (
	"net/http"

	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/model"
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/repository"
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/usecase"
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	productRepo := repository.NewProductRepository()
	productUseCase := usecase.NewProductUseCase(productRepo)
	categoryRepo := repository.NewCategoryRepository()
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	routerEngine := gin.Default()
	routerGroup := routerEngine.Group("/api")

	routerGroup.POST("/product", func(ctx *gin.Context) {
		var newProduct model.Product
		if err := ctx.BindJSON(&newProduct); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
			err := productUseCase.AddNewProduct(&newProduct)
			utils.RaiseError(err)
		}
	})
	routerGroup.GET("/product", func(ctx *gin.Context) {
		productList := productUseCase.GetAllProduct()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    productList,
		})
	})
	routerGroup.GET("/product/:id", func(ctx *gin.Context) {
		var productSelected model.Product
		id := ctx.Param("id")
		err := productUseCase.GetProductById(id, &productSelected)
		utils.RaiseError(err)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    productSelected,
		})
	})
	routerGroup.GET("/product/:id", func(ctx *gin.Context) {
		var productSelected model.Product
		id := ctx.Param("id")
		err := productUseCase.GetProductById(id, &productSelected)
		utils.RaiseError(err)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    productSelected,
		})
	})

	routerGroup.POST("/category", func(ctx *gin.Context) {
		var newCategory model.Category
		if err := ctx.BindJSON(&newCategory); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
			err := categoryUseCase.AddNewCategory(&newCategory)
			utils.RaiseError(err)
		}
	})
	routerGroup.GET("/category", func(ctx *gin.Context) {
		categoryList := categoryUseCase.GetAllCategory()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    categoryList,
		})
	})

	err := routerEngine.Run(":8888") // ganti port
	if err != nil {
		panic(err)
	}
}

/*
{
	"id": "P001",
	"name": "jus",
	"category": {
	"id": "C001",
	"name": "drink"
	},
	"is_active": true
}
*/
