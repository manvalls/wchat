package controllers

import (
	"github.com/manvalls/wok"
)

type errorController struct{ wok.Default }

func (e errorController) Plan() wok.Plan {
	return nil
}
