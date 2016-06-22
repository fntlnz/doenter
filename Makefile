APP=doenter
REPO=fntlnz/doenter

export GO15VENDOREXPERIMENT=1
export GOPATH:=$(PWD)/vendor:$(GOPATH)
export GOOS=darwin
export CGO=1

all: build-container

build: format build-app

build-app:
	go build -o $(APP) .

build-container:
	@docker build -t doenter-build -f Dockerfile.build .
	@docker run -it -e BUILD -e TAG --name doenter-build -ti doenter-build make build
	@docker cp doenter-build:/go/src/github.com/$(REPO)/$(APP) ./$(APP)
	@docker rm -fv doenter-build

format:
	@gofmt -s -w .

clean:
	@rm $(APP)

.PHONY: all build build-app build-container format clean
