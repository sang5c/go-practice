package server

import (
	"fmt"
	"net/http"
)

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "HELLO WORLD")
	})
	return mux
}

// run http server
//func main() {
//	http.ListenAndServe(":8080", MakeWebHandler())
//}
