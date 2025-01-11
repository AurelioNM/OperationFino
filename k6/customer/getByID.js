import http from 'k6/http';
import * as util from '../util/util.js';
import * as fixture from '../util/fixture.js';

export const options = {
	vus: 10,
	duration: '10m',
};

export default function() {
	http.get(`${util.customerBaseUrl}/v1/customers/${util.randomItemFromArray(fixture.customerIds)}`)
	http.get(`${util.customerBaseUrl}/v2/customers/${util.randomItemFromArray(fixture.customerIds)}`)
}
