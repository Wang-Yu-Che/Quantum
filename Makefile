# Makefile to process .api files in subdirectories of 'restful' using goctl

# Define the directory to scan
API_DIR := restful

# Define a variable to store the list of filenames without extension
# Scan for .api files in subdirectories of 'restful' and extract filenames without extension
FILENAMES := $(shell find $(API_DIR) -mindepth 2 -type f -name "*.api" -exec sh -c 'basename "{}" .api' \;)
# 定义变量存储./interface文件夹及其子文件夹下所有.proto后缀的文件路径
PROTO_FILES := $(shell find ./interface -type f -name "*.proto")

# 排除文件名中包含 enum.proto 的文件
PROTO_FILES := $(filter-out ./interface/%enum.proto,$(PROTO_FILES))
ENUM_PROTO_FILES := $(shell find ./interface -type d -name enum -exec find {} -name "*enum.proto" \;)
# Default target
.PHONY: all
all: help

# api服务 make demo
.PHONY: $(FILENAMES)
$(FILENAMES):
	@echo "Processing $@..."
	goctl api format --dir "$(API_DIR)/$@/$@.api"
	goctl api go --api "$(API_DIR)/$@/$@.api" --dir "$(API_DIR)/$@" --style goZero
.PHONY: proto
proto:
	@echo "Generating code from .proto files..."
	@for file in $(PROTO_FILES); do \
		filename=$$(basename "$$file" .proto); \
		goctl rpc protoc "$$file" --go_out=./interface/$$filename/pb --go-grpc_out=./interface/$$filename/pb --zrpc_out=./service/$$filename --client=false; \
	done
	@echo "Code generation completed."
# 枚举定义 make enum
.PHONY: enum
enum:
	@echo "Generating Go files from proto files..."
	@for file in $(ENUM_PROTO_FILES); do \
		dir=$$(dirname $$file); \
		protoc --go_out=$$dir --proto_path=$$dir $$file; \
	done
	@echo "Generation complete."
.PHONY: test
test:
	@echo services: $(FILENAMES)
	@echo .proto: $(PROTO_FILES)
	@echo enum: $(ENUM_PROTO_FILES)
# Help target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make <filename>   Process a specific .api file by its filename (without extension)"
	@echo "  make help         Show this help message"