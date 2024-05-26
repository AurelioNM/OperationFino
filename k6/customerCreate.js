import http from 'k6/http';
import { check } from 'k6';

export const options = {
	vus: 20,
	duration: '200s',
};

export default function() {
	const url = "http://127.0.0.1:8001/v1/customers"

	const payload = JSON.stringify({
		name: randomString(6),
		surname: randomString(6),
		email: `${randomString(8)}@gmail.com`,
		birthdate: "1998-08-18"
	})

	const res = http.post(url, payload)

	check(res, {
		'status 201': (r) => r.status === 201
	})
}

function randomString(length, charset = '') {
	if (!charset) charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
	let res = ''
	while (length --) res += charset[(Math.random() * charset.length) | 0]
	return res
}

function randomWordFromArray() {
	const words = ['PS5', 'Shirt', 'Notebook', 'Coffee', 'JBL-5']
	return words[Math.floor(Math.random() * words.length)]
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

function randomDate() {
	const year = Math.floor(Math.random() * 26) + 2000
	const month = String(Math.floor(Math.random() * 12) + 1).padStart(2, '0')
	const day = String(Math.floor(Math.random() * 31) + 1).padStart(2, '0')
	return `${year}-${month}-${day}`
}

