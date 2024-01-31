.PHONY: build frontend_build build_mac build_win build_linux

build: frontend_build build_mac build_win build_linux

frontend_build:
	cd frontend && bun install && bun run build

build_mac: frontend_build
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w -X main.AppMode=prod" -trimpath -o ./bin/mac/go-react-local ./main.go
	$(MAKE) copy_data PLATFORM=mac

# use mingw-w64
build_win: frontend_build
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -ldflags="-s -w -X main.AppMode=prod" -trimpath -o ./bin/win/go-react-local.exe ./main.go
	$(MAKE) copy_data PLATFORM=win

build_linux: frontend_build
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X main.AppMode=prod" -trimpath -o ./bin/linux/go-react-local ./main.go

dev:
	cd frontend && bun install && bun run dev & ENV=dev air && fg

copy_data:
	cp -r ./data ./bin/$(PLATFORM)/