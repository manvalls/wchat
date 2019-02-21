package selectors

import "github.com/manvalls/wq"

// Selectors for the app
var (
	MainContainer = wq.S(".container")
	Messages      = wq.S("#messages")
	MessageInput  = wq.S("input[name=message]")
	ChatTab       = wq.S("#chat-tab")
	LandingTab    = wq.S("#settings-tab")
	NavbarItems   = wq.S(".navbar-item")
	Errors        = wq.S(".error")
	ErrorMessage  = wq.S(".error-message")
	NameInput     = wq.S("input[name=name]")
)
