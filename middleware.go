package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/api/auth/login5" {
			ctx.Next()
		} else {
			h := AuthHeader{}
			if err := ctx.ShouldBindHeader(&h); err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
			}

			if h.AuthorizationHeader == "ini_token" {
				ctx.Next()
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "token invalid",
				})
				ctx.Abort()
			}
		}
	}
}
