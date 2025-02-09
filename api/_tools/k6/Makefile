# =================================================================
# K6
# =================================================================
.PHONY: k6-build k6-test-run k6-scenario1-run k6-end k6-down

k6-build:
	docker build --no-cache \
		-t k6-load-test:scenario-1 \
		-f ./build/Dockerfile.k6 .

k6-test-run: k6-build
	docker run --rm \
		-e SHOP_API_URL="http://localhost:8080" \
		-e SCRIPT_PATH="/k6/src/scripts/main.ts" \
		--network host \
		--name scenario-1 \
		k6-load-test:scenario-1

k6-scenario1-run: k6-build
	docker run --rm \
		-e SHOP_API_URL="http://localhost:8080" \
		-e SCRIPT_PATH="/k6/src/scenarios/scenario1_user_purchase_flow/main.ts" \
		--network host \
		--name scenario-1 \
		k6-load-test:scenario-1

k6-end:
	docker stop scenario-1

k6-down: k6-end
	docker rm scenario-1
