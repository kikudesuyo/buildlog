import type { PageServerLoad } from './$types';
import { fetchCurrentGoals, fetchDiaryEntries } from '$lib/api/client';

const ADMIN_CACHE_CONTROL = 'private, max-age=30, stale-while-revalidate=60';

export const load: PageServerLoad = async ({ fetch, setHeaders, url }) => {
	setHeaders({ 'cache-control': ADMIN_CACHE_CONTROL });
	const sort = url.searchParams.get('sort') === 'likes' ? 'likes' : 'newest';
	const order = url.searchParams.get('order') === 'asc' ? 'asc' : 'desc';
	const [diaryEntries, goals] = await Promise.all([
		fetchDiaryEntries(fetch, true, 0, 0, sort, order),
		fetchCurrentGoals(fetch)
	]);

	return { diaryEntries, goals };
};
