const SESSION_MAX_AGE_MS = 1000 * 60 * 60 * 24 * 30;
const encoder = new TextEncoder();

function toBase64Url(bytes: Uint8Array): string {
	let binary = '';
	for (const byte of bytes) binary += String.fromCharCode(byte);
	return btoa(binary).replaceAll('+', '-').replaceAll('/', '_').replace(/=+$/, '');
}

async function signTimestamp(timestamp: string, secret: string): Promise<string> {
	const key = await crypto.subtle.importKey(
		'raw',
		encoder.encode(secret),
		{ name: 'HMAC', hash: 'SHA-256' },
		false,
		['sign', 'verify']
	);
	const signature = await crypto.subtle.sign('HMAC', key, encoder.encode(timestamp));
	return toBase64Url(new Uint8Array(signature));
}

export async function createAdminSession(secret: string): Promise<string> {
	const timestamp = String(Date.now());
	return `${timestamp}.${await signTimestamp(timestamp, secret)}`;
}

export async function isValidAdminSession(value: string | undefined, secret: string): Promise<boolean> {
	if (!value) return false;
	const [timestamp, signature] = value.split('.');
	if (!timestamp || !signature || !/^\d+$/.test(timestamp)) return false;
	const issuedAt = Number(timestamp);
	if (!Number.isSafeInteger(issuedAt) || Date.now() - issuedAt < 0 || Date.now() - issuedAt > SESSION_MAX_AGE_MS) return false;
	return signature === (await signTimestamp(timestamp, secret));
}
