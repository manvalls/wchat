package controllers

import (
	"github.com/manvalls/wok"
)

type landing struct{ wok.Default }

func (l landing) Plan() wok.Plan {
	return nil
}
