CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server src/server.go
chmod 777 server
