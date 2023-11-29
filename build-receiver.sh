#!/bin/sh
#
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o receiver-binary  receiver/main.go
mkdir -p /tmp/bin
pwd
ls -al
tar -czvf /tmp/bin/receiver-binary .tar.gz --transform='s|.*/||'  ./receiver-binary 
#tar --ignore-zeros -xf
ls -l /tmp/bin
