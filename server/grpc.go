package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/CuriousHet/Notify/data"
	"github.com/CuriousHet/Notify/notification"
	"github.com/CuriousHet/Notify/proto/postpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type postServer struct {
	postpb.UnimplementedPostServiceServer
	Queue *notification.Queue
}

func (s *postServer) PublishPost(ctx context.Context, req *postpb.Post) (*postpb.NotificationResponse, error) {
	author := req.AuthorId
	postID := req.PostId
	content := req.Content

	log.Printf("Received post from %s: [%s] %s\n", author, postID, content)

	followers, ok := data.Followers[author]
	if !ok || len(followers) == 0 {
		log.Printf("No followers found for user: %s\n", author)
		return &postpb.NotificationResponse{
			Message: fmt.Sprintf("Post received, but no followers for user: %s", author),
		}, nil
	}

	for _, follower := range followers {
		notif := notification.Notification{
			FollowerID: follower,
			AuthorID:   author,
			PostID:     postID,
			Content:    content,
		}
		s.Queue.Enqueue(notif)
	}

	return &postpb.NotificationResponse{
		Message: fmt.Sprintf("Post from %s sent to %d followers", author, len(followers)),
	}, nil
}

func StartGRPCServer(queue *notification.Queue) {
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	postpb.RegisterPostServiceServer(s, &postServer{Queue: queue})
	reflection.Register(s)

	log.Println("gRPC server running on port 5050...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
