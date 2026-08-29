import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

const MAX_RESPONSE_BYTES = 512 * 1024;
const FETCH_TIMEOUT_MS = 5000;

export const GET: RequestHandler = async ({ url }) => {
	const targetUrl = url.searchParams.get('url');
	if (!targetUrl) return json({ error: 'Missing url parameter' }, { status: 400 });

	let parsedUrl: URL;
	try {
		parsedUrl = new URL(targetUrl);
		if (parsedUrl.protocol !== 'https:' || parsedUrl.username || parsedUrl.password) {
			return json({ error: 'Only public HTTPS URLs are supported' }, { status: 400 });
		}
		if (isPrivateAddress(parsedUrl.hostname)) {
			return json({ error: 'Private network URLs are not supported' }, { status: 400 });
		}
		await assertPublicHostname(parsedUrl.hostname);
	} catch {
		return json({ error: 'Invalid or private URL' }, { status: 400 });
	}

	try {
		const controller = new AbortController();
		const timeoutId = setTimeout(() => controller.abort(), FETCH_TIMEOUT_MS);
		let response: Response;
		try {
			response = await fetch(parsedUrl, {
				signal: controller.signal,
				redirect: 'error',
				headers: { 'User-Agent': 'Mozilla/5.0 (compatible; BuildlogLinkBot/1.0; +https://buildlog.dev)' }
			});
		} finally {
			clearTimeout(timeoutId);
		}

		if (!response.ok) {
			return json({
				title: targetUrl,
				description: `Could not fetch preview (Status ${response.status})`,
				image: '',
				siteName: parsedUrl.hostname
			});
		}

		const html = await readResponseText(response, MAX_RESPONSE_BYTES);
		const title = extractMeta(html, [
			/<meta[^>]*property=["']og:title["'][^>]*content=["']([^"']+)["']/i,
			/<meta[^>]*content=["']([^"']+)["'][^>]*property=["']og:title["']/i,
			/<title>([^<]+)<\/title>/i
		]) || targetUrl;
		const description = extractMeta(html, [
			/<meta[^>]*property=["']og:description["'][^>]*content=["']([^"']+)["']/i,
			/<meta[^>]*content=["']([^"']+)["'][^>]*property=["']og:description["']/i,
			/<meta[^>]*name=["']description["'][^>]*content=["']([^"']+)["']/i,
			/<meta[^>]*content=["']([^"']+)["'][^>]*name=["']description["']/i
		]);
		const image = extractMeta(html, [
			/<meta[^>]*property=["']og:image["'][^>]*content=["']([^"']+)["']/i,
			/<meta[^>]*content=["']([^"']+)["'][^>]*property=["']og:image["']/i
		]);
		const siteName = extractMeta(html, [
			/<meta[^>]*property=["']og:site_name["'][^>]*content=["']([^"']+)["']/i,
			/<meta[^>]*content=["']([^"']+)["'][^>]*property=["']og:site_name["']/i
		]) || parsedUrl.hostname;

		return json({ title, description, image, siteName });
	} catch {
		return json({ title: targetUrl, description: 'Link preview currently unavailable.', image: '', siteName: parsedUrl.hostname });
	}
};

async function readResponseText(response: Response, maxBytes: number): Promise<string> {
	if (!response.body) return '';
	const reader = response.body.getReader();
	const chunks: Uint8Array[] = [];
	let totalBytes = 0;
	try {
		while (true) {
			const { done, value } = await reader.read();
			if (done) break;
			if (totalBytes + value.byteLength > maxBytes) throw new Error('OGP response is too large');
			totalBytes += value.byteLength;
			chunks.push(value);
		}
	} finally {
		reader.releaseLock();
	}
	const bytes = new Uint8Array(totalBytes);
	let offset = 0;
	for (const chunk of chunks) {
		bytes.set(chunk, offset);
		offset += chunk.byteLength;
	}
	return new TextDecoder().decode(bytes);
}

async function assertPublicHostname(hostname: string): Promise<void> {
	if (isPrivateAddress(hostname)) throw new Error('Private network hostname is not supported');
	const { lookup } = await import('node:dns/promises');
	const addresses = await lookup(hostname, { all: true, verbatim: true });
	if (addresses.length === 0 || addresses.some(({ address }) => isPrivateAddress(address))) {
		throw new Error('Private network hostname is not supported');
	}
}

function isPrivateAddress(hostname: string): boolean {
	const normalized = hostname.replace(/^\[|\]$/g, '').toLowerCase();
	return normalized === 'localhost' || normalized === 'localhost.localdomain' || normalized === 'metadata.google.internal' ||
		/^127\./.test(normalized) || /^10\./.test(normalized) || /^192\.168\./.test(normalized) ||
		/^172\.(1[6-9]|2\d|3[0-1])\./.test(normalized) || /^169\.254\./.test(normalized) ||
		normalized === '::1' || normalized.startsWith('fc') || normalized.startsWith('fd') || normalized.startsWith('fe80:');
}

function extractMeta(html: string, regexList: RegExp[]): string {
	for (const regex of regexList) {
		const match = html.match(regex);
		if (match?.[1]) return match[1].replace(/&amp;/g, '&').replace(/&lt;/g, '<').replace(/&gt;/g, '>').replace(/&quot;/g, '"').replace(/&#39;/g, "'").trim();
	}
	return '';
}
