import http from 'k6/http';
import * as util from '../util/util.js';
import * as fixture from '../util/fixture.js';

export const options = {
	vus: 5,
	duration: '10m',
};

export default function() {
	http.get(`${util.productBaseUrl}/v1/products/name/${util.randomItemFromArray(fixture.productNames)}`)
}
