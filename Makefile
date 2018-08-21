PROG_NAME = go-watcher

PROJECT ?= github.com/vterdunov/${PROG_NAME}

GOLANGCI_LINTER_VERSION = v1.10

lint:
	@echo Linting...
	@docker run -it --rm -v $(CURDIR):/go/src/$(PROJECT) -w /go/src/$(PROJECT) golangci/golangci-lint:$(GOLANGCI_LINTER_VERSION) run
