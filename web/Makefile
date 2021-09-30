.PHONY: generate
generate:
	oapi-codegen ../reference/openapi.yml > ./Openapi/interface.gen.go

.PHONY: generate_client
generate_client:
	oapi-codegen -generate "client" ../reference/openapi.yml > ./Openapi/client.gen.go

.PHONY: generate_types
generate_types:
	oapi-codegen -generate "types" ../reference/openapi.yml > ./Openapi/types.gen.go

.PHONY: generate_server
generate_server:
	oapi-codegen -generate "server" ../reference/openapi.yml > ./Openapi/server.gen.go

.PHONY: generate_all
generate_all: generate_client generate_types generate_server

.PHONY: integration_tests
integration_tests: ENV_FILE := integration_tests.env
integration_tests:
	ENV_FILE="${ENV_FILE}" \
	./tools/integration_tests.sh
