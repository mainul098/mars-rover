NAME=mars-exploration
REPO=github.com/mainul098/mars-rover
SRC_DIRS=rover

# Colors
NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

.PHONY: all deps

all: deps

# Install dependencies
deps:
	git config --global http.https://gopkg.in.followRedirects true
	@printf "$(OK_COLOR)==> Ensuring dependencies$(NO_COLOR)\n"
	@go mod vendor

# Unit tests
test-unit:
	@printf "$(OK_COLOR)==> Running unit tests$(NO_COLOR)\n"
	@CGO_ENABLED=1 go test -race -coverprofile=coverage_unit.txt -covermode=atomic -v $$(for d in $(SRC_DIRS); do echo ./$$d/...; done)
