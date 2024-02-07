# syntax=docker/dockerfile:1.2
FROM golang:1.22-bullseye as builder
LABEL maintainer="allan.nava@hiway.media"
#
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o main receiver/main.go
#
FROM phusion/baseimage:focal-1.2.0
WORKDIR /app
#
COPY --from=builder /app /app
#
CMD [ "./main" ]
#
