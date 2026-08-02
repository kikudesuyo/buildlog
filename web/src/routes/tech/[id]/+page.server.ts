import type { PageServerLoad } from './$types';
import { fetchTech } from '$lib/api/client';

export const load: PageServerLoad = async ({ params, fetch, setHeaders }) => {
	setHeaders({ 'cache-control': 'public, max-age=30, s-maxage=30, stale-while-revalidate=60' });
	const id = parseInt(params.id, 10);
	const tech = await fetchTech(fetch, id);
	return { tech: { ...tech, hasLiked: false } };
};
