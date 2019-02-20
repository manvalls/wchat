package controllers

import (
	"net/http"

	"github.com/manvalls/wchat/routes"

	"github.com/manvalls/wchat/templates"
	"github.com/manvalls/wit"
	"github.com/manvalls/wok"
	"github.com/manvalls/wq"
)

type landing struct{ wok.Default }

func (l landing) Plan() wok.Plan {
	return wok.List(
		wq.Body.Set(templates.Landing()),
		wok.Post().Handle(func(r wok.Request) wit.Command {
			r.SetCookie(&http.Cookie{
				Name:  "name",
				Value: r.FormValue("name"),
			})

			r.URLRedirect(303, nil, routes.Chat)
			return nil
		}),
	)
}
