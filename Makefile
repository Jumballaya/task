# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
ENTRY=main.go
BINARY_NAME=task
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build:
	$(GOBUILD) -v -o dist/$(BINARY_NAME) $(ENTRY)
test:
	$(GOTEST) ./...
cover:
	$(GOTEST) ./... -cover
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -v -o dist/$(BINARY_NAME) $(ENTRY)
	./dist/$(BINARY_NAME)
deps:
	echo "No dependencies..."
build-linux:
	make deps
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -v -o dist/$(BINARY_UNIX) $(ENTRY)
