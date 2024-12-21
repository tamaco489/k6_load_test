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
.PHONY: healthcheck create-user get-me create-profile get-profile-me

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

create-profile:
	curl -sX 'POST' \
		'http://localhost:8080/jester/v1/users/profiles' \
		-H 'accept: application/json' \
		-H 'Content-Type: application/json' \
		-d '{"name": {"first_name": "世一","last_name": "潔","first_name_roman": "Yoichi","last_name_roman": "Isagi"}, "address": { "region": "関東", "zip_code": "150-8377", "prefecture": "東京都", "city": "渋谷区", "street": "宇田川町", "other": "15番1号"}, "birthdate": "2000-01-01", "image_url": "https://example"}' | jq .

get-profile-me:
	curl -sX 'GET' \
		'http://localhost:8080/jester/v1/users/profiles/me' \
		-H 'accept: application/json' | jq .
