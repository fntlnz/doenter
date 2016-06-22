APP=doenter
REPO=fntlnz/doenter

export GO15VENDOREXPERIMENT=1
export GOPATH:=$(PWD)/vendor:$(GOPATH)
export GOOS=darwin
export CGO=1

all: build-app

build: format build-app

build-app:
	go build -o $(APP) .

format:
	@gofmt -s -w .

clean:
	@rm $(APP)

.PHONY: all build build-app format clean
