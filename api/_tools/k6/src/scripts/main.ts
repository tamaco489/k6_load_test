import * as http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  // vus: 10,
  // duration: '10s',

  thresholds: {
    // 99%のリクエストが3000ms以内に終了することを保証する。
    http_req_duration: ['p(99) < 3000'],
  },

  // 最初の15秒で0人から10人に増やし、次の20秒で10人を維持、最後の10秒で0人に減らしていく。
  stages: [
    { duration: '15s', target: 10 },
    { duration: '20s', target: 10 },
    { duration: '10s', target: 0 },
  ],
};

export default function () {
  const BASE_URL = __ENV.SHOP_API_URL || 'http://localhost:8080';
  const headers = { 'Content-Type': 'application/json' };

  let res = http.get(`${BASE_URL}/shop/healthcheck`, { headers });
  check(res, {'status is health': (r) => r.status === 200 });
  sleep(1);
}
