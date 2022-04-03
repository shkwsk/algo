GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
APPNAME=algo

.PHONY: all
all: test build

.PHONY: build
build:
	$(GOBUILD) -o $(APPNAME) cmd/$(APPNAME)/main.go

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: run
run:
	$(GOBUILD) -o $(APPNAME) cmd/$(APPNAME)/main.go
	./$(APPNAME)

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(APPNAME)