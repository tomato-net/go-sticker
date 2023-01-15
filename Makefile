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

.PHONY: deploy-published
deploy-published: set-image
	$(KUBECTL) apply -k deployments/overlays/published

.PHONY: deploy-local
deploy-local: docker-build set-image minikube-load
	$(KUBECTL) apply -k deployments/overlays/local


.PHONY: minikube-load
minikube-load:
	minikube image load $(IMG)
