
all:
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o serverbin ./server/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o clientbin ./client/main.go
