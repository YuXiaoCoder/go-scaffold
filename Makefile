.PHONY: build tools clean

BINARY="go-scaffold"

all: tools build-darwin

tools:
	@go fmt ./...
	@go vet ./...

build-darwin:
	@mkdir -p logs
	@mkdir -p output
	@swag init && CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -o output/${BINARY} main.go

build-linux:
	@swag init && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o output/${BINARY} main.go

clean:
	@rm -rf logs
	@rm -rf output

help:
	@echo "make - 格式化代码和静态检查代码，然后编译生成二进制文件"
	@echo "make tools - 格式化代码和静态检查代码"
	@echo "make build-darwin - 编译生成二进制文件（MacOS）"
	@echo "make build-linux - 编译生成二进制文件（Linux）"
	@echo "make clean - 移除二进制文件和日志文件"
