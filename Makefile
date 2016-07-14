APP=doenter
REPO=fntlnz/doenter

export GO15VENDOREXPERIMENT=1

all: build-container

build: deps format build-app

build-app:
	GOOS=darwin CGO=1 go build -o $(APP) .

build-container:
	@docker build -t doenter-build -f Dockerfile.build .
	@docker run -it -e BUILD -e TAG --name doenter-build -ti doenter-build make build
	@docker cp doenter-build:/go/src/github.com/$(REPO)/$(APP) ./$(APP)
	@docker rm -fv doenter-build

deps:
	@go get -u github.com/FiloSottile/gvt
	@gvt restore

format:
	@gofmt -s -w .

clean:
	@docker rmi doenter-build
	@rm $(APP)

.PHONY: all build build-app build-container format clean
