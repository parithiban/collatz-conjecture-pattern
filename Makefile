DOCKER_PREFIX = parithiban
DOCKER_REPO = collatz

ifeq ($(REQUIRE_APPROVAL), )
REQUIRE_APPROVAL=any-change
endif

.DEFAULT_GOAL := explain
.PHONY: explain
explain:
	### Welcome
	#
	#
	#	
	#	  _______  _______  _        _        _______ _________ _______ 
	#	 (  ____ \(  ___  )( \      ( \      (  ___  )\__   __// ___   )
	#	 | (    \/| (   ) || (      | (      | (   ) |   ) (   \/   )  |
	#	 | |      | |   | || |      | |      | (___) |   | |       /   )
	#	 | |      | |   | || |      | |      |  ___  |   | |      /   / 
	#	 | |      | |   | || |      | |      | (   ) |   | |     /   /  
	#	 | (____/\| (___) || (____/\| (____/\| )   ( |   | |    /   (_/\
	#	 (_______/(_______)(_______/(_______/|/     \|   )_(   (_______/
	#	 
	#	
	#
	# $$ make all
	#
	### Targets
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: ## Install the local dependencies
	@echo "Getting go tools"
	go get golang.org/x/lint/golint
	go get github.com/securego/gosec/cmd/gosec

.PHONY: vet
vet: ## Vet the code
	go vet ./...

.PHONY: security-check
security-check: ## Inspect code for security vulnerabilities
	gosec ./...

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_PREFIX)/$(DOCKER_REPO) .

.PHONY: docker-push
docker-push:
	docker push $(DOCKER_PREFIX)/$(DOCKER_REPO):latest
