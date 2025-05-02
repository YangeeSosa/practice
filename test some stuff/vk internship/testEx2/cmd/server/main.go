package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"testEx2/api"
	"testEx2/config"
	"testEx2/pkg/subpub"
)

type server struct {
	api.UnimplementedPubSubServer
	pubsub subpub.SubPub
}

func (s *server) Subscribe(req *api.SubscribeRequest, stream api.PubSub_SubscribeServer) error {
	if req.Key == "" {
		return status.Error(codes.InvalidArgument, "ключ не может быть пустым")
	}

	// Создание канала для получения сообщений
	msgChan := make(chan interface{}, 100)

	// Подписка на систему pubsub
	sub, err := s.pubsub.Subscribe(req.Key, func(msg interface{}) {
		msgChan <- msg
	})
	if err != nil {
		return status.Error(codes.Internal, "не удалось подписаться")
	}
	defer sub.Unsubscribe()

	// Обработка входящих сообщений
	for {
		select {
		case msg := <-msgChan:
			if str, ok := msg.(string); ok {
				if err := stream.Send(&api.Event{Data: str}); err != nil {
					return status.Error(codes.Internal, "не удалось отправить событие")
				}
			}
		case <-stream.Context().Done():
			return nil
		}
	}
}

func (s *server) Publish(ctx context.Context, req *api.PublishRequest) (*emptypb.Empty, error) {
	if req.Key == "" {
		return nil, status.Error(codes.InvalidArgument, "ключ не может быть пустым")
	}

	if err := s.pubsub.Publish(req.Key, req.Data); err != nil {
		return nil, status.Error(codes.Internal, "не удалось опубликовать сообщение")
	}

	return &emptypb.Empty{}, nil
}

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Printf("Не удалось загрузить конфигурацию, используем значения по умолчанию: %v", err)
		cfg = &config.Config{GRPCPort: ":50051"}
	}

	// Создаем новый pubsub
	ps, err := subpub.NewSubPub()
	if err != nil {
		log.Fatalf("Не удалось создать pubsub: %v", err)
	}

	// Создаем gRPC сервер
	grpcServer := grpc.NewServer()
	api.RegisterPubSubServer(grpcServer, &server{pubsub: ps})

	// Начинаем слушать
	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Не удалось начать прослушивание: %v", err)
	}

	// Обрабатываем graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Println("Завершаем работу gRPC сервера...")
		grpcServer.GracefulStop()

		if err := ps.Close(context.Background()); err != nil {
			log.Printf("Ошибка при закрытии pubsub: %v", err)
		}
	}()

	log.Printf("Запускаем gRPC сервер на %s", cfg.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
