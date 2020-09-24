.PHONY: opa-up
opa-up: ## Starts OPA Server
	@docker run --rm -p 8181:8181 --name gopa -d openpolicyagent/opa run --server --log-level debug

.PHONY: opa-down
opa-down: ## Stops OPA Server
	@docker kill gopa

.PHONY: test
test: ##
	@go test ./...
