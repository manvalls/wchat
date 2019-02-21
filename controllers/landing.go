package controllers

import (
	"net/http"

	"github.com/manvalls/wchat/routes"
	"github.com/manvalls/wchat/selectors"
	"github.com/manvalls/wq"

	"github.com/manvalls/wchat/templates"
	"github.com/manvalls/wit"
	"github.com/manvalls/wok"
)

type landing struct{ wok.Default }

func (l landing) Plan() wok.Plan {
	return wok.List(
		wq.Title.SetText("Settings"),
		selectors.NavbarItems.RmClass("is-active"),
		selectors.LandingTab.AddClass("is-active"),
		selectors.MainContainer.Set(templates.Landing()),
		wok.Post().Handle(func(r wok.Request) wit.Command {
			r.SetCookie(&http.Cookie{
				Name:  "name",
				Value: r.FormValue("name"),
			})

			r.URLRedirect(303, nil, routes.MainContainer, routes.Chat)
			return nil
		}),
	)
}
