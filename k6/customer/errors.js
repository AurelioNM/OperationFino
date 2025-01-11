import http from 'k6/http';
import { check } from 'k6';
import * as util from '../util/util.js';
import * as fixture from '../util/fixture.js';

export const options = {
	stages: [
		{ duration: '2m', target: 1},
		{ duration: '3m', target: 2},
		{ duration: '4m', target: 3},
		{ duration: '6m', target: 4},
		{ duration: '7m', target: 5},
	]
}

export default function() {
	let url = `${util.customerBaseUrl}/v1/customers`
	const params = {
		headers: {
			'Content-Type': 'application/json',
		}
	}

	// POST
	const createResponse = http.post(url, generateJson(), params)
	check(createResponse, {
		'create status 201': (r) => r.status === 201
	})
	const createdId = "errorID"
	url = `${url}/${createdId}`

	// GET
	const getResponse = http.get(url)
	check(getResponse, {
		'get status 200': (r) => r.status === 200
	})

	// PUT
	const updateResponse = http.put(
		url, 
		generateJson(),
		params
	)
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
		surname: util.randomString(6),
		email: `errorEmail@gmail.com`,
		birthdate: "1998-08-18"
	})
}

