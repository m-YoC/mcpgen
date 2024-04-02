.DEFAULT_GOAL := help

hello: ## golang hello world test
	cd hello-world && go run hello.go
build: ## build mcpgen (test)
	cd mcpgen_dir && CGO_ENABLED=0 go build -o ../mcpgen -ldflags '-s -w' -trimpath
install: ## mcpgen install (test)
	cd mcpgen_dir && CGO_ENABLED=0 go install -ldflags '-s -w' -trimpath mcpgen.go 
create-yaml: ## create-mode
	cd mcpgen_dir && MCPGEN_CREATE_YAML_MODE= go run main.go

cross-compile: ## cross compile
	@cd mcpgen_dir && GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o ../compiled/linux/mcpgen -ldflags '-s -w' -trimpath
	@cd mcpgen_dir && GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ../compiled/win/mcpgen.exe -ldflags '-s -w' -trimpath
	@cd mcpgen_dir && GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o ../compiled/mac_intel/mcpgen -ldflags '-s -w' -trimpath
	@cd mcpgen_dir && GOOS=darwin  GOARCH=arm64 CGO_ENABLED=0 go build -o ../compiled/mac_arm/mcpgen -ldflags '-s -w' -trimpath

help: ## Display this help screen.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk -F ':.*?## ' '{printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
