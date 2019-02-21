package controllers

import (
	"github.com/manvalls/wchat/routes"
	"github.com/manvalls/wchat/templates"
	"github.com/manvalls/wok"
	"github.com/manvalls/wq"
)

type mainContainer struct{ wok.Default }

func (m mainContainer) Plan() wok.Plan {
	return wq.Body.Set(templates.MainContainer())
}

func (m mainContainer) Resolve(route string) wok.Controller {
	switch route {
	case routes.Chat:
		return chat{}
	case routes.Landing:
		return landing{}
	}

	return wok.Default{}
}
