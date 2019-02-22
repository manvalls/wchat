package controllers

import (
	"github.com/manvalls/wchat/routes"
	"github.com/manvalls/wok"
)

// Root controller of the page
func Root() wok.Controller {
	return root{}
}

type root struct{ wok.Default }

func (r root) Plan() wok.Plan {
	return nil
}

func (r root) Resolve(route string) wok.Controller {
	switch route {
	case routes.MainContainer:
		return mainContainer{}
	case routes.Error:
		return errorController{}
	}

	return wok.Default{}
}
