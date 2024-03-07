# syntax=docker/dockerfile:1.7
FROM golang:1.20-bullseye as builder
LABEL maintainer="allan.nava@hiway.media"
#
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o main sender/main.go
#
FROM phusion/baseimage:focal-1.2.0
WORKDIR /app
#
COPY --from=builder /app /app
#
CMD [ "./main" ]
#
