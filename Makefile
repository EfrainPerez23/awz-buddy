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

# Colors
RED := \033[1;31m
GREEN := \033[1;32m
YELLOW := \033[1;33m
CYAN := \033[1;36m
RESET := \033[0m

# ----------------------------------------
# Default target
.PHONY: all
all: build

# ----------------------------------------
.PHONY: prepare
prepare:
	@if [ -d $(OUTPUT_DIR) ]; then \
		rm -rf $(OUTPUT_DIR); \
		echo "$(CYAN)Folder $(OUTPUT_DIR) cleaned$(RESET)"; \
	fi
	@mkdir -p $(OUTPUT_DIR)

# ----------------------------------------
# Build
.PHONY: build
build: prepare $(BINARIES)
	$(SILENT)echo "$(GREEN)Builds completed in $(OUTPUT_DIR) $(RESET)"

# ----------------------------------------
# Linux AMD64
$(OUTPUT_DIR)/$(BINARY_NAME)-linux-amd64: $(OUTPUT_DIR) main.go
	$(SILENT)GOOS=$(LINUX) GOARCH=$(AMD64) $(GO) build -o $@ main.go

# Windows AMD64
$(OUTPUT_DIR)/$(BINARY_NAME)-windows-amd64.exe: $(OUTPUT_DIR) main.go
	$(SILENT)GOOS=$(WINDOWS) GOARCH=$(AMD64) $(GO) build -o $@ main.go

# macOS Intel AMD64
$(OUTPUT_DIR)/$(BINARY_NAME)-darwin-amd64: $(OUTPUT_DIR) main.go
	$(SILENT)GOOS=$(DARWIN) GOARCH=$(AMD64) $(GO) build -o $@ main.go

# macOS ARM64 (Apple Silicon)
$(OUTPUT_DIR)/$(BINARY_NAME)-darwin-arm64: $(OUTPUT_DIR) main.go
	$(SILENT)GOOS=$(DARWIN) GOARCH=$(ARM64) $(GO) build -o $@ main.go

# ----------------------------------------
# Clean
.PHONY: clean
clean:
	$(SILENT)rm -rf $(OUTPUT_DIR)
	$(SILENT)echo "$(CYAN)Folder $(OUTPUT_DIR) deleted$(RESET)"

# ----------------------------------------
# Nix environment
.PHONY: develop
develop:
	@echo "$(CYAN)🔦🔦 Verifying Nix installation 🔦🔦$(RESET)"
	@if ! command -v nix >/dev/null 2>&1; then \
		echo "$(RED)❌❌ Nix is not installed ❌❌$(RESET)"; \
		echo "$(YELLOW)🏄🏄 You can install it going to $(CYAN)https://nixos.org/download 🏄🏄$(RESET)"; \
		exit 1; \
	fi
	@echo "$(GREEN)✅✅ Nix detected ✅✅$(RESET)"
	@if [ -f flake.nix ]; then \
		echo "$(CYAN)⚗️ ⚗️ Opening environment with flake.nix ⚗️ ⚗️$(RESET)"; \
		nix develop; \
	elif [ -f shell.nix ]; then \
		echo "$(CYAN)⚗️ ⚗️ Opening environment with shell.nix ⚗️ ⚗️$(RESET)"; \
		nix-shell; \
	else \
		echo "$(RED)❌❌ flake.nix or shell.nix not found ❌❌$(RESET)"; \
		exit 1; \
	fi
