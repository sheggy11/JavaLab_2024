package ru.itis.hr.config;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import ru.itis.hr.grpc.pb.cv.CVServiceGrpc;
import ru.itis.hr.grpc.pb.job.JobServiceGrpc;

@Configuration
public class GrpcServicesConfig {

    @Bean("cvService")
    public CVServiceGrpc.CVServiceBlockingStub cvService(@Value(value = "${grpc.cv-service.address}") String address,
                                                         @Value(value = "${grpc.cv-service.port}") int port) {

        ManagedChannel channel = ManagedChannelBuilder.forAddress(address, port)
                .usePlaintext()
                .build();
        return CVServiceGrpc.newBlockingStub(channel);
    }

    @Bean("jobService")
    public JobServiceGrpc.JobServiceBlockingStub jobService(@Value(value = "${grpc.job-client.address}") String address,
                                                            @Value(value = "${grpc.job-client.port}") int port) {

        ManagedChannel channel = ManagedChannelBuilder.forAddress(address, port)
                .usePlaintext()
                .build();
        return JobServiceGrpc.newBlockingStub(channel);
    }
}
