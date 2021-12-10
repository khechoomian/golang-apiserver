#Golang
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOFMT=$(GOCMD) fmt
GOMOD=$(GOCMD) mod
mainFile="./cmd/golang-apiserver"

go-fmt:
	$(GOFMT) ./...

go-package:
	$(GOMOD) download
	$(GOMOD) tidy
	$(GOMOD) verify

run-linux:
	$(MAKE) go-fmt
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GORUN) ${mainFile}

build-linux:
	$(MAKE) go-fmt
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -ldflags="-w -s" ${mainFile}

build-docker:
	docker build --rm -f "build/docker/Dockerfile" -t golang/apiserver:v1 .

docker-compose:
	$(MAKE) go-fmt
	docker-compose up -d --build