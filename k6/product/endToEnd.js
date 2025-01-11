import http from 'k6/http';
import { check } from 'k6';
import * as util from '../util/util.js';

export const options = {
	vus: 1,
	duration: '30m',
};

export default function() {
	let url = `${util.productBaseUrl}/v1/products`

	// POST
	const createResponse = http.post(url, generateJson())
	check(createResponse, {
		'create status 201': (r) => r.status === 201
	})

	const createdId = createResponse.json().data.id
	url = `${url}/${createdId}`

	// GET
	const getResponse = http.get(url)
	check(getResponse, {
		'get status 200': (r) => r.status === 200
	})

	// PUT
	const updateResponse = http.put(url, generateJson())
	check(updateResponse, {
		'update status 200': (r) => r.status === 200
	})

	// DELETE
	const deleteResponse = http.del(url)
	check(deleteResponse, {
		'delete status 200': (r) => r.status === 200
	})
}

function generateJson() {
	return JSON.stringify({
		name: util.randomString(6),
		description: util.randomString(30),
		price: util.randomPrice(),
		quantity: util.randomInteger(1, 100)
	})
}
