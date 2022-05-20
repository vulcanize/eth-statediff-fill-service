BIN = $(GOPATH)/bin

# Tools
## Testing library
GINKGO = $(BIN)/ginkgo
$(BIN)/ginkgo:
	go get github.com/onsi/ginkgo/ginkgo

## Migration tool
GOOSE = $(BIN)/goose
$(BIN)/goose:
	go get github.com/pressly/goose/cmd/goose

.PHONY: installtools
installtools: | $(LINT) $(GOOSE) $(GINKGO)
	echo "Installing tools"

.PHONY: test
test: | $(GINKGO) $(GOOSE)
	go vet ./...
	go fmt ./...
	$(GINKGO) -r --skipPackage=test

.PHONY: integrationtest
integrationtest: | $(GINKGO) $(GOOSE)
	go vet ./...
	go fmt ./...
	$(GINKGO) -r test/ -v

build:
	go fmt ./...
	GO111MODULE=on go build
