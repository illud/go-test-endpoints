package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Sum(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func GetPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "{'name': 'test','age': 22}",
	})
}

func main() {
	fmt.Println(Sum(2, 3))
	fmt.Println(Multiply(2, 3))

	SetupServer().Run()
}

func SetupServer() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", GetPing)
	r.GET("/users", GetUsers)

	return r
}
