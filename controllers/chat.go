package controllers

import (
	"sync"

	"github.com/manvalls/channel"
	"github.com/manvalls/wchat/types"
	"github.com/manvalls/wok"
)

var (
	messages      = []types.Message{}
	messagesMutex = sync.RWMutex{}
	chatChannel   = channel.NewChannel()
)

type chat struct{ wok.Default }

func (c chat) Plan() wok.Plan {
	return nil
}
