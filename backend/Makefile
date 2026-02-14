run:
	go run ./main.go

dev:
	air run

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./local/backend_linux ./server
	goupx ./local/backend_linux
