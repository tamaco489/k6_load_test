# k6

### 実行環境整備

```bash
$ pwd
/<your-project>/k6_load_test/api
```

k6をTypeScriptで実行しているため、必要なモジュールをインストールする
```bash
npm -i
```

負荷テストを実行する
```bash
make k6-run
```

### 実行例

```bash
$ make k6-run
cd ./_tools/k6/ && npx tsc ./src/scripts/main.ts --outDir ./dist && cd -
/home/miyabiii1210/src/github.com/tamaco489/k6_load_test/api
docker-compose up k6
Pulling k6 (grafana/k6:latest)...
latest: Pulling from grafana/k6
da9db072f522: Pull complete
48d9e22b0142: Pull complete
cfcefb3cfb8f: Pull complete
4f4fb700ef54: Pull complete
Digest: sha256:c362e72377bb63a897346d549b5ecc17bddd5ad5d3458c0307e03873ab2807c7
Status: Downloaded newer image for grafana/k6:latest
Creating k6 ... done
Attaching to k6
k6            |
k6            |          /\      Grafana   /‾‾/
k6            |     /\  /  \     |\  __   /  /
k6            |    /  \/    \    | |/ /  /   ‾‾\
k6            |   /          \   |   (  |  (‾)  |
k6            |  / __________ \  |_|\_\  \_____/
k6            |
k6            |      execution: local
k6            |         script: /_tools/k6/dist/main.js
k6            |         output: -
k6            |
k6            |      scenarios: (100.00%) 1 scenario, 10 max VUs, 1m15s max duration (incl. graceful stop):
k6            |               * default: Up to 10 looping VUs for 45s over 3 stages (gracefulRampDown: 30s, gracefulStop: 30s)
k6            |
k6            |
k6            | running (0m01.0s), 01/10 VUs, 0 complete and 0 interrupted iterations
k6            | default   [   2% ] 01/10 VUs  01.0s/45.0s

... 省略

k6            | running (0m45.0s), 01/10 VUs, 332 complete and 0 interrupted iterations
k6            | default   [ 100% ] 01/10 VUs  45.0s/45.0s
k6            |
k6            |      ✓ status is 200
k6            |
k6            |      checks.........................: 100.00% 333 out of 333
k6            |      data_received..................: 42 kB   924 B/s
k6            |      data_sent......................: 44 kB   968 B/s
k6            |      http_req_blocked...............: avg=26.38µs  min=2.27µs   med=5.32µs   max=1.03ms   p(90)=13.47µs  p(95)=21.02µs
k6            |      http_req_connecting............: avg=12.86µs  min=0s       med=0s       max=793.21µs p(90)=0s       p(95)=0s
k6            |    ✓ http_req_duration..............: avg=716.22µs min=257.62µs med=625.2µs  max=1.74ms   p(90)=1.2ms    p(95)=1.35ms
k6            |        { expected_response:true }...: avg=716.22µs min=257.62µs med=625.2µs  max=1.74ms   p(90)=1.2ms    p(95)=1.35ms
k6            |      http_req_failed................: 0.00%   0 out of 333
k6            |      http_req_receiving.............: avg=108.8µs  min=18.1µs   med=83.44µs  max=396.38µs p(90)=206.22µs p(95)=234.35µs
k6            |      http_req_sending...............: avg=27.95µs  min=5.54µs   med=16.33µs  max=286µs    p(90)=49.49µs  p(95)=98.45µs
k6            |      http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s       p(90)=0s       p(95)=0s
k6            |      http_req_waiting...............: avg=579.47µs min=172.9µs  med=472.04µs max=1.57ms   p(90)=1.02ms   p(95)=1.11ms
k6            |      http_reqs......................: 333     7.389834/s
k6            |      iteration_duration.............: avg=1s       min=1s       med=1s       max=1s       p(90)=1s       p(95)=1s
k6            |      iterations.....................: 333     7.389834/s
k6            |      vus............................: 1       min=1          max=10
k6            |      vus_max........................: 10      min=10         max=10
k6            |
k6            |
k6            | running (0m45.1s), 00/10 VUs, 333 complete and 0 interrupted iterations
k6            | default ✓ [ 100% ] 00/10 VUs  45s
k6 exited with code 0
```
