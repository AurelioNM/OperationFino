import http from 'k6/http';
import { check } from 'k6';

export const options = {
	vus: 3,
	duration: '15m',
};

function generateRandomString(length, charset = '') {
	if (!charset) charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
	let str = ''
	while (length --) str += charset[(Math.random() * charset.length) | 0]
	return str
}

function generateJson() {
	return JSON.stringify({
		name: generateRandomString(6),
		surname: generateRandomString(6),
		email: `${generateRandomString(8)}@gmail.com`,
		birthdate: "1998-08-18"
	})
}

export default function() {
	let url = "http://127.0.0.1:8001/v1/customers"
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
	const createdId = createResponse.json().data.id
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
