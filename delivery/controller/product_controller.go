package controller

import (
	"net/http"

	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/model"
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/repository"
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/usecase"
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/utils"
	"github.com/gin-gonic/gin"
)

func PostProduct(productRepo repository.ProductRepository, dest *gin.HandlerFunc) {
	productUseCase := usecase.NewProductUseCase(productRepo)
	*dest = func(ctx *gin.Context) {
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
	}
}

func GetProductById(productRepo repository.ProductRepository, dest *gin.HandlerFunc) {
	productUseCase := usecase.NewProductUseCase(productRepo)
	*dest = func(ctx *gin.Context) {
		var productSelected model.Product
		id := ctx.Param("id")
		err := productUseCase.GetProductById(id, &productSelected)
		utils.RaiseError(err)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    productSelected,
		})
	}
}

func GetProductAll(productRepo repository.ProductRepository, dest *gin.HandlerFunc) {
	productUseCase := usecase.NewProductUseCase(productRepo)
	*dest = func(ctx *gin.Context) {
		productList := productUseCase.GetAllProduct()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    productList,
		})
	}
}
