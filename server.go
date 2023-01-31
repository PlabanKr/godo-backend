package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	var err error

	DB, err = gorm.Open("mysql", DbURL(BuildDBConfig()))
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer DB.Close()
	DB.AutoMigrate(&Todo{})

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/todos", GetTodos)
	router.POST("/todo", CreateATodo)
	router.GET("/todo/:id", GetATodo)
	router.PUT("/todo/:id", UpdateATodo)
	router.DELETE("/todo/:id", DeleteATodo)

	router.Run()
}
