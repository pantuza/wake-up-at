
#
# Makefile of the wake-up-at project
#


MAIN_FILE := wake-up-at.go

BINDIR := /usr/local/bin
bindir := $(BINDIR) # Makefile compatibility with generic bindir variable
BINFILE := wake-up-at


help:
	@echo "wake-up-at"
	@echo
	@echo "Target rules:"
	@echo
	@echo "    build - Builds the project and creates the binary file"
	@echo "    run - Runs the project"
	@echo "    test - Runs project tests"
	@echo "    format - Runs go fmt tool at source files"
	@echo "    clean - Remove binary file"
	@echo "    install - Installs binary file on system path"
	@echo "    uninstall - Uninstall program from system"

build:
	@echo "Building project"
	@go build

run:
	@clear
	@go run $(MAIN_FILE)

test:
	@clear
	@go test ./... -v

format:
	@go fmt .

clean:
	@rm -rvf wake-up-at

install: $(BINFILE)
	@echo "Installing project"
	@install -v $(BINFILE) $(BINDIR)/$(BINFILE)

uninstall: $(BINDIR)/$(BINFILE)
	@echo "Uninstalling project"
	@rm -vf $+
