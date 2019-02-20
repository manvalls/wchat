package controllers

import (
	"github.com/manvalls/wchat/templates"
	"github.com/manvalls/wok"
	"github.com/manvalls/wq"
)

type chat struct{ wok.Default }

func (c chat) Plan() wok.Plan {
	return wok.List(
		wq.Body.Set(templates.Chat()),
	)
}
