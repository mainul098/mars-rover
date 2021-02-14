NAME=mars-exploration
REPO=github.com/mainul098/mars-rover
SRC_DIRS=rover

# Unit tests
test-unit:
	@printf "$(OK_COLOR)==> Running unit tests$(NO_COLOR)\n"
	@CGO_ENABLED=1 go test -race -coverprofile=coverage_unit.txt -covermode=atomic -v $$(for d in $(SRC_DIRS); do echo ./$$d/...; done)