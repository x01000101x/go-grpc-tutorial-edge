package chat

import (
	"context"
	"log"

	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, err) {
	log.Printf("Received message body from the client %s", message.Body)
	return &Message{Body: "Hello from the server!"}, nil
}
