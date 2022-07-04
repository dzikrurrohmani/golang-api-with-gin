package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	routerEngine := gin.Default()
	routerEngine.GET("/", func(c *gin.Context) {
		c.String(200, "Health Check")
	})

	routerEngine.GET("/greeting1/:name", greeting1)
	routerEngine.GET("/greeting2/:name", greeting2)

	routerEngine.POST("/login1", login1)
	routerEngine.POST("/login2", login2)
	routerEngine.POST("/login3", login3)
	routerEngine.POST("/login4", login4)
	routerEngine.POST("/login5", login5)
	err := routerEngine.Run() // secara default menggunakan port :8080
	// err := routerEngine.Run(:8888) // ganti port
	if err != nil {
		panic(err)
	}
}

func greeting1(c *gin.Context) {
	name := c.Param("name")
	// c.String(200, "Greeting %s", name)
	c.String(http.StatusOK, "Greeting %s", name)
}

// Query param / String in path -> ?key=value

func greeting2(c *gin.Context) {
	name := c.Param("name")
	kec := c.Query("kec")
	kel := c.Query("kel")
	// c.String(200, "Greeting %s", name)
	c.String(http.StatusOK, "Hello %s di %s, %s", name, kec, kel)
}

func login1(c *gin.Context) {
	username := c.PostForm("username")
	c.PostForm("password")
	c.String(http.StatusOK, "Hello %s", username)
}

func login2(c *gin.Context) {
	var uc UserCredential
	if c.ShouldBind(&uc) == nil {
		// c.String(http.StatusOK, "Hello %s", uc.Username)
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    uc.Username,
		})
	}
}

func login3(c *gin.Context) {
	var uc UserCredential
	if err := c.ShouldBind(&uc); err != nil {
		// c.String(http.StatusOK, "Hello %s", uc.Username)
		fmt.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    uc.Username,
		})
	}
}

func login4(c *gin.Context) {
	var uc UserCredential
	if err := c.ShouldBind(&uc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    uc.Username,
		})
	}
}

func login5(c *gin.Context) {
	var uc UserCredential
	if err := c.ShouldBindJSON(&uc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    uc.Username,
		})
	}
}
