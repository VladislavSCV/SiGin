package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/VladislavSCV/SiGin"
	"github.com/VladislavSCV/SiGin/Models"
)



func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("[%s] %s %s %s", c.Request.Method, c.Request.URL.Path, c.Request.RemoteAddr, time.Since(start))
	}
}


func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	api := r.Group("/api")
	api.Use(Logger())

	r.GET("/", Index)

	api.GET("/ping", Ping)
	
	// api.GET("/users", GetUsers)
	// api.GET("/users/:id", GetUser)
	// api.POST("/users", CreateUser)
	// api.PUT("/users/:id", UpdateUser)
	// api.DELETE("/users/:id", DeleteUser)

	// api.GET("/autos", GetAutos)
	// api.GET("/autos/:id", GetAuto)
	// api.POST("/autos", CreateAuto)
	// api.PUT("/autos/:id", UpdateAuto)
	// api.DELETE("/autos/:id", DeleteAuto)

	r.Run(":8000");
}


func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "TITLE",
	})
}

func Ping(c *gin.Context) {
	data := map[string]interface{}{
		"message": "pong",
	}
	c.AsciiJSON(http.StatusOK, data)
}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Api GetUsers
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
// func GetUsers(c *gin.Context) {

// }