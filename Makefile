.SILENT:

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## test: run tests
test:
ifdef year
ifdef day
	go test -v github.com/teodorpopa/advent-of-code-go/y$(year)/day$(day) -run TestDayPart1
	go test -v github.com/teodorpopa/advent-of-code-go/y$(year)/day$(day) -run TestDayPart2
else
	go test -v ./...
endif
endif

## run: run specific day
.PHONY: run
run:
	@go run main.go --year=$(year) --day=$(day)


## create: create a day from template
.PHONY: create
create:
	@go run main.go --create --year=$(year) --day=$(day)