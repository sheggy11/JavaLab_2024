FROM golang:1.22-alpine as builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o ./bin/cv-service ./cmd/server/main.go





FROM alpine
COPY --from=builder /build/bin/cv-service /cv-service
COPY --from=builder /build/configs/config.yml /configs/config.yaml

EXPOSE 5011

ENTRYPOINT ["/cv-service"]