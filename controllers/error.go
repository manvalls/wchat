package controllers

import (
	"github.com/manvalls/wchat/templates"
	"github.com/manvalls/wit"
	"github.com/manvalls/wok"
	"github.com/manvalls/wq"
)

type errorController struct{ wok.Default }

func (e errorController) Plan() wok.Plan {
	return wok.List(
		wok.Handle(func(r wok.Request) wit.Command {
			r.SetStatusCode(404)
			return nil
		}),
		wq.Title.SetText("404"),
		wq.Body.Set(templates.Error()),
	)
}
