package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"job-service/internal/config"
	"job-service/internal/repository/mongo"
	"job-service/internal/service"
	"job-service/proto"
	"log"
	"net"
)

var (
	host = "localhost"
	port = "5010"
)

func main() {
	ctx := context.Background()

	err := setupViper()

	if err != nil {
		log.Fatalf("Error while reading uml file: %v", err)
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("TCP", addr)

	if err != nil {
		log.Fatalf("Error while starting TCP listener: %v", err)
	}

	mongoDataBase, err := config.SetupMongoDataBase(ctx)

	if err != nil {
		log.Fatalf("Error while starting mongo: %v", err)
	}

	jobRepository := mongo.NewJobRepository(mongoDataBase.Collection("job"))
	jobService := service.NewJobService(jobRepository)

	grpcServer := grpc.NewServer()

	proto.RegisterJobServiceServer(grpcServer, jobService)

	log.Printf("gRPC stated at %v\n", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error while statring gRPC: %v", err)
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
