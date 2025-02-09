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
  const url = 'http://shop-api:8080/shop/healthcheck';
  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.get(url, params);

  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  sleep(1);
}
