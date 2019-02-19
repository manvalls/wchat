package controllers

import (
	"github.com/manvalls/wchat/routes"
	"github.com/manvalls/wchat/templates"
	"github.com/manvalls/wok"
	"github.com/manvalls/wq"
)

// Root controller of the page
func Root() wok.Controller {
	return root{}
}

type root struct{ wok.Default }

func (r root) Plan() wok.Plan {
	return wq.Head.Set(templates.Head())
}

func (r root) Resolve(route string) wok.Controller {
	switch route {
	case routes.Chat:
		return chat{}
	case routes.Landing:
		return landing{}
	}

	return errorController{}
}
