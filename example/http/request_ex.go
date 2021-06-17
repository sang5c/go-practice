package main

import (
	Get "example/http/get"
	"example/http/post"
)

func main() {
	Get.DoGet()
	Post.DoPost()

	jsonHandling()
}
