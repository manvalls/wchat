package controllers

import (
	"sync"

	"github.com/manvalls/channel"
	"github.com/manvalls/wchat/selectors"
	"github.com/manvalls/wchat/templates"
	"github.com/manvalls/wchat/types"
	"github.com/manvalls/wit"
	"github.com/manvalls/wok"
	"github.com/manvalls/wq"
)

var (
	messages      = []types.Message{}
	messagesMutex = sync.RWMutex{}
	chatChannel   = channel.NewChannel()
)

type chat struct{ wok.Default }

func (c chat) Plan() wok.Plan {
	return wok.List(
		wok.Socket().Do(func(r wok.ReadOnlyRequest) {
			chatChannel.Join(r)
		}),
		wok.Sync().Post().Handle(func(r wok.Request) wit.Command {
			messagesMutex.Lock()
			defer messagesMutex.Unlock()

			msg := types.Message{
				User:    "Anon",
				Message: r.FormValue("message"),
			}

			userCookie, err := r.Cookie("name")
			if err == nil {
				msg.User = userCookie.Value
			}

			chatChannel.Broadcast(
				selectors.Messages.Append(templates.Message(msg)),
			)

			messages = append(messages, msg)
			return selectors.MessageInput.AddAttr(map[string]string{
				"value": "",
			})
		}),
		wok.Run(func(r wok.Request) wit.Command {
			messagesMutex.RLock()
			defer messagesMutex.RUnlock()
			return wq.Body.Set(templates.Chat(messages))
		}),
	)
}
