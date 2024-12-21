# =================================================================
# K6
# =================================================================
.PHONY: k6-compile k6-run k6-end k6-down

k6-compile:
	cd ./_tools/k6/ && npx tsc ./src/scripts/main.ts --outDir ./dist && cd -

k6-run: k6-compile
	docker-compose up k6

k6-end:
	docker-compose stop k6

k6-down: k6-end
	docker rm k6


# =================================================================
# API Request
# =================================================================
.PHONY: healthcheck create-user get-me

healthcheck:
	curl -i -X 'GET' \
		'http://localhost:8080/jester/healthcheck' \
		-H 'accept: application/json'

create-user:
	curl -i -X 'POST' \
		'http://localhost:8080/jester/v1/users' \
		-H 'accept: application/json' \
		-d ''

get-me:
	curl -i -X 'GET' \
		'http://localhost:8080/jester/v1/users/me' \
		-H 'accept: application/json'
