import http from 'k6/http';

export const options = {
	vus: 10,
	duration: '5m',
};

export default function() {
	http.get(`http://127.0.0.1:8001/v1/customers/${randomIdFromArray()}`)
	http.get(`http://127.0.0.1:8001/v2/customers/${randomIdFromArray()}`)
}

function randomIdFromArray() {
	const ids = [
		"01HZ60SQE7DC7JQ1Z6EPXVVXDW",
		"01HZ60SQE3CDM3PFDGNG882W6N",
		"01HZ60SQDZ838AXF5PM242QBGM",
		"01HZ60SQDNN97KR9ZVFBR1AC06",
		"01HZ60SQDFEQP7B0PC521718QG",
	];
	return ids[Math.floor(Math.random() * ids.length)]
}
