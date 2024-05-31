import http from 'k6/http';

export const options = {
	vus: 10,
	duration: '15m',
};

export default function() {
	http.get(`http://127.0.0.1:8001/v1/customers/email/${randomEmailFromArray()}`)
	http.get(`http://127.0.0.1:8001/v2/customers/email/${randomEmailFromArray()}`)
}

function randomEmailFromArray() {
	const emails = [
		"f0zfahOE@gmail.com",
		"3PUJWCnE@gmail.com",
		"gGvt9EH8@gmail.com",
		"tu1QnS1J@gmail.com",
		"0RmGUKg0@gmail.com",
		"d1lXhURb@gmail.com",
		"H4WPQ73M@gmail.com",
		"doGQawAZ@gmail.com",
		"5Z1jCw1z@gmail.com",
		"AABNvVk0@gmail.com",
		"lbxnMBGi@gmail.com",
		"WnSqTV2G@gmail.com",
		"X8kueWbo@gmail.com",
		"nQcDQWIz@gmail.com",
		"MC60mEA2@gmail.com",
		"YK0Bl5ly@gmail.com",
		"jutE4Ek4@gmail.com",
		"gaCq1e8T@gmail.com",
		"1mHY6Kto@gmail.com",
		"orQ6vTy0@gmail.com",
	];
	return emails[Math.floor(Math.random() * emails.length)]
}
