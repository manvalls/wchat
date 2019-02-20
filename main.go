package main

import (
	"log"
	"net/http"

	"github.com/manvalls/wchat/controllers"
	"github.com/manvalls/wchat/routes"
	"github.com/manvalls/wok"
)

//go:generate qtc -dir=./templates

func main() {
	http.Handle("/", wok.Handler{
		Router: routes.Router,
		Root: func() wok.Controller {
			return controllers.Root()
		},
	})

	log.Fatal(http.ListenAndServe(":8085", nil))
}
