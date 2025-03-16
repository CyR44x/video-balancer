package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"video-balancer/internal/service"
	pb "video-balancer/proto"
)

func main() {
	cdnHost := os.Getenv("CDN_HOST")
	if cdnHost == "" {
		log.Fatal("Переменная окружения CDN_HOST не установлена")
	}

	port := flag.String("port", "50051", "порт для прослушивания gRPC сервера")
	flag.Parse()

	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("Ошибка прослушивания порта: %v", err)
	}

	grpcServer := grpc.NewServer()

	svc := service.NewBalancerService(cdnHost)
	pb.RegisterServiceServer(grpcServer, svc)

	reflection.Register(grpcServer)

	fmt.Printf("Сервер запущен на порту %s\n", *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка запуска gRPC сервера: %v", err)
	}
}
