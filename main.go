package main

import (
	"github.com/fqzz2000/tiny-tictok/model"
	"github.com/fqzz2000/tiny-tictok/service"
	"github.com/gin-gonic/gin"
)
func main() {
	go service.RunMessageServer()
	model.Init_DB(false)
	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
