"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.options = void 0;
exports.default = default_1;
var http_1 = require("k6/http");
var k6_1 = require("k6");
exports.options = {
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
function default_1() {
    var url = 'http://jester-api:8080/jester/healthcheck';
    var params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };
    var res = http_1.default.get(url, params);
    (0, k6_1.check)(res, {
        'status is 200': function (r) { return r.status === 200; },
    });
    (0, k6_1.sleep)(1);
}
