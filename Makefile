VERSION=$(shell git describe --always --match "v[0-9]*" HEAD)
BUILD_INFO=-ldflags "-X github.com/kilingzhang/BingGPT/internal/version.version=$(VERSION)"
GO_BUILD_TAGS="jsoniter"
DOCKER_BASE=kilingzhang
DOCKER_IMAGE_NAME=BingGPT
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

.PHONY: tidy
gotidy:
	export GOPROXY="https://goproxy.io,direct" &&  go mod tidy -compat=1.20 && go mod vendor

.PHONY: run
run:BingGPT
	./bin/BingGPT_$(GOOS)_$(GOARCH)$(EXTENSION) server

.PHONY: BingGPT
BingGPT:
	GO111MODULE=on CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -trimpath -o ./bin/BingGPT_$(GOOS)_$(GOARCH)$(EXTENSION) \
		$(BUILD_INFO) -tags $(GO_BUILD_TAGS) ./cmd/server

.PHONY: BingGPT-all-sys
BingGPT-all-sys: BingGPT-darwin_amd64 BingGPT-darwin_arm64 BingGPT-linux_amd64 BingGPT-linux_arm64 BingGPT-windows_amd64

.PHONY: BingGPT-darwin_amd64
BingGPT-darwin_amd64:
	GOOS=darwin GOARCH=amd64 $(MAKE) BingGPT

.PHONY: BingGPT-darwin_arm64
BingGPT-darwin_arm64:
	GOOS=darwin GOARCH=arm64 $(MAKE) BingGPT

.PHONY: BingGPT-linux_amd64
BingGPT-linux_amd64:
	GOOS=linux GOARCH=amd64 $(MAKE) BingGPT

.PHONY: BingGPT-linux_arm64
BingGPT-linux_arm64:
	GOOS=linux GOARCH=arm64 $(MAKE) BingGPT

.PHONY: BingGPT-windows_amd64
BingGPT-windows_amd64:
	GOOS=windows GOARCH=amd64 EXTENSION=.exe $(MAKE) BingGPT

.PHONY: build-docker
build-docker:
	docker buildx build --push --platform linux/amd64,linux/arm64 -t $(DOCKER_BASE)/$(DOCKER_IMAGE_NAME):$(VERSION) .


#--net=bitwarden_gcloud_default
.PHONY: run-docker
run-docker:
	docker stop BingGPT_container; \
	docker rm BingGPT_container; \
	docker run -itd \
	--name=BingGPT_container \
	--cap-add=SYS_PTRACE  \
	-p 12527:12527 \
    $(DOCKER_BASE)/BingGPT:$(VERSION)

.PHONY: clean
clean:
	@rm -rf ./bin
	@rm -rf ./tmp
	@rm -rf ./logs
