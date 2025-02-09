import * as http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = __ENV.SHOP_API_URL || 'http://localhost:8080';
const headers = { 'Content-Type': 'application/json' };

export const options = {
  stages: [
    { duration: '15s', target: 10 },  // 15秒間で10人のユーザに増加
    { duration: '20s', target: 10 },  // 20秒間10人で維持
    { duration: '10s', target: 0 },   // 10秒間で0人に減少
  ],
  thresholds: {
    'http_req_duration': ['p(99) < 3000'],  // 99%のリクエストが3秒以内に終了
  },
};

export default function () {
  // ユーザー登録
  let res = http.post(`${BASE_URL}/shop/v1/users`, JSON.stringify({ /* 登録データ */ }), { headers });
  check(res, { 'ユーザー登録成功': (r) => r.status === 201 });
  sleep(1);

  // ログイン
  res = http.get(`${BASE_URL}/shop/v1/users/me`, { headers });
  check(res, { 'ログイン成功': (r) => r.status === 200 });
  sleep(1);

  // プロフィール登録
  res = http.post(`${BASE_URL}/shop/v1/users/profiles`, JSON.stringify({ /* プロフィールデータ */ }), { headers });
  check(res, { 'プロフィール登録成功': (r) => r.status === 201 });
  sleep(1);

  // プロフィール取得
  res = http.get(`${BASE_URL}/shop/v1/users/profiles/me`, { headers });
  check(res, { 'プロフィール取得成功': (r) => r.status === 200 });
  sleep(1);

  // 商品一覧取得
  res = http.get(`${BASE_URL}/shop/v1/products`, { headers });
  check(res, { '商品一覧取得成功': (r) => r.status === 200 });
  sleep(1);

  // 商品詳細ページ
  const productID = 20001001;
  res = http.get(`${BASE_URL}/shop/v1/products/${productID}`, { headers });
  check(res, { '商品詳細ページ取得成功': (r) => r.status === 200 });
  sleep(1);

  // クレジットカード取得
  res = http.get(`${BASE_URL}/shop/v1/payments/cards`, { headers });
  check(res, { 'クレジットカード情報取得成功': (r) => r.status === 200 });
  sleep(1);

  // クレジットカード登録
  res = http.post(`${BASE_URL}/shop/v1/payments/cards`, JSON.stringify({ /* クレジットカード情報 */ }), { headers });
  check(res, { 'クレジットカード登録成功': (r) => r.status === 204 });
  sleep(1);

  // 予約リクエスト
  res = http.post(`${BASE_URL}/shop/v1/payments/reservations`, JSON.stringify([{ product_id: productID, quantity: 1 }]), { headers });
  check(res, { '予約リクエスト成功': (r) => r.status === 201 });
  sleep(1);

  // 決済
  res = http.post(`${BASE_URL}/shop/v1/payments/charges`, JSON.stringify({ reservation_id: '予約ID' }), { headers });
  check(res, { '決済成功': (r) => r.status === 204 });
  sleep(1);

  // 決済履歴取得
  res = http.get(`${BASE_URL}/shop/v1/payments/charges/histories`, { headers });
  check(res, { '決済履歴取得成功': (r) => r.status === 200 });
  sleep(1);
}
