package main

import (
	"goGettingStarted/controller"
	"net/http"
)

func main() {

	controller.RegisterControllers()
	http.ListenAndServe(":8081", nil)
}
