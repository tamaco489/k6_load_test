# k6

### 1. 実行環境整備

```bash
$ pwd
/<your-project>/k6_load_test/api
```

k6をTypeScriptで実行しているため、必要なモジュールをインストールする
```bash
npm -i
```

### 2. ディレクトリ構成

```bash
$ tree -I "node_modules"
.
├── README.md
├── dist
│   └── main.js
├── package-lock.json
├── package.json
├── src
│   ├── @types
│   ├── config
│   ├── scenarios
│   └── scripts
│       └── main.ts
└── tsconfig.json

6 directories, 6 files
```

### 3. 負荷テストを実行する

```bash
# テスト実行（healthcheckのみ）
make k6-run-test
```

### 4. シナリオを定義して負荷テストを実行する

#### 🟢 負荷試験シナリオ (1): 新規ユーザーの購買フロー

| 行動 | HTTP Method | エンドポイント |
|------|------------|---------------|
| ユーザー登録 | `POST` | `/v1/users` |
| ログイン | `GET` | `/v1/users/me` |
| プロフィール登録 | `POST` | `/v1/users/profiles` |
| プロフィール登録 | `GET` | `/v1/users/profiles/me` |
| 商品一覧を取得 | `GET` | `/v1/products` |
| 商品詳細ページを見る | `GET` | `/v1/products/{productID}` |
| クレジットカード取得 | `GET` | `/v1/payments/cards` |
| クレジットカード登録 | `POST` | `/v1/payments/cards` |
| 予約リクエスト (カートに追加) | `POST` | `/v1/payments/reservations` |
| 決済 | `POST` | `/v1/payments/charges` |
| 決済履歴取得 | `GET` | `/v1/payments/charges/histories` |

---

#### 🟡 負荷試験シナリオ (2): 既存ユーザーの通常購買フロー

| 行動 | HTTP Method | エンドポイント |
|------|------------|---------------|
| ログイン | `GET` | `/v1/users/me` |
| 商品一覧を取得 | `GET` | `/v1/products` |
| 商品詳細ページを見る | `GET` | `/v1/products/{productID}` |
| 予約リクエスト (カートに追加) | `POST` | `/v1/payments/reservations` |
| 決済 | `POST` | `/v1/payments/charges` |
| 決済履歴取得 | `GET` | `/v1/payments/charges/histories` |

---

#### 🔴 負荷試験シナリオ (3): 高負荷テスト (API のスケール耐性を確認)

| 行動 | HTTP Method | エンドポイント |
|------|------------|---------------|
| **1000ユーザーがログイン & 商品一覧取得** | `GET` | `/v1/users/me`, `/v1/products` |
| **500ユーザーがカートに商品を追加** | `POST` | `/v1/payments/reservations` |
| **200ユーザーが決済処理に進む** | `POST` | `/v1/payments/charges` |
| **継続的に決済履歴を取得** | `GET` | `/v1/payments/charges/histories` |

---

```bash
# シナリオに基づいた負荷テストを実行
docker run --rm \
        -e SHOP_API_URL="http://localhost:8080" \
        -e SCRIPT_PATH="/k6/src/scenarios/scenario1_user_purchase_flow/main.ts" \
        --name k6-load-test \
        --network host \
        k6-load-test

         /\      Grafana   /‾‾/
    /\  /  \     |\  __   /  /
   /  \/    \    | |/ /  /   ‾‾\
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/

     execution: local
        script: /k6/src/scenarios/scenario1_user_purchase_flow/main.ts
        output: -

     scenarios: (100.00%) 1 scenario, 10 max VUs, 1m15s max duration (incl. graceful stop):
              * default: Up to 10 looping VUs for 45s over 3 stages (gracefulRampDown: 30s, gracefulStop: 30s)


running (0m01.0s), 01/10 VUs, 0 complete and 0 interrupted iterations
default   [   2% ] 01/10 VUs  01.0s/45.0s

running (0m02.0s), 02/10 VUs, 0 complete and 0 interrupted iterations
default   [   4% ] 02/10 VUs  02.0s/45.0s

# 省略

running (0m59.0s), 01/10 VUs, 26 complete and 0 interrupted iterations
default ↓ [ 100% ] 06/10 VUs  45s

     ✓ ユーザー登録成功
     ✓ ログイン成功
     ✓ プロフィール登録成功
     ✓ プロフィール取得成功
     ✓ 商品一覧取得成功
     ✓ 商品詳細ページ取得成功
     ✓ クレジットカード情報取得成功
     ✓ クレジットカード登録成功
     ✓ 予約リクエスト成功
     ✓ 決済成功
     ✓ 決済履歴取得成功

     checks.........................: 100.00% 297 out of 297
     data_received..................: 134 kB  2.3 kB/s
     data_sent......................: 45 kB   753 B/s
     http_req_blocked...............: avg=34.54µs  min=2.93µs   med=8.55µs   max=1.08ms   p(90)=16.55µs  p(95)=25.29µs
     http_req_connecting............: avg=17µs     min=0s       med=0s       max=814.92µs p(90)=0s       p(95)=0s
   ✓ http_req_duration..............: avg=347.08ms min=476.14µs med=1.9ms    max=2s       p(90)=1s       p(95)=2s
       { expected_response:true }...: avg=347.08ms min=476.14µs med=1.9ms    max=2s       p(90)=1s       p(95)=2s
     http_req_failed................: 0.00%   0 out of 297
     http_req_receiving.............: avg=124.59µs min=31.06µs  med=103.85µs max=347.11µs p(90)=215.38µs p(95)=258.91µs
     http_req_sending...............: avg=74.5µs   min=7.57µs   med=54.37µs  max=378.1µs  p(90)=143.02µs p(95)=160.46µs
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s       p(90)=0s       p(95)=0s
     http_req_waiting...............: avg=346.88ms min=390.6µs  med=1.63ms   max=2s       p(90)=1s       p(95)=2s
     http_reqs......................: 297     5.006843/s
     iteration_duration.............: avg=14.82s   min=14.82s   med=14.83s   max=14.83s   p(90)=14.83s   p(95)=14.83s
     iterations.....................: 27      0.455168/s
     vus............................: 1       min=1          max=10
     vus_max........................: 10      min=10         max=10


running (0m59.3s), 00/10 VUs, 27 complete and 0 interrupted iterations
default ✓ [ 100% ] 00/10 VUs  45s
```
