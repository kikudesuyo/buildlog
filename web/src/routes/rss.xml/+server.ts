import type { RequestHandler } from './$types';
import { fetchDiaryEntries, fetchTechFeed } from '$lib/api/client';

export const GET: RequestHandler = async ({ fetch }) => {
	const siteUrl = 'http://localhost:5173';

	// 並行して日記と技術記事を取得
	const [diaries, techFeed] = await Promise.all([
		fetchDiaryEntries(fetch),
		fetchTechFeed(fetch)
	]);

	// すべての記事を統合
	const articles = [
		...diaries.map(d => ({
			title: d.title,
			content: d.content,
			createdAt: d.createdAt,
			link: `/diary/${d.id}`
		})),
		...techFeed.techArticles.map(t => ({
			title: t.title,
			content: t.content,
			createdAt: t.createdAt,
			link: t.external?.url ?? `${siteUrl}/tech`
		}))
	];

	// 日付順にソート (最新が上)
	articles.sort((a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime());

	// 最新の 20 件に絞り込み
	const latestArticles = articles.slice(0, 20);

	const rssItems = latestArticles.map(item => `
		<item>
			<title>${escapeXml(item.title)}</title>
			<link>${item.link}</link>
			<guid>${item.link}</guid>
			<pubDate>${new Date(item.createdAt).toUTCString()}</pubDate>
			<description>${escapeXml(item.content)}</description>
		</item>
	`).join('');

	const xml = `<?xml version="1.0" encoding="UTF-8" ?>
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
<channel>
	<title>Buildlog</title>
	<link>${siteUrl}</link>
	<description>思考の断片を、構造化された知性へ。技術録と日々の思考の記録。</description>
	<language>ja</language>
	<atom:link href="${siteUrl}/rss.xml" rel="self" type="application/rss+xml" />
	${rssItems}
</channel>
</rss>`.trim();

	return new Response(xml, {
		headers: {
			'Cache-Control': 'max-age=0, s-maxage=3600',
			'Content-Type': 'application/xml; charset=utf-8'
		}
	});
};

function escapeXml(unsafe: string): string {
	return unsafe.replace(/[<>&'"]/g, (c) => {
		switch (c) {
			case '<': return '&lt;';
			case '>': return '&gt;';
			case '&': return '&amp;';
			case '\'': return '&apos;';
			case '"': return '&quot;';
			default: return c;
		}
	});
}
