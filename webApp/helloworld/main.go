package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!", r.URL.Path, "Love u...")
}

func main() {
	http.HandleFunc("/", handler)

	//创建路由
	http.ListenAndServe(":8080", nil)
}
