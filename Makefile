BINDIR := ./bin
TARGET := autocomplete
ARGS := # Used in prog exec

all: build

build:
	@echo "Building..."
	mkdir -p $(BINDIR)
	go build -o $(BINDIR) ./cmd/...

run:
	$(BINDIR)/$(TARGET) $(ARGS)

clean:
	$(RM) -r $(BINDIR)

PHONY: all build run clean
