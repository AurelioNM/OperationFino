import http from 'k6/http';

export const options = {
	vus: 1,
	duration: '10m',
};

export default function() {
	http.get(`http://127.0.0.1:8002/v1/products/${randomIdFromArray()}`)
}

function randomIdFromArray() {
	const ids = [
		"01JGVB6K5PENPFQK6G8WGQPARE",
		"01JGVB7N668AWJ6SE4159NHB8W",
		"01JGVC70XR8AR810W1SAJVW5W3",
		"01JGVJQJCQ4PF9Y7XHH78GCK5D",
	];
	return ids[Math.floor(Math.random() * ids.length)]
}
