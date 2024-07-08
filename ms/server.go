package ms

import (
	"context"
	pb "dst-admin-go/proto/messaging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"net"
	"sync"
)

var Server *server

func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := ss.Context()
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		// 将客户端标识信息存储在 context 中
		log.Println("-------------", md.Get("client-id"), md)
		ctx = context.WithValue(ctx, "client-id", md.Get("client-id"))
	}
	return handler(srv, &wrappedServerStream{ss, ctx})
}

type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}

type server struct {
	pb.UnimplementedMessagingServiceServer
	clients      map[string]chan pb.Message
	clientsExtra map[string]string
	mu           sync.Mutex
}

func (s *server) Connect(stream pb.MessagingService_ConnectServer) error {
	ch := make(chan pb.Message)
	var clientID string

	// Listen for incoming messages from the client
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Client: %s Failed to receive message: %v", clientID, err)
				break
			}
			if clientID == "" {
				clientID = in.ClientId
				s.mu.Lock()
				s.clients[clientID] = ch
				s.mu.Unlock()
			}
			log.Printf("Received message from client %s: %s", in.ClientId, in.Content)
			if clientID != "" {
				if in.FnName == CLIENT_STATUS {
					s.clientsExtra[clientID] = in.Content
				}
			}
			// Register the client channel
			s.mu.Lock()
			s.clients[clientID] = ch
			s.mu.Unlock()
		}

		// Client disconnected, clean up resources
		s.mu.Lock()
		if clientID != "" {
			delete(s.clients, clientID)
			delete(s.clientsExtra, clientID)
		}
		s.mu.Unlock()
		close(ch)
	}()

	// Send messages to the client
	for msg := range ch {
		if err := stream.Send(&msg); err != nil {
			log.Printf("Failed to send message: %v", err)
			break
		}
	}

	return nil
}

func (s *server) BroadcastMessage(content string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for clientID, ch := range s.clients {
		ch <- pb.Message{ClientId: clientID, Content: content, Type: pb.MessageType_BROADCAST}
	}
}

func (s *server) SyncLevelToClient(clientID, levelJson string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if ch, exists := s.clients[clientID]; exists {
		ch <- pb.Message{ClientId: clientID, Content: levelJson, Type: pb.MessageType_PRIVATE, FnName: SYNC_LEVEL}
	} else {
		log.Printf("Client with ID %s not found", clientID)
	}
}

func (s *server) SyncStartToClient(clientID, start string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if ch, exists := s.clients[clientID]; exists {
		ch <- pb.Message{ClientId: clientID, Content: start, Type: pb.MessageType_PRIVATE, FnName: SYNC_START}
	} else {
		log.Printf("Client with ID %s not found", clientID)
	}
}

func (s *server) SyncAllLevelConfig(content string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for clientID, ch := range s.clients {
		ch <- pb.Message{ClientId: clientID, Content: content, Type: pb.MessageType_BROADCAST, FnName: SYNC_LEVEL}
	}
}

func (s *server) SyncStartAllLevel(content string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for clientID, ch := range s.clients {
		ch <- pb.Message{ClientId: clientID, Content: content, Type: pb.MessageType_BROADCAST, FnName: SYNC_START}
	}
}

func (s *server) SyncClientLevelStatus() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for clientID, ch := range s.clients {
		ch <- pb.Message{ClientId: clientID, Content: "", Type: pb.MessageType_BROADCAST, FnName: SERVER_SYNC_STATUS}
	}
}

func (s *server) SendMessageToClient(clientID, content string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if ch, exists := s.clients[clientID]; exists {
		ch <- pb.Message{ClientId: clientID, Content: content, Type: pb.MessageType_PRIVATE}
	} else {
		log.Printf("Client with ID %s not found", clientID)
	}
}

func (s *server) GetClient() []string {
	keys := make([]string, 0, len(s.clients))
	for key := range s.clients {
		keys = append(keys, key)
	}
	return keys
}

func (s *server) GetClientExtra() map[string]string {
	return s.clientsExtra
}

func NewServer(rgpcPort string) *server {
	log.Println("===========GRPC================")
	log.Println("正在启动 GRPC server port: ", rgpcPort)
	lis, err := net.Listen("tcp", ":"+rgpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := &server{clients: make(map[string]chan pb.Message), clientsExtra: map[string]string{}}
	pb.RegisterMessagingServiceServer(s, srv)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	return srv
}
