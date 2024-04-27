import http from 'k6/http';
import { check } from 'k6';

export const options = {
	vus: 1,
	duration: '1s',
};

function randomString(length, charset = '') {
	if (!charset) charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
	let res = ''
	while (length --) res += charset[(Math.random() * charset.length) | 0]
	return res
}

export default function() {
	const url = "http://127.0.0.1:8001/v1/customers"
	const params = {
		headers: {
			'Content-Type': 'application/json',
		}
	}

	const payload = JSON.stringify({
		name: randomString(6),
		surname: randomString(6),
		email: `${randomString(8)}@gmail.com`,
		birthdate: "1998-08-18"
	})

	const res = http.post(url, payload, params)

	check(res, {
		'status 201': (r) => r.status === 201
	})
}
