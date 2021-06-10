all: build
.PHONY: all

build:
	go build -ldflags "-X k8s.io/client-go/pkg/version.gitVersion=$$(git describe --abbrev=8 --dirty --always)" -o bin/kcp ./cmd/kcp
	go build -ldflags "-X k8s.io/client-go/pkg/version.gitVersion=$$(git describe --abbrev=8 --dirty --always)" -o bin/syncer ./cmd/syncer
	go build -ldflags "-X k8s.io/client-go/pkg/version.gitVersion=$$(git describe --abbrev=8 --dirty --always)" -o bin/cluster-controller ./cmd/cluster-controller
	go build -ldflags "-X k8s.io/client-go/pkg/version.gitVersion=$$(git describe --abbrev=8 --dirty --always)" -o bin/deployment-splitter ./cmd/deployment-splitter
	go build -ldflags "-X k8s.io/client-go/pkg/version.gitVersion=$$(git describe --abbrev=8 --dirty --always)" -o bin/group-deployment ./cmd/group-deployment
.PHONY: build

vendor:
	go mod tidy
	go mod vendor
.PHONY: vendor

codegen:
	./hack/update-codegen.sh
.PHONY: codegen
