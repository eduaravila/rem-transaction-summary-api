.PHONY: openapi_http
openapi_http: 
	@./scripts/openapi-http.sh summary ports pkg/summary/ports

test:
	@./scripts/test.sh summary .env.test .env 