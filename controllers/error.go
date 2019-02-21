package controllers

import (
	"github.com/manvalls/wchat/templates"
	"github.com/manvalls/wok"
	"github.com/manvalls/wq"
)

type errorController struct{ wok.Default }

func (e errorController) Plan() wok.Plan {
	return wok.List(
		wq.Title.SetText("404"),
		wq.Body.Set(templates.Error()),
	)
}
