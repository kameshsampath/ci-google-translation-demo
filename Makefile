GO_PROTOC_IMAGE?=kameshsampath/protoc-go
WEB_PROTOC_IMAGE?=kameshsampath/protoc-grpc-web
GO_PROTOC_IMAGE_CACHE?=kameshsampath/protoc-go-cache
WEB_PROTOC_IMAGE_CACHE?=kameshsampath/protoc-grpc-web-cache
UI_CACHE?=type=registry,ref=kameshsampath/docker-extension-ui-cache
UI_BAS_IMAGE=kameshsampath/drone-ci-extension-ui-base
TAG?=latest

BUILDER=buildx-multi-arch

build-go-protoc:	prepare-buildx ## Build protoc with go plugins
	docker buildx build --builder=$(BUILDER) -f docker/Dockerfile.go.protoc . --push --pull=true --cache-to=$(GO_PROTOC_IMAGE_CACHE) --cache-from=$(GO_PROTOC_IMAGE_CACHE) --platform linux/amd64,linux/arm64 -t $(GO_PROTOC_IMAGE):$(TAG)

build-web-protoc:	prepare-buildx ## Build protoc with grpc-web plugins
	docker buildx build --builder=$(BUILDER) -f docker/Dockerfile.web.protoc . --push --pull=true --cache-to=$(WEB_PROTOC_IMAGE_CACHE) --cache-from=$(WEB_PROTOC_IMAGE_CACHE) --platform linux/amd64,linux/arm64 -t $(WEB_PROTOC_IMAGE):$(TAG)

prepare-buildx: ## Create buildx builder for multi-arch build, if not exists
	docker buildx inspect $(BUILDER) || docker buildx create --name=$(BUILDER) --driver=docker-container --driver-opt=network=host

build-server:	prepare-buildx ## Build protoc with go plugins
	docker buildx build --builder=$(BUILDER) --output="type=docker" -t kameshsampath/lingua-greeter-server:$(TAG) cmd/server 

build-client:	prepare-buildx ## Build protoc with go plugins
	docker buildx build --builder=$(BUILDER) --output="type=docker" -t kameshsampath/lingua-greeter-client:$(TAG) cmd/client

images:
	drone exec --trusted --exclude="protoc-server" --exclude="protoc-web"

all:	build-go-protoc

.PHONY:	build-go-protoc all build-server build-client