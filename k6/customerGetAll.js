import http from 'k6/http';
import { check } from 'k6';

export const options = {
	stages: [
		{ duration: '10s', target: 100},
		{ duration: '15s', target: 200},
		{ duration: '30s', target: 300},
		{ duration: '40s', target: 600},
	]
}

export default function() {
	const res = http.get("http://127.0.0.1:8001/customers")
	check(res, {
		'status 200': (r) => r.status === 200
	})
}
