import http from 'k6/http';
import { check } from 'k6';

export const options = {
	vus: 1,
	duration: '30m',
};

export default function() {
	let url = "http://127.0.0.1:8002/v1/products"

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

function generateRandomString(length, charset = '') {
	if (!charset) charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
	let str = ''
	while (length --) str += charset[(Math.random() * charset.length) | 0]
	return str
}

function generateJson() {
	return JSON.stringify({
		name: generateRandomString(6),
		description: generateRandomString(30),
		price: randomPrice(),
		quantity: randomInteger(1, 100)
	})
}

function randomPrice() {
	const min = 1.00
	const max = 1000000.00
	const decimals = 2
	const price = (Math.random() * (max - min) + min).toFixed(decimals)
	return parseFloat(price)
}

function randomInteger(min, max) {
	return Math.floor(Math.random() * (max - min + 1) + min)
}

