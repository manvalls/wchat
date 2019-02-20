package selectors

import "github.com/manvalls/wq"

// Messages matches the chat messages container
var Messages = wq.S("#messages")

// MessageInput matches the message input
var MessageInput = wq.S("input[name=message]")
