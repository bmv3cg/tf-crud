GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=tfe-ws-manager
BINARY_UNIX=$(BINARY_NAME)_unix

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v