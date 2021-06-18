package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// https://github.com/gin-gonic/gin
func main() {
	router := gin.Default() // 생성
	router.GET("/get", get)
	router.GET("/user/:name/*action", getWithParameter)
	router.POST("/post", post)
	router.Run() // 실행
}

func get(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"123": "test",
	})
}

func getWithParameter(context *gin.Context) {
	name := context.Param("name")
	action := context.Param("action")
	message := name + " is " + action
	context.String(http.StatusOK, message)
}

func post(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"result": "good",
	})
}
