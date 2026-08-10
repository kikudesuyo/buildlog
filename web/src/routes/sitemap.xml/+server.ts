import type { RequestHandler } from './$types';
import { fetchDiaryEntries } from '$lib/api/client';

export const GET: RequestHandler = async ({ fetch }) => {
	const domain = 'https://buildlog.dev'; 

	let diaries: { id: number; updatedAt: string }[] = [];
	try {
		const diaryList = await fetchDiaryEntries(fetch);
		diaries = diaryList.map(diary => ({ id: diary.id, updatedAt: diary.updatedAt }));
	} catch (e) {
		console.error('Failed to fetch diaries for sitemap:', e);
	}

	const sitemap = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	<url>
		<loc>${domain}/</loc>
		<changefreq>daily</changefreq>
		<priority>1.0</priority>
	</url>
	<url>
		<loc>${domain}/profile</loc>
		<changefreq>monthly</changefreq>
		<priority>0.8</priority>
	</url>
	<url>
		<loc>${domain}/contact</loc>
		<changefreq>monthly</changefreq>
		<priority>0.7</priority>
	</url>
	<url>
		<loc>${domain}/apps</loc>
		<changefreq>weekly</changefreq>
		<priority>0.8</priority>
	</url>
	<url>
		<loc>${domain}/tech</loc>
		<changefreq>daily</changefreq>
		<priority>0.9</priority>
	</url>
	<url>
		<loc>${domain}/diary</loc>
		<changefreq>daily</changefreq>
		<priority>0.9</priority>
	</url>
	${diaries.map(diary => `
	<url>
		<loc>${domain}/diary/${diary.id}</loc>
		<lastmod>${diary.updatedAt ? new Date(diary.updatedAt).toISOString().split('T')[0] : new Date().toISOString().split('T')[0]}</lastmod>
		<changefreq>weekly</changefreq>
		<priority>0.7</priority>
	</url>`).join('')}
</urlset>`;

	return new Response(sitemap.trim(), {
		headers: {
			'Content-Type': 'application/xml',
			'Cache-Control': 'max-age=0, s-maxage=3600'
		}
	});
};
