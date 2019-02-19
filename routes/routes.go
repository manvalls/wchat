package routes

import (
	"github.com/manvalls/way"
)

// Route constants
const (
	Chat    = "chat"
	Landing = "landing"
	Error   = "error"
)

// Router for the app
var Router = way.BuildRouter(way.RouteMap{
	"/":      {Landing},
	"/chat":  {Chat},
	":path*": {Error},
})
