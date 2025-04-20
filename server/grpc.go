// server/grpc.go
package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/CuriousHet/Notify/proto/postpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type PostServer struct {
	postpb.UnimplementedPostServiceServer
}

func (s *PostServer) PublishPost(ctx context.Context, post *postpb.Post) (*postpb.NotificationResponse, error) {
	log.Printf("Received Post from %s: %s", post.AuthorId, post.Content)

	// Later we'll create and dispatch notifications here
	return &postpb.NotificationResponse{
		Message: fmt.Sprintf("Post %s published successfully!", post.PostId),
	}, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	postpb.RegisterPostServiceServer(s, &PostServer{})

	reflection.Register(s)

	log.Println("gRPC server running at :5050")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
