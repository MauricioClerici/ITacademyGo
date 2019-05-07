package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/go-routine-test/src/api/controllers/goroutine"

)

const (
	port = ":8080"
)

var (
	router = gin.Default()

)



func main() {
	router.GET("/users/:id", goroutine.Controller)
	router.Run(port)

}