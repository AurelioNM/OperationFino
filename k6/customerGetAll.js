import http from 'k6/http';
import { check } from 'k6';

export const options = {
	stages: [
		{ duration: '3s', target: 100},
		{ duration: '5s', target: 200},
		{ duration: '7s', target: 300},
	]
}

export default function() {
	const res = http.get("http://127.0.0.1:8001/customers")
	check(res, {
		'status 200': (r) => r.status === 200
	})
}
