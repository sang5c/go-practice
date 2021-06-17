package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func PrintResponse(response *http.Response) {
	defer response.Body.Close()

	bytes, _ := ioutil.ReadAll(response.Body)
	str := string(bytes)
	fmt.Println(str)
}
