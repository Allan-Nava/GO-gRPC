# syntax=docker/dockerfile:1.2
FROM golang:1.20-bullseye as builder
LABEL maintainer="allan.nava@hiway.media"
#
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o ./go-grpc . 
#
#
FROM phusion/baseimage:focal-1.2.0
WORKDIR /app
#
COPY --from=builder /app /app
#
CMD [ "./go-grpc" ]
#
