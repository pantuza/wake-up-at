
#
# Makefile of the wake-up-at project
#


MAIN_FILE := wake-up-at.go

BINDIR := /usr/local/bin
bindir := $(BINDIR) # Makefile compatibility with generic bindir variable
BINFILE := wake-up-at


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
