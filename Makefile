SWEEP?=global
TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
WEBSITE_REPO=github.com/hashicorp/terraform-website
PKG_NAME=iamtf
GOLINT=./bin/golangci-lint
GOFMT:=gofumpt
#TFPROVIDERLINT=/home/sgonzalez/wa/git/go/bin/tfproviderlint
TF_ACC_TEMP_DIR=.tmp

#GO_SRC?=..
OS_ARCH?=$(shell go env GOOS)_$(shell go env GOARCH)
HOSTNAME=atricore.com
NAMESPACE=iam
#VERSION=0.1.8
VERSION=$(shell git describe --tags --always --dirty)

NAME=iamtf
OUT_DIR=./.tmp/$(GOOS)/$(GOARCH)/$(VERSION)
BINARY=terraform-provider-${NAME}

PLATFORMS=darwin linux windows openbsd
ARCHITECTURES=386 amd64,

# Setup linker flags option for build that interoperate with variable names in src code
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Expression to match against tests
# go test -run <filter>
# e.g. Iden will run all TestAccIdentity tests
ifdef TEST_FILTER
	TEST_FILTER := -run $(TEST_FILTER)
endif

default: build

dep-replace: # Download required dependencies (DEV MODE: use local dependencies)
	go mod edit -replace github.com/atricore/josso-api-go='$(GO_SRC)'../josso-api-go
	go mod edit -replace github.com/atricore/josso-sdk-go='$(GO_SRC)'../josso-sdk-go

dep-dropreplace:
	go mod edit -dropreplace github.com/atricore/josso-api-go
	go mod edit -dropreplace github.com/atricore/josso-sdk-go


dep: # Download required dependencies
	go mod tidy
	go mod vendor

# For develoeprs!
dep-dev: dep-replace dep dep-dropreplace
build-dev: dep-replace fmtcheck install dep-dropreplace
test-dev: fmtcheck install test dep-dropreplace
testacc-dev: fmtcheck install testacc dep-dropreplace

build: fmtcheck install

install:
	go build ${LDFLAGS} -o ${BINARY}
#	go install ./...


build_all:
	$(foreach GOOS, $(PLATFORMS),\
		$(foreach GOARCH, $(ARCHITECTURES), \
			$(shell export GOOS=$(GOOS); \
				export GOARCH=$(GOARCH); \
				go build -v -o $(OUT_DIR)/$(BINARY); \
				if test -f "$(OUT_DIR)/$(BINARY)" ; then cd $(OUT_DIR) ; zip -q ../../../$(BINARY)-$(GOOS)-$(GOARCH)-$(VERSION).zip $(BINARY) ; fi)))

sweep:
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	go test $(TEST) -v -sweep=$(SWEEP) $(SWEEPARGS)

test: fmtcheck
	go test $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) $(TEST_FILTER) -timeout=30s -parallel=4

testacc: fmtcheck
	mkdir -p $(TF_ACC_TEMP_DIR)
	cp -r ./acctest-data/ $(TF_ACC_TEMP_DIR)
	go clean -testcache
	TF_ACC_TEMP_DIR=`pwd`/$(TF_ACC_TEMP_DIR) TF_ACC=1 go test $(TEST) -v $(TESTARGS) $(TEST_FILTER) -timeout 120m

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

.PHONY: fmt install build build_all
fmt: check-fmt # Format the code
	@echo "formatting the code..."
	@$(GOFMT) -l -w $$(find . -name '*.go' |grep -v vendor)

check-fmt:
	@which $(GOFMT) > /dev/null || (echo "downloading formatter..." && GO111MODULE=on go get mvdan.cc/gofumpt)

fmtcheck: 
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

install-provider: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/$(OS_ARCH)
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/$(OS_ARCH)

lint: tools
	@echo "==> Checking source code against linters..."
	@GOGC=30 $(GOLINT) run ./$(PKG_NAME)
	@$(TFPROVIDERLINT) \
		-c 1 \
		-AT001 \
    -R004 \
		-S001 \
		-S002 \
		-S003 \
		-S004 \
		-S005 \
		-S007 \
		-S008 \
		-S009 \
		-S010 \
		-S011 \
		-S012 \
		-S013 \
		-S014 \
		-S015 \
		-S016 \
		-S017 \
		-S019 \
		./$(PKG_NAME)

tools:
	@which $(GOLINT) || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.33.0
	@which $(TFPROVIDERLINT) || go install github.com/bflad/tfproviderlint/cmd/tfproviderlint

docs: build
	go run ./iamtf-docs/iamtf-docs.go "$(CURDIR)/reference/.tmp" "$(CURDIR)/reference"


