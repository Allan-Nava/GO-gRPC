#!/bin/sh
#
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o sender-binary  sender/main.go
mkdir -p /tmp/bin
pwd
ls -al
tar -czvf /tmp/bin/sender-binary .tar.gz --transform='s|.*/||'  ./sender-binary 
#tar --ignore-zeros -xf
ls -l /tmp/bin
