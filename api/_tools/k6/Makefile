# =================================================================
# K6
# =================================================================
.PHONY: k6-build k6-run-test k6-run-scenario1 k6-end k6-down

k6-build:
	docker build --no-cache \
		-t k6-load-test \
		-f ./build/Dockerfile.k6 .

k6-run-test: k6-build
	docker run --rm \
		-e SHOP_API_URL="http://localhost:8080" \
		-e SCRIPT_PATH="/k6/src/scripts/main.ts" \
		--name k6-load-test \
		--network host \
		k6-load-test

k6-run-scenario1: k6-build
	docker run --rm \
		-e SHOP_API_URL="http://localhost:8080" \
		-e SCRIPT_PATH="/k6/src/scenarios/scenario1_user_purchase_flow/main.ts" \
		--name k6-load-test \
		--network host \
		k6-load-test

k6-end:
	docker stop k6-load-test

k6-down: k6-end
	docker rm k6-load-test
