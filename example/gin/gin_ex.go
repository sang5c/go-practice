package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// https://github.com/gin-gonic/gin
func main() {
	r := gin.Default() // 생성
	r.GET("/get", get)
	r.POST("/post", post)
	r.Run() // 실행
}

func post(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"result": "good",
	})
}

func get(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"123": "test",
	})
}
