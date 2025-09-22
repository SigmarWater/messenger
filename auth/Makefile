#include vendor.proto.mk

# Генерация Go и gRPC файлов через buf
.buf-generate: .bin-deps
	$(info Running buf generate...)
	buf generate

# Основная цель
generate: .buf-generate .tidy

.tidy:
	go mod tidy

# Установка бинарных зависимостей
.bin-deps:
	$(info Installing binary dependencies...)
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: .bin-deps .buf-generate generate
