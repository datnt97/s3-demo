
include .env
OS := $(shell uname)
GOPATH:=$(shell go env GOPATH)
TARGET := $(firstword $(MAKECMDGOALS))
OPTIONAL_PARAM := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
API_PROTO_FILES=$(shell find api -name *.proto)
EXT_PROTO_FILES=$(shell find external -name *.proto)

# Export environment variables based on OS
ifeq ($(OS),Darwin)
    export $(shell sed 's/=.*//' .env)
else
	ifeq ($(OS),Linux)
		export $(shell cat .env | xargs)
	else
		export
	endif
endif

.PHONY: init
# init env
init:
	go install github.com/zeromicro/go-zero/tools/goctl@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	#goctl env check --install --verbose --force

.PHONY: commit
# push new change with upgrade commit message
commit:
	git add . && git commit -m "$(OPTIONAL_PARAM)" && git push

.PHONY: add
# add new proto file
add:
	goctl rpc -o api/$(path)/$(path).proto

.PHONY: grpc
# generate grpc code
grpc:
	protoc 	--proto_path=. \
			--proto_path=./protos \
			--go_out=paths=source_relative:. \
			--go-grpc_out=paths=source_relative:. \
	       	--openapi_out=fq_schema_naming=true,default_response=false:. \
			$(API_PROTO_FILES)

.PHONY: gateway
# generate gateway code
gateway:
	protoc 	--proto_path=. \
			--proto_path=./api \
			--proto_path=./protos \
			--go_out=paths=source_relative:. \
			--grpc-gateway_opt logtostderr=true \
			--grpc-gateway_opt paths=source_relative \
			--grpc-gateway_opt generate_unbound_methods=true \
			--grpc-gateway_out=paths=source_relative:. \
			$(API_PROTO_FILES)

.PHONY: oswagger
# generate grpc code
oswagger:
	protoc 	--proto_path=. \
			--proto_path=./protobuf \
			--openapiv2_out . \
			--openapiv2_opt logtostderr=true \
			--openapiv2_opt json_names_for_fields=false \
			$(API_PROTO_FILES)

.PHONY: ext
# generate grpc code
ext:
	protoc --proto_path=./ \
	       --proto_path=./protos \
 	       --go_out=paths=source_relative:./ \
 	       --go-grpc_out=paths=source_relative:./ \
	       $(EXT_PROTO_FILES)

.PHONY: run
# run app service
run:
ifeq ($(OS),Darwin)
	@echo "Setting environment variables from .env on macOS"
	@set -o allexport && source .env && set +o allexport && go run cmd/main/main.go
else
ifeq ($(OS),Linux)
	@echo "Setting environment variables from .env on Linux"
	@export $(shell cat .env | xargs) && go run cmd/main/main.go
else
	@echo "Unsupported operating system: $(OS)"
endif
endif

.PHONY: cron
# run cron service
cron:
	go run cmd/cron/main.go $(OPTIONAL_PARAM)

.PHONY: consumer
# run consumer service
consumer:
	go run cmd/consumer/main.go

.PHONY: swagger
# run swagger service
swagger:
	swagger generate spec -m -o ./swagger.yaml
	@echo "---"
	@echo Starting to validate the swagger file
	swagger validate ./swagger.yaml --skip-warnings