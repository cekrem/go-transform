# Go commands
GO := go
GOBUILD := $(GO) build
GOCLEAN := $(GO) clean

# Directories
BUILD_DIR := build
PLUGIN_DIR := plugins
CMD_DIR := cmd

.PHONY: all
all: build plugins

.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/transform $(CMD_DIR)/main.go

.PHONY: plugins
plugins: build-plugins

.PHONY: build-plugins
build-plugins:
	@mkdir -p $(BUILD_DIR)/plugins
	@echo "Building plugins..."
	@for plugin in $(PLUGIN_DIR)/*/ ; do \
		if [ -f $$plugin/go.mod ]; then \
			plugin_name=$$(basename $$plugin); \
			echo "Building plugin: $$plugin_name"; \
			cd $$plugin && go mod tidy && \
			$(GOBUILD) -buildmode=plugin -o ../../$(BUILD_DIR)/plugins/$$plugin_name.so || exit 1; \
			cd ../../; \
		fi \
	done

.PHONY: clean
clean:
	@rm -rf $(BUILD_DIR)
	$(GOCLEAN) -cache

.PHONY: test
test:
	$(GO) test ./processor/... ./interfaces/...
	@cd $(PLUGIN_DIR)/passthrough && $(GO) test ./... 

.PHONY: plugin
plugin:
	@if [ "$(filter-out $@,$(MAKECMDGOALS))" = "" ]; then \
		echo "Error: Plugin name not specified. Usage: make plugin <plugin-name>", or make plugins to build all; \
		exit 1; \
	fi
	@mkdir -p $(BUILD_DIR)/plugins
	@plugin_name=$(filter-out $@,$(MAKECMDGOALS)); \
	echo "Building plugin: $$plugin_name"; \
	if [ ! -d "$(PLUGIN_DIR)/$$plugin_name" ]; then \
		echo "Error: Plugin '$$plugin_name' not found in $(PLUGIN_DIR)"; \
		exit 1; \
	fi; \
	cd $(PLUGIN_DIR)/$$plugin_name && go mod tidy && \
	$(GOBUILD) -buildmode=plugin -o ../../$(BUILD_DIR)/plugins/$$plugin_name.so || exit 1

%:
	@: 