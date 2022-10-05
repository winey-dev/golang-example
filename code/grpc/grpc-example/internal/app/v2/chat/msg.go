package chat

import (
	"errors"
	"log"

	chatpb "github.com/yiaw/grpc-example/protos/v2/chat"
)

type RecvStream struct {
	msg *chatpb.Message
	err error
}

type Channel struct {
	name    string
	channel chan *chatpb.Message
}

var channel map[string]*Channel

type rc struct{}

func init() {
	channel = make(map[string]*Channel)
}
func Client() *rc {
	return &rc{}
}

func (r *rc) InitChannel(name string) (chan *chatpb.Message, error) {
	if len(name) == 0 {
		return nil, errors.New("required user name")
	}

	c := &Channel{
		name:    name,
		channel: make(chan *chatpb.Message),
	}

	channel[name] = c
	return c.channel, nil
}

func (r *rc) WriteChannel(msg *chatpb.Message) error {
	for user, c := range channel {
		log.Printf("send to %s\n", user)
		c.channel <- msg
	}
	return nil
}
