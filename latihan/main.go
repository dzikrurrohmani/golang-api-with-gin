package main

import (
	"net/http"

	"github.com/dzikrurrohmani/golang-api-with-gin/delivery/controller"
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/model"
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/repository"
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/usecase"
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	productRepo := repository.NewProductRepository()
	categoryRepo := repository.NewCategoryRepository()
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	routerEngine := gin.Default()
	routerGroup := routerEngine.Group("/api")

	var PostProductFn, GetProductAllFn, GetProductByIdFn gin.HandlerFunc
	controller.PostProduct(productRepo, &PostProductFn)
	controller.GetProductAll(productRepo, &PostProductFn)
	controller.GetProductById(productRepo, &PostProductFn)
	routerGroup.POST("/product", PostProductFn)
	routerGroup.GET("/product", GetProductAllFn)
	routerGroup.GET("/product/:id", GetProductByIdFn)

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
