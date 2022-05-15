import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  stages: [
    { duration: '5s', target: 10 },
    { duration: '16s', target: 32 },
    { duration: '24s', target: 48 },
    { duration: '10s', target: 20 },
    { duration: '5s', target: 10 },
  ],
  thresholds: {
      http_req_duration: ['p(90) < 200', 'p(95) < 500', 'p(99.9) < 1200'],
    }
};

export default function () {
  let req = http.get('http://localhost:5000/api/v1/citizens/1');
  check(req, { 'status was 200': (r) => r.status == 200 });
}