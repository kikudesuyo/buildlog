import type { RequestHandler } from './$types';
import { fetchDiaryEntries, fetchTechFeed } from '$lib/api/client';

export const GET: RequestHandler = async ({ fetch }) => {
	const domain = 'https://buildlog.dev'; 

	let techs: { id: number; updatedAt: string }[] = [];
	let diaries: { id: number; updatedAt: string }[] = [];
	try {
		const [{ featuredArticle, techArticles }, diaryList] = await Promise.all([
			fetchTechFeed(fetch),
			fetchDiaryEntries(fetch)
		]);
		const all = [];
		if (featuredArticle && featuredArticle.id !== 0) {
			all.push(featuredArticle);
		}
		if (techArticles && techArticles.length > 0) {
			all.push(...techArticles);
		}
		techs = all.map(t => ({
			id: t.id,
			updatedAt: t.updatedAt
		}));
		diaries = diaryList.map(diary => ({ id: diary.id, updatedAt: diary.updatedAt }));
	} catch (e) {
		console.error('Failed to fetch techs for sitemap:', e);
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
	${techs.map(tech => `
	<url>
		<loc>${domain}/tech/${tech.id}</loc>
		<lastmod>${tech.updatedAt ? new Date(tech.updatedAt).toISOString().split('T')[0] : new Date().toISOString().split('T')[0]}</lastmod>
		<changefreq>weekly</changefreq>
		<priority>0.7</priority>
	</url>`).join('')}
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
