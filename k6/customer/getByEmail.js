import http from 'k6/http';
import * as util from '../util/util.js';
import * as fixture from '../util/fixture.js';

export const options = {
	vus: 10,
	duration: '15m',
};

export default function() {
	http.get(`${util.customerBaseUrl}/v1/customers/email/${util.randomItemFromArray(fixture.customerEmails)}`)
	http.get(`${util.customerBaseUrl}/v2/customers/email/${util.randomItemFromArray(fixture.customerEmails)}`)
}
