BINDIR := ./bin
TARGET := autocomplete
ARGS := # Used in prog exec

all: build

build:
	@echo "Building..."
	mkdir -p $(BINDIR)
	go build -o $(BINDIR) ./cmd/...

run:
	@echo "Running..."
	$(BINDIR)/$(TARGET) $(ARGS)

test:
	@echo "Testing..."
	go test -cover ./internal/...

clean:
	@echo "Cleaning..."
	$(RM) -r $(BINDIR)

PHONY: all build run clean
