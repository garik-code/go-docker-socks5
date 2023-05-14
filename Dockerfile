# Base image, golang 1.18
FROM golang:1.18.3
WORKDIR /workspace
# Copy all files into the image
COPY . .
# Run go mod
RUN go mod init main
RUN go get github.com/things-go/go-socks5
# Expose ports
EXPOSE 8000
# Run Go program, just like locally
ENTRYPOINT ["go","run","main.go"]