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
# 生存確認API
.PHONY: healthcheck
healthcheck:
	curl -i -X 'GET' \
		'http://localhost:8080/shop/healthcheck' \
		-H 'accept: application/json'

# ユーザ系API
.PHONY: create-user get-me create-profile get-profile-me
create-user:
	curl -i -X 'POST' \
		'http://localhost:8080/shop/v1/users' \
		-H 'accept: application/json' \
		-d ''

get-me:
	curl -i -X 'GET' \
		'http://localhost:8080/shop/v1/users/me' \
		-H 'accept: application/json'

create-profile:
	curl -sX 'POST' \
		'http://localhost:8080/shop/v1/users/profiles' \
		-H 'accept: application/json' \
		-H 'Content-Type: application/json' \
		-d '{"name": {"first_name": "世一","last_name": "潔","first_name_roman": "Yoichi","last_name_roman": "Isagi"}, "address": { "region": "関東", "zip_code": "150-8377", "prefecture": "東京都", "city": "渋谷区", "street": "宇田川町", "other": "15番1号"}, "birthdate": "2000-01-01", "image_url": "https://example"}' | jq .

get-profile-me:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/users/profiles/me' \
		-H 'accept: application/json' | jq .

# 商品系API
.PHONY: get-products get-product-by-id
get-products:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/products?cursor=MjAwMDEwMDE%3D&limit=10' \
		-H 'accept: application/json' | jq .

get-product-by-id:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/products/20001001' \
		-H 'accept: application/json' | jq .

# カード系API
.PHONY: get-credit-cards create-credit-card delete-credit-card
get-credit-cards:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/payments/cards' \
		-H 'accept: application/json' | jq .

create-credit-card:
	curl -i -sX 'POST' \
		'http://localhost:8080/shop/v1/payments/cards' \
		-H 'accept: */*' \
		-H 'Content-Type: application/json' \
		-d '{"token": "9f9213b2-9687-6910-d2ec-3a247582be2d", "is_default": true}'

delete-credit-card:
	curl -i -sX 'DELETE' \
		'http://localhost:8080/shop/v1/payments/cards' \
		-H 'accept: */*'

# 予約系API
.PHONY: create-reservations
create-reservations:
	curl -sX 'POST' \
		'http://localhost:8080/shop/v1/payments/reservations' \
		-H 'accept: application/json' \
		-H 'Content-Type: application/json' \
		-d '[{"product_id": 10001001, "quantity": 2}, {"product_id": 10001002, "quantity": 3}]' | jq .

# 注文系API
.PHONY: create-charge get-charge-histories
create-charge:
	curl -i -sX 'POST' \
		'http://localhost:8080/shop/v1/payments/charges' \
		-H 'accept: */*' \
		-H 'Content-Type: application/json' \
		-d '{"reservation_id": "9f9213b2-9687-6910-d2ec-3a247582be2d"}'

get-charge-histories:
	curl -sX 'GET' \
		'http://localhost:8080/shop/v1/payments/charges/histories?offset=0&limit=10' \
		-H 'accept: application/json' | jq .
