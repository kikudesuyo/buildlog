import type { PageServerLoad } from './$types';
import { fetchTechFeed, type TechSort } from '$lib/api/client';

export const load: PageServerLoad = async ({ fetch, setHeaders, url }) => {
	setHeaders({ 'cache-control': 'public, max-age=30, s-maxage=30, stale-while-revalidate=60' });
	const order = url.searchParams.get('order') === 'asc' ? 'asc' : 'desc';
	const sort: TechSort = url.searchParams.get('sort') === 'likes' ? 'likes' : 'newest';
	try {
		const data = await fetchTechFeed(fetch, false, 0, 4, order, sort);
		return {
			...data,
			techArticles: data.techArticles.map((article) => ({ ...article, hasLiked: false })),
			hasMore: data.hasMore
		};
	} catch {
		setHeaders({ 'cache-control': 'no-store' });
		return { techArticles: [], hasMore: false, loadError: true };
	}
};
