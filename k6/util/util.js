export const customerBaseUrl = "http://127.0.0.1:8001"
export const productBaseUrl = "http://127.0.0.1:8002"
export const orderBaseUrl = "http://127.0.0.1:8003"

export function randomString(length, charset = '') {
    if (!charset) charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    let res = '';
    while (length--) res += charset[(Math.random() * charset.length) | 0];
    return res;
}

export function randomItemFromArray(array) {
    return array[Math.floor(Math.random() * array.length)];
}

export function randomDate() {
    const year = Math.floor(Math.random() * 26) + 1980;
    const month = String(Math.floor(Math.random() * 12) + 1).padStart(2, '0');
    const day = String(Math.floor(Math.random() * 28) + 1).padStart(2, '0');
    return `${year}-${month}-${day}`;
}

export function randomPrice() {
    const min = 1.00;
    const max = 1000000.00;
    const decimals = 2;
    const price = (Math.random() * (max - min) + min).toFixed(decimals);
    return parseFloat(price);
}

export function randomInteger(min, max) {
    return Math.floor(Math.random() * (max - min + 1) + min);
}
