.PHONY: clean build run

run:
	bin/49775.exe

build:
	go env -w CGO_CFLAGS="-g -O2 -I E:/Go/cgo49775/include"
	go env -w CGO_LDFLAGS="-g -O2 -L E:/Go/cgo49775/bin"
	go build -o bin/49775.exe .
	go env -w CGO_CFLAGS="-g -O2"
	go env -w CGO_LDFLAGS="-g -O2"

# 清理
clean:
	- go clean 
	- go clean -cache