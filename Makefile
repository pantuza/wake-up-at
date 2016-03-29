
#
# Makefile of the wake-up-at project
#


MAIN_FILE := wake-up-at.go


build:
	@echo "Building project"
	@go build

run:
	@clear
	@go run $(MAIN_FILE)

test:
	@clear
	@go test ./... -v

clean:
	@rm -rvf wake-up-at
