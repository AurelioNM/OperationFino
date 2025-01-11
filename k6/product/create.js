import http from 'k6/http';
import { check } from 'k6';
import * as util from '../util/util.js';
import * as fixture from '../util/fixture.js';

export const options = {
	vus: 20,
	duration: '10m',
	// iterations: 3,
};

export default function() {
	const url = `${util.productBaseUrl}/v1/products`

	const payload = JSON.stringify({
		name: util.randomString(10),
		description: util.randomString(30),
		price: util.randomPrice(),
		quantity: util.randomInteger(1, 100)
	})

	const res = http.post(url, payload)
	check(res, {
		'status 201': (r) => r.status === 201
	})
}
