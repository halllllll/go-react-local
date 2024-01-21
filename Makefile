build:
	cd frontend && bun run build
	GIN_MODE=release & GOOS=darwin GOARCH=amd64 go build  -ldflags="-s -w" -trimpath -o ./bin/go-react-local ./main.go

dev:
	cd frontend && bun run dev & ENV=dev air && fg