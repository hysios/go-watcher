PROG_NAME = go-watcher

PROJECT ?= github.com/vterdunov/${PROG_NAME}

GO_VARS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
GO_LDFLAGS :="
GO_LDFLAGS += -s -w
GO_LDFLAGS +="


GOLANGCI_LINTER_VERSION = v1.10.1

all: lint test

lint:
	@echo Linting...
	@docker run -it --rm -v $(CURDIR):/go/src/$(PROJECT) -w /go/src/$(PROJECT) golangci/golangci-lint:$(GOLANGCI_LINTER_VERSION) run

start:
	@go run -race ./cmd/watcher/main.go

start-binary: compile
	./$(PROG_NAME)

compile: clean
	@$(GO_VARS) go build -v -ldflags $(GO_LDFLAGS) -o $(PROG_NAME) ./cmd/watcher/main.go

clean:
	@rm -f ${PROG_NAME}

dep:
	@dep ensure -v

test:
	@go test -v -race ./...
