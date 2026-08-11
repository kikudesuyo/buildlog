import type { PageLoad } from './$types';
import { fetchTechFeed } from '$lib/api/client';

export const load: PageLoad = async ({ fetch, url }) => {
	const order = url.searchParams.get('order') === 'asc' ? 'asc' : 'desc';
	return fetchTechFeed(fetch, false, 0, 0, order);
};
