package main

import (
	"context"
	"log"
	"net"

	"github.com/rodrigodelmonte/grpc-assistant/pkg/assistant"
	"github.com/rodrigodelmonte/grpc-assistant/pkg/llm"
	"google.golang.org/grpc"
)

// assistantServer implements the assistant server.
type assistantServer struct {
	assistant.UnimplementedAssistantServer
}

// Handle is the handler for the assistant server.
func (s *assistantServer) Handle(ctx context.Context, in *assistant.AssistantRequest) (*assistant.AssistantResponse, error) {

	res, err := llm.Ask(in.Request)
	if err != nil {
		log.Fatalf("Could not ask %s", err)
	}
	return &assistant.AssistantResponse{Result: res}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Could not listen %s", err)
	}
	log.Println("Listening on :8080")

	svc := grpc.NewServer()
	srv := &assistantServer{}

	assistant.RegisterAssistantServer(svc, srv)

	err = svc.Serve(listener)
	if err != nil {
		log.Fatalf("Could not serve %s", err)
	}
}
