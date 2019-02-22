package controllers

import (
	"github.com/manvalls/wchat/routes"
	"github.com/manvalls/wok"
)

type mainContainer struct{ wok.Default }

func (m mainContainer) Plan() wok.Plan {
	return nil
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
