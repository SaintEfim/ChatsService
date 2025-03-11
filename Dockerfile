FROM golang:1.23.1 AS builder

WORKDIR /chatservice
COPY . .
RUN go mod download && go build -o chatservice cmd/main.go

FROM alpine:3.18.4

WORKDIR /app
RUN apk add gcompat
VOLUME /app/config
COPY --from=builder /chatservice/chatservice .

ENTRYPOINT [ "/app/chatservice" ]
