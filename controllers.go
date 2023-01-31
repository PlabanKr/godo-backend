package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateATodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	err := CreateATodoInDB(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodos(c *gin.Context) {
	var todos []Todo
	err := GetAllTodosFromDB(&todos)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	err := GetATodoFromDB(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	err := GetATodoFromDB(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	err = UpdateATodoInDB(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	err := DeleteATodoFromDB(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
