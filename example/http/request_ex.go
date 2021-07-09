package main

import (
	Get2 "example/http/client/get"
	"example/http/client/post"
)

func main() {
	Get2.DoGet()
	Post.DoPost()

	jsonHandling()
}
