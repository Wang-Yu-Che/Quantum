# Makefile to process .api files in subdirectories of 'restful' using goctl

# Define the directory to scan
API_DIR := restful

# Apifox
APIFOX_PROJECT_ID := 5185764
APIFOX_TOKEN := afxp_d61b1beVqkVIcdsyg4nM8rIYwRzR16tBCk2a

# Define a variable to store the list of filenames without extension
# Scan for .api files in subdirectories of 'restful' and extract filenames without extension
FILENAMES := $(shell find $(API_DIR) -mindepth 2 -type f -name "*.api" -exec sh -c 'basename "{}" .api' \;)

# 定义变量存储./interface文件夹及其子文件夹下所有.proto后缀的文件路径
PROTO_FILES := $(shell find ./interface -type f -name "*.proto")

# 排除文件名中包含 enum.proto 的文件
PROTO_FILES := $(filter-out ./interface/%enum.proto,$(PROTO_FILES))
ENUM_PROTO_FILES := $(shell find ./interface -type d -name enum -exec find {} -name "*enum.proto" \;)

# 数据库名
TARGETS := $(shell go run ./script/gorm/list_db_name.go)

# Default target
.PHONY: all
all: help

.PHONY: install
install:
	@echo "Install Go-Zero"
	go get -u github.com/zeromicro/go-zero@latest
	go install github.com/zeromicro/go-zero/tools/goctl@latest
	goctl env check --install --verbose --force
	@echo "Install GORM"
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/sqlite
	go get -u gorm.io/gen
	go install gorm.io/gen/tools/gentool@latest

# api服务: make 服务名
.PHONY: $(FILENAMES)
$(FILENAMES):
	@echo "Processing $@..."
	goctl api format --dir "$(API_DIR)/$@/$@.api"
	goctl api go --api "$(API_DIR)/$@/$@.api" --dir "$(API_DIR)/$@" --style go_zero

# rpc服务: 一键生成(包括枚举)
.PHONY: proto
proto:
	@echo "Generating Enum Go files from proto files..."
	@for file in $(ENUM_PROTO_FILES); do \
		dir=$$(dirname $$file); \
		echo "Processing enum file: $$file"; \
		protoc --go_out=$$dir --proto_path=$$dir $$file; \
	done
	@echo "------------------------------------------------------------------"
	@echo "Generating Service code from .proto files..."
	@for file in $(PROTO_FILES); do \
		filename=$$(basename "$$file" .proto); \
		echo "Processing proto file: $$file"; \
		goctl rpc protoc "$$file" --go_out=./interface/$$filename/pb --go-grpc_out=./interface/$$filename/pb --zrpc_out=./service/$$filename --style go_zero; \
	done
	@echo "Code generation completed."
	@$(MAKE) --no-print-directory reset

# GORM数据库表model生成
# 定义模式匹配规则
%_model:
	@echo "Building $@"
	@go run ./script/gorm/gen_db_model.go -db=$(patsubst %_model,%,$@)

# 分隔符差异消除
.PHONY: reset
reset:
#   @git config --global core.autocrlf false
	@rm -f .git/index
	@git reset -q

.PHONY: test
test:
	@echo services: $(FILENAMES)
	@echo .proto: $(PROTO_FILES)
	@echo enum: $(ENUM_PROTO_FILES)
	@echo models: $(TARGETS)

# 生成所有 API 的 swagger 文档
.PHONY: swagger
swagger:
	@echo "Generating swagger files for all api services..."
	@mkdir -p ./docs/swagger
	@for name in $(FILENAMES); do \
		api_file="$(API_DIR)/$$name/$$name.api"; \
		echo "Generating swagger for $$api_file"; \
		goctl api swagger --api "$$api_file" --dir ./docs/swagger --filename "$$name"; \
	done
	@echo "Swagger generation completed. Files are in ./docs/swagger"

.PHONY: apifox
apifox:
	@echo "Pushing swagger files in ./docs/swagger to Apifox..."
	@for file in ./docs/swagger/*.yaml ./docs/swagger/*.json; do \
		if [ -f "$$file" ]; then \
			echo "Uploading $$file to Apifox..."; \
			curl -X POST "https://api.apifox.com/v1/projects/$(APIFOX_PROJECT_ID)/import-openapi" \
              -H "Authorization: Bearer $(APIFOX_TOKEN)" \
              -F "file=@$$file"
		fi \
	done
	@echo "Apifox sync completed."



# git rm -r --cached .  #清除缓存
# git add . #重新trace file

# Help target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make install        Install required tools (Go-Zero, GORM)"
	@echo "  make <filename>     Process a specific .api file by its filename (without extension)"
	@echo "  make proto          Generate RPC code from .proto files (excluding enum files)"
	@echo "  make <model_name>_model  Generate GORM models for a specific database"
	@echo "  make swagger        Generate swagger files for all api services into ./docs"
	@echo "  make reset          Reset the repository (remove .git/index)"
	@echo "  make test           Print debug information (services, proto files, enum files, models)"
	@echo "  make help           Show this help message"
