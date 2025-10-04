# Makefile para awz-buddy
# ----------------------------------------
# Variables
BINARY_NAME := awz-buddy
OUTPUT_DIR := dist

GO := go
SILENT := @

LINUX := linux
WINDOWS := windows
DARWIN := darwin

AMD64 := amd64
ARM64 := arm64

BINARIES := \
	$(OUTPUT_DIR)/$(BINARY_NAME)-linux-amd64 \
	$(OUTPUT_DIR)/$(BINARY_NAME)-windows-amd64.exe \
	$(OUTPUT_DIR)/$(BINARY_NAME)-darwin-amd64 \
	$(OUTPUT_DIR)/$(BINARY_NAME)-darwin-arm64

# ----------------------------------------
# Default target
.PHONY: all
all: build

$(OUTPUT_DIR):
	$(SILENT)mkdir -p $(OUTPUT_DIR)


.PHONY: build
build: $(OUTPUT_DIR) $(BINARIES)
	$(SILENT)echo "Build completed in $(OUTPUT_DIR)/"

# Linux AMD64
$(OUTPUT_DIR)/$(BINARY_NAME)-linux-amd64: main.go
	$(SILENT)GOOS=$(LINUX) GOARCH=$(AMD64) $(GO) build -o $@ main.go

# Windows AMD64
$(OUTPUT_DIR)/$(BINARY_NAME)-windows-amd64.exe: main.go
	$(SILENT)GOOS=$(WINDOWS) GOARCH=$(AMD64) $(GO) build -o $@ main.go

# macOS Intel AMD64
$(OUTPUT_DIR)/$(BINARY_NAME)-darwin-amd64: main.go
	$(SILENT)GOOS=$(DARWIN) GOARCH=$(AMD64) $(GO) build -o $@ main.go

# macOS ARM64 (Apple Silicon)
$(OUTPUT_DIR)/$(BINARY_NAME)-darwin-arm64: main.go
	$(SILENT)GOOS=$(DARWIN) GOARCH=$(ARM64) $(GO) build -o $@ main.go


.PHONY: clean
clean:
	$(SILENT)rm -rf $(OUTPUT_DIR)
	$(SILENT)echo "Carpeta $(OUTPUT_DIR) eliminada."
