import http from 'k6/http';

export const options = {
	vus: 1,
	duration: '5m',
};

export default function() {
	http.get(`http://127.0.0.1:8001/v1/customers/${randomIdFromArray()}`)
	http.get(`http://127.0.0.1:8001/v2/customers/${randomIdFromArray()}`)
}

function randomIdFromArray() {
	const ids = [
		"01HZ77Y96590PQ3Q9G0AHHTRDZ",
		"01HZ77Y96590PQ3Q9G0R0NECHR",
		"01HZ77Y966KZ9PD7RP28ESEHNR",
		"01HZ77Y968BVPT7TEP8Z8CCF3H",
		"01HZ77Y9692H3AX90BD0VJVJXD",
		"01HZ77Y96BW26F83DWGT53J6VC",
		"01HZ77Y96CJ3FB7WVVJ4J1KW0E",
		"01HZ77Y96DGE5S6GK5N3V4A45N",
		"01HZ77Y96GEXA1A2J7XBS3DDZ0",
		"01HZ77Y96H3QHZY93HSN63G36N",
		"01HZ77Y96KW2BRWPFN5XQ3QDB9",
		"01HZ77Y96M5QJ1FT5ZNJB9416V",
		"01HZ77Y96VFQ2JNSJ0KH1CXEN7",
		"01HZ77Y96YJB75JKHR0C9BF6Z5",
		"01HZ77Y971M38V7HA6KN8DY1F8",
		"01HZ77Y976T4RJEGBW1M01ZD8Z",
		"01HZ77Y97GJGP2SFZ3CVJN1XAE",
	];
	return ids[Math.floor(Math.random() * ids.length)]
}
