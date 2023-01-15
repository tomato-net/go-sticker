ENV ?= published
IMG ?= tomatod4r/stock-ticker

KUBECTL = "kubectl"
KUSTOMIZE = "kustomize"

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: docker-build
docker-build: ## Build docker image with the manager.
	docker build -t ${IMG} -f build/server/Dockerfile .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker push ${IMG}

.PHONY: set-image
set-image:
	cd deployments/base && $(KUSTOMIZE) edit set image server=$(IMG)

.PHONY: deploy
deploy:
	$(KUBECTL) apply -k deployments/overlays/$(ENV)