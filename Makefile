GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
APPNAME=algo

.PHONY: all
all: test build

.PHONY: build
build-cli:
	$(GOBUILD) -o $(APPNAME) cmd/cli/main.go

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(APPNAME)