import type { PageServerLoad } from './$types';
import { fetchDiaryEntries } from '$lib/api/client';

export const load: PageServerLoad = async ({ fetch, setHeaders, url }) => {
	setHeaders({ 'cache-control': 'public, max-age=30, s-maxage=30, stale-while-revalidate=60' });
	const sort = url.searchParams.get('sort') === 'likes' ? 'likes' : 'newest';
	const order = url.searchParams.get('order') === 'asc' ? 'asc' : 'desc';
	const diaryEntries = await fetchDiaryEntries(fetch, false, 0, 0, sort, order);
	return {
		diaryEntries: diaryEntries.map((entry) => ({ ...entry, hasLiked: false }))
	};
};
