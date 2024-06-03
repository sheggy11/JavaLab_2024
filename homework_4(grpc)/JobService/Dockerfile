FROM golang:1.22-alpine as builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o ./bin/job-service ./cmd/server/main.go








FROM alpine
COPY --from=builder /build/bin/job-service /job-service
COPY --from=builder /build/configs/config.yml /configs/config.yaml

EXPOSE 5010

ENTRYPOINT ["/job-service"]