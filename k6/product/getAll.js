import http from 'k6/http';
import { check } from 'k6';

export const options = {
	vus: 1,
	duration: '10s',
	iterations: 1,
};

export default function() {
	const res = http.get("http://127.0.0.1:8002/v1/products")
	check(res, {
		'status 200': (r) => r.status === 200
	})
}
