
import http from 'k6/http';
import { check, sleep } from 'k6';
import { randomSeed } from 'k6';

export let options = {
  stages: [
    { duration: '5s', target: 1 },
    { duration: '10s', target: 2 },
    { duration: '25s', target: 5 },
    { duration: '20s', target: 1 },
  ],
  thresholds: {
    http_req_duration: ['p(90) < 200', 'p(95) < 500', 'p(99.9) < 1200'],
  }
};

export default function () {
  randomSeed(123456789);
  let rnd = Math.random();
  let data = {
    "name": "string" + rnd,
    "species": "string" + rnd,
    "description": "string" + rnd,
    "has_pet_human": (rnd % 2 == 1),
    "weight": 50.0,
    "height": 1.58,
    "photo_url": "http://animalia-api.com/"+"string" + rnd +".jpg"
  };

  let headers = { 'Content-Type': 'application/json' };

  let resCreate = http.post('http://localhost:5000/api/v1/citizens', JSON.stringify(data), { headers: headers });
  let id = JSON.parse(resCreate.body).id;
  check(resCreate, { 'status was 201': (r) => r.status == 201 });
  
  let resRead = http.get('http://localhost:5000/api/v1/citizens/' + id);
  check(resRead, { 'status was 200': (r) => r.status == 200 });

  let resUpdate = http.put('http://localhost:5000/api/v1/citizens/'+ id, JSON.stringify(data), { headers: headers });
  check(resUpdate, { 'status was 200': (r) => r.status == 200 });
  
  let resDelete = http.del('http://localhost:5000/api/v1/citizens/' + id);
  check(resDelete, { 'status was 204': (r) => r.status == 204 });
  sleep(1);
}