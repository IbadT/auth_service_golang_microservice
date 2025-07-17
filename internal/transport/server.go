package transportgrpc

import (
	"log"
	"net"

	"github.com/IbadT/auth_service_golang_microservice/internal/auth"
	authpb "github.com/IbadT/project-protos/proto/auth"
	"google.golang.org/grpc"
)

func RunGRPC(svc auth.Service) error {
	listen, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Ошибка при запуске сервиса %v", err)
		return err
	}

	grpcSrv := grpc.NewServer()

	authpb.RegisterAuthServiceServer(grpcSrv, NewHandler(svc))

	log.Printf("GRPC запущен на порту 50053")
	if err := grpcSrv.Serve(listen); err != nil {
		log.Fatalf("Ошибка при запуске grpc сервера %v", err)
		return err
	}
	return nil
}
