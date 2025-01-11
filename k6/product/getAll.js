import http from 'k6/http';
import { check } from 'k6';
import * as util from '../util/util.js';

export const options = {
	vus: 5,
	duration: '1m',
};

export default function() {
	const url = `${util.productBaseUrl}/v1/products`

	const res = http.get(url)
	check(res, {
		'status 200': (r) => r.status === 200
	})
}
