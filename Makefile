.PHONY: build frontend_build build_mac build_win build_linux

build: frontend_build build_mac build_win build_linux

frontend_build:
	cd frontend && bun run build

build_mac: frontend_build
	GIN_MODE=release & GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o ./bin/mac/go-react-local ./main.go

build_win: frontend_build
	GIN_MODE=release & GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o ./bin/win/go-react-local.exe ./main.go

build_linux: frontend_build
	GIN_MODE=release & GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o ./bin/linux/go-react-local ./main.go

dev:
	cd frontend && bun run dev & ENV=dev air && fg