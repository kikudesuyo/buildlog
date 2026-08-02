import type { PageServerLoad } from './$types';
import { fetchTechFeed } from '$lib/api/client';

export const load: PageServerLoad = async ({ fetch, setHeaders }) => {
	setHeaders({ 'cache-control': 'public, max-age=30, s-maxage=30, stale-while-revalidate=60' });
	try {
		const data = await fetchTechFeed(fetch);
		return {
			...data,
			featuredArticle: data.featuredArticle ? { ...data.featuredArticle, hasLiked: false } : null,
			techArticles: data.techArticles.map((article) => ({ ...article, hasLiked: false }))
		};
	} catch {
		setHeaders({ 'cache-control': 'no-store' });
		return { featuredArticle: null, techArticles: [], loadError: true };
	}
};
