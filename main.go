package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
	
	api.GET("/users", GetUsers)
	api.GET("/users/:id", GetUser)
	api.POST("/addUser", AddUser)
	api.PUT("/users/:id/:name", UpdateUser)
	api.DELETE("/users/:id", DeleteUser)

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


func GetUsers(c *gin.Context) {
	users := models.GetUsers()
	for _, user := range users {
		c.JSON(http.StatusOK, user)
	}
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print("Warning")
	}
	user := models.GetUserById(id)
	if user.Id == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	data := map[string]interface{}{
		"id":       user.Id,
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	}
	c.AsciiJSON(http.StatusOK, data)
}

func AddUser(c *gin.Context) {
	if models.UsersDB == nil {
		log.Println("UsersDB is nil")
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	id := int(len(models.UsersDB) + 1)
	user := models.User{id, "name", "email", "passw"}
	models.UsersDB[id] = user
	c.String(http.StatusOK, "OK")
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	name := c.Param("name")
	if err != nil {
		log.Print("Error updating user")
	}
	models.UpdateUserById(id, name)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print("Warning")
	}
	models.DeleteUser(id)
	c.JSON(http.StatusOK, "OK")
}
