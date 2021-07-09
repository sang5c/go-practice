package Post

import (
	"bytes"
	"encoding/json"
	"example/http/common"
	"net/http"
	"net/url"
)

type User struct {
	Name string
	Age  int
}

const requestUrl = "http://localhost:8080/post"

func DoPost() {
	simplePostJson()
	simplePostForm()
}

func simplePostForm() {
	values := url.Values{"Name": {"tester2"}, "Age": {"11"}}
	response, _ := http.PostForm(requestUrl, values)

	common.PrintResponse(response)
}

func simplePostJson() {
	contentType := "application/json"
	body, _ := json.Marshal(User{"tester", 29})
	buffer := bytes.NewBuffer(body)
	response, _ := http.Post(requestUrl, contentType, buffer)
	common.PrintResponse(response)
}

func getJsonString(v interface{}) string {
	marshal, _ := json.Marshal(v)
	return string(marshal)
}
