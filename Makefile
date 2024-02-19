include .env

PORT ?= ${PORT}

.PHONY: build frontend_build build_mac build_win build_linux

current_dir := $(shell pwd)

build: frontend_build build_mac build_win build_linux

frontend_build: oas_ts_fetch
	cd frontend && bun install && bun run build

build_mac: frontend_build
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w -X main.AppMode=prod" -trimpath -o $(current_dir)/bin/mac/go-react-local $(current_dir)/main.go

# use mingw-w64
build_win: frontend_build
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -ldflags="-s -w -X main.AppMode=prod" -trimpath -o $(current_dir)/bin/win/go-react-local.exe $(current_dir)/main.go

build_linux: frontend_build
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X main.AppMode=prod" -trimpath -o ./bin/linux/go-react-local ./main.go

dev: port_check
	cd frontend && bun install && bun run dev & ENV=dev air && fg

oas_ts_fetch:
	docker run --rm -v $(current_dir):/local \
    openapitools/openapi-generator-cli \
    generate \
    -g typescript-fetch \
    -i /local/openapi.yml \
    -o /local/frontend/src/openapi \
    --api-package api \
    --model-package model \
    --generate-alias-as-model \
    --additional-properties withInterfaces=true \
    --additional-properties withSeparateModelsAndApi=true \
		--additional-properties enumPropertyNaming=PascalCase

# やってはみたものの導入方法がよくわからない
__oas_go_kiota:
	docker run --rm -v ${current_dir}:/local \
	mcr.microsoft.com/openapi/kiota \
	generate --language go \
	--openapi /local/openapi.yml \
	-o /local/kiota_test \
	-n sample/go-react-local-app/openapi

# oganを使いたいがginにうまく組み込むやり方がわからない
oas_ogen:
	ogen -package openapi -target ${current_dir}/internal/ogen_openapi -clean ${current_dir}/openapi.yml

copy_data:
	cp -r ./data ./bin/$(PLATFORM)/


port_check:
	@echo "Checking the availability of port $(PORT)..."
	@lsof -i :$(PORT) > /dev/null 2>&1; if [ $$? -eq 0 ]; then \
		echo "👺 Error: Port $(PORT) is already in use."; \
		echo "Details of the process occupying the port:"; \
		lsof -i :$(PORT) | awk 'NR>1 {print "PID: "$$2", User: "$$3", Command: "$$1}'; \
		exit 1; \
	else \
		echo "🎉 Port $(PORT) is not in use. let's go!"; \
	fi