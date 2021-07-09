package Get

import (
	"example/http/common"
	"fmt"
	"net/http"
	"net/url"
)

func DoGet() {
	simpleGet()
	urlEncode1()
	urlEncode2()
	getWithHeader()
}

func simpleGet() {
	response, err := http.Get("https://api.github.com")
	if err != nil {
		return
	}

	common.PrintResponse(response)
}

func urlEncode1() {
	params := url.Values{}
	params.Add("key1", "value1")
	params.Add("key2", "value2")
	fmt.Println(params.Encode())
}

func urlEncode2() {
	str := "key=" + url.QueryEscape("value")
	fmt.Println(str)
}

func getWithHeader() {
	request, err := http.NewRequest("GET", "https://api.github.com", nil)
	if err != nil {
		return
	}
	request.Header.Add("hello", "world")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}

	common.PrintResponse(response)
}
