# ==================================================================================== #
# HELPERS
# ==================================================================================== #
.SILENT:

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go test -vet=off ./...

## test: run tests
.PHONY: test
test:
ifdef year
ifdef day
	go test -v github.com/teodorpopa/advent-of-code-go/y$(year) -run TestDay$(day)First
	go test -v github.com/teodorpopa/advent-of-code-go/y$(year) -run TestDay$(day)Second
else
	go test -v ./...
endif
endif

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## run: run the application
.PHONY: run
run:
	@go run main.go $(year) $(day)


## create: create a day from template
.PHONY: create
create:
	@go run main.go create $(year) $(day)
