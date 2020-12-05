package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println(path)

		name := context.DefaultQuery("name", "hello")
		fmt.Println(name)

		context.Writer.Write([]byte("Hello," + name))

	})

	engine.Run()
}
