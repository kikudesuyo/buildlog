import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ url }) => {
	const targetUrl = url.searchParams.get('url');
	if (!targetUrl) {
		return json({ error: 'Missing url parameter' }, { status: 400 });
	}

	try {
		// 外部サーバーへリクエスト（タイムアウトを5秒に制限）
		const controller = new AbortController();
		const timeoutId = setTimeout(() => controller.abort(), 5000);

		const response = await fetch(targetUrl, {
			signal: controller.signal,
			headers: {
				'User-Agent': 'Mozilla/5.0 (compatible; BuildlogLinkBot/1.0; +http://localhost:5173)'
			}
		});
		clearTimeout(timeoutId);

		if (!response.ok) {
			return json({
				title: targetUrl,
				description: `Could not fetch preview (Status ${response.status})`,
				image: '',
				siteName: new URL(targetUrl).hostname
			});
		}

		const html = await response.text();

		// OGP情報の抽出
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
		]) || new URL(targetUrl).hostname;

		return json({
			title,
			description,
			image,
			siteName
		});

	} catch {
		return json({
			title: targetUrl,
			description: 'Link preview currently unavailable.',
			image: '',
			siteName: new URL(targetUrl).hostname
		});
	}
};

function extractMeta(html: string, regexList: RegExp[]): string {
	for (const regex of regexList) {
		const match = html.match(regex);
		if (match && match[1]) {
			return match[1]
				.replace(/&amp;/g, '&')
				.replace(/&lt;/g, '<')
				.replace(/&gt;/g, '>')
				.replace(/&quot;/g, '"')
				.replace(/&#39;/g, '\'')
				.trim();
		}
	}
	return '';
}
