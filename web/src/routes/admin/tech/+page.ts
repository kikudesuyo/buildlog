import type { PageLoad } from './$types';
import { fetchTechFeed, type TechSort } from '$lib/api/client';

export const load: PageLoad = async ({ fetch, url }) => {
	const order = url.searchParams.get('order') === 'asc' ? 'asc' : 'desc';
	const sort: TechSort = url.searchParams.get('sort') === 'likes' ? 'likes' : 'newest';
	return fetchTechFeed(fetch, false, 0, 0, order, sort);
};
