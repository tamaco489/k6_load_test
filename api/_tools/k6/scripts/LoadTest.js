// NOTE: 諸々終わったらTypeScriptに書き換える

import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
  // vus: 10,
  // duration: '10s',

  thresholds: {
    // Assert that 99% of requests finish within 3000ms.
    http_req_duration: ["p(99) < 3000"],
  },
  // Ramp the number of virtual users up and down
  stages: [
    { duration: "15s", target: 10 },
    { duration: "20s", target: 10 },
    { duration: "10s", target: 0 },
  ],
};

export default function () {
  const url = 'http://jester-api:8080/jester/healthcheck';
  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  const res = http.get(url, params);
  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  sleep(1);
}
