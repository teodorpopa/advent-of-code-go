.SILENT:

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## test: run tests
test:
ifdef year
ifdef day
	go test -v github.com/teodorpopa/advent-of-code-go/y$(year)/day$(day)
else
	go test -v ./...
endif
endif

## run: run specific day
.PHONY: run
run:
ifdef year
ifdef day
	@go run y$(year)/day$(day)/main.go --part=$(part)
endif
endif

## scaffold: scaffold AoC day
.PHONY: scaffold
scaffold:
ifdef year
ifdef day
	@go run main.go scaffold --year=$(year) --day=$(day)
endif
endif

## calendar: view the AoC calendar for the specified year
.PHONY: calendar
calendar:
ifdef year
	@go run main.go calendar --year=$(year)
endif

## view: view an AoC puzzle
.PHONY: view
view:
ifdef year
	@go run main.go view --year=$(year) --day=$(day)
endif