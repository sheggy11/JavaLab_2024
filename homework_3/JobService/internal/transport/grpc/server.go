package grpc

import (
    "fmt"
    "google.golang.org/grpc"
    "job-service/internal/config"
    "job-service/internal/service"
    "job-service/proto"
    "net"
)

func ServeGRPC(config *config.GRPCConfig, jobService *service.JobService) error {
    addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
    lis, err := net.Listen("tcp", addr)
    if err != nil {
        return fmt.Errorf("start tcp listener: %w", err)
    }

    grpcServer := grpc.NewServer()

    proto.RegisterJobServiceServer(grpcServer, jobService)

    go func() {
        if err := grpcServer.Serve(lis); err != nil {
            panic(err)
        }
    }()

    return nil
}
