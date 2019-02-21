package routes

import (
	"github.com/manvalls/way"
)

// Route constants
const (
	MainContainer = "main"
	Chat          = "chat"
	Landing       = "landing"
	Error         = "error"
)

// Router for the app
var Router = way.BuildRouter(way.RouteMap{
	"/":      {MainContainer, Landing},
	"/chat":  {MainContainer, Chat},
	":path*": {Error},
})
