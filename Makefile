build:
	cd frontend && bun run build
	GIN_MODE=release & GOOS=darwin GOARCH=arm64 go build  -ldflags="-s -w" -trimpath -o ./bin/mac/go-react-local ./main.go
	GIN_MODE=release & GOOS=windows GOARCH=amd64 go build  -ldflags="-s -w" -trimpath -o ./bin/win/go-react-local.exe ./main.go
	GIN_MODE=release & GOOS=linux GOARCH=amd64 go build  -ldflags="-s -w" -trimpath -o ./bin/linux/go-react-local ./main.go

dev:
	cd frontend && bun run dev & ENV=dev air && fg