import http from 'k6/http'
import { check } from 'k6'

export const options = {
    vus: 50,
    duration: '30s',
};

export default function(){
    const res = http.get('http://127.0.0.1:3000/dhPX2', {redirects: 0});
    check(res, {
        'status is 302': (r) => r.status === 302,
    });}
