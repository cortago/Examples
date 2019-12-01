package main

import (
	"Examples/basic-example/controller"

	"github.com/cortago/cortago/router"
)

func main() {
	router.Port = 8000

	y := router.New()
	y.FileServer("/assets")
	y.AppendRouter("/main", controller.Main)
	y.AppendRouter("/second", controller.Second)

	y.StartApp()
}
