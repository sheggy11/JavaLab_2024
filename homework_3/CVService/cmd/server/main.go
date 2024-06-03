package main

import (
	"CVService/internal/config"
	"CVService/internal/repository/mongo"
	"CVService/internal/service"
	"CVService/proto"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	host = "localhost"
	port = "5000"
)

func main() {
	ctx := context.Background()

	err := setupViper()

	if err != nil {
		log.Fatalf("error reading uml file: %v", err)
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("error starting tcp listener: %v", err)
	}

	mongoDataBase, err := config.SetupMongoDataBase(ctx)

	if err != nil {
		log.Fatalf("error starting mongo: %v", err)
	}

	cvRepository := mongo.NewCVRepository(mongoDataBase.Collection("cvs"))
	cvService := service.NewCVService(cvRepository)

	grpcServer := grpc.NewServer()

	proto.RegisterCVServiceServer(grpcServer, cvService)

	log.Printf("gRPC stated at %v\n", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error statring gRPC: %v", err)
	}

}

func setupViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
