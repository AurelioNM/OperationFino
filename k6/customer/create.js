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
	const url = `${util.customerBaseUrl}/v1/customers`

	const payload = JSON.stringify({
		name: util.randomItemFromArray(fixture.customerNames),
		surname: util.randomItemFromArray(fixture.customerNames),
		email: `${util.randomString(8)}@gmail.com`,
		birthdate: util.randomDate()
	})

	const res = http.post(url, payload)
	check(res, {
		'status 201': (r) => r.status === 201
	})
}
