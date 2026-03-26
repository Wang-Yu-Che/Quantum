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
.PHONY: gen-swagger upload-swagger

gen-swagger:
	@echo "Generating swagger files..."
	@mkdir -p ./docs/swagger
	@for name in $(FILENAMES); do \
		api_file="$(API_DIR)/$$name/$$name.api"; \
		goctl api swagger --api "$$api_file" --dir ./docs/swagger --filename "$$name" || exit 1; \
	done
	@echo "Swagger generation completed. Files are in ./docs/swagger"


upload-swagger:
	@echo "Uploading all swagger files to Apifox..."
	@go run ./script/swagger/upload_apifox.go || exit 1
	@echo "Apifox sync completed."

swagger: gen-swagger upload-swagger

# Docker image generation from Go files in service and restful folders
.PHONY: docker
docker:
	@echo "Generating Docker files for services..."
	@for base_dir in service restful; do \
		find "$$base_dir" -mindepth 1 -maxdepth 1 -type d | while read dir; do \
			dir_name=$$(basename "$$dir"); \
			go_file="$$dir/$$dir_name.go"; \
			if [ -f "$$go_file" ]; then \
				echo "Processing Docker for $$dir_name in $$dir"; \
				(cd "$$dir" && goctl docker -go "$$dir_name.go"); \
			fi; \
		done; \
	done
	@echo "Docker generation completed."

# 基础镜像前缀和命名空间
IMAGE_PREFIX = ghcr.io/wang-yu-che/quantum
NAMESPACE = default
OUTPUT_DIR = build/k8s

.PHONY: kube
kube:
	@echo "Generating K8s Deployment files into $(OUTPUT_DIR)..."
	@mkdir -p $(OUTPUT_DIR)
	@for base_dir in restful service; do \
		if [ -d "$$base_dir" ]; then \
			find "$$base_dir" -mindepth 1 -maxdepth 1 -type d | while read dir; do \
				dir_name=$$(basename "$$dir"); \
				yaml_file="$$dir/etc/$$dir_name.yaml"; \
				if [ -f "$$yaml_file" ]; then \
					if [ "$$base_dir" = "restful" ]; then \
						port=$$(grep "^Port:" "$$yaml_file" | awk '{print $$2}'); \
						kube_name="$$dir_name-api"; \
					else \
						port=$$(grep "^ListenOn:" "$$yaml_file" | awk -F ':' '{print $$NF}'); \
						kube_name="$$dir_name-rpc"; \
					fi; \
					\
					if [ -z "$$port" ]; then \
						echo "Warning: Port/ListenOn not found in $$yaml_file, skipping..."; \
					else \
						echo "Processing $$kube_name (Port: $$port) from $$yaml_file"; \
						goctl kube deploy \
							-name "$$kube_name" \
							-namespace $(NAMESPACE) \
							-image "$(IMAGE_PREFIX)/$$kube_name:latest" \
							-port "$$port" \
							-targetPort "$$port" \
							-replicas 1 \
							-minReplicas 1 \
							-maxReplicas 2 \
							-requestCpu 50 \
							-requestMem 64 \
							-limitCpu 800 \
							-limitMem 160 \
							-revisions 3 \
							-imagePullPolicy IfNotPresent \
							-o "$(OUTPUT_DIR)/$$kube_name.yaml"; \
					fi; \
				fi; \
			done; \
		fi; \
	done
	@echo "K8s generation completed."

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
	@echo "  make docker        Generate Docker files for services with matching .go files"
	@echo "  make kube           Batch generate K8s YAMLs into build/k8s (scans Port/ListenOn from etc/*.yaml)"
	@echo "  make swagger        Generate swagger files for all api services into ./docs"
	@echo "  make reset          Reset the repository (remove .git/index)"
	@echo "  make test           Print debug information (services, proto files, enum files, models)"
	@echo "  make help           Show this help message"
