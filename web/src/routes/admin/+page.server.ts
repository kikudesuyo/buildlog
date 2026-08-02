import type { PageServerLoad } from './$types';
import { fetchCurrentGoals, fetchDiaryEntries } from '$lib/api/client';

const ADMIN_CACHE_CONTROL = 'private, max-age=30, stale-while-revalidate=60';

export const load: PageServerLoad = async ({ fetch, setHeaders }) => {
	setHeaders({ 'cache-control': ADMIN_CACHE_CONTROL });
	const [diaryEntries, goals] = await Promise.all([
		fetchDiaryEntries(fetch, true),
		fetchCurrentGoals(fetch)
	]);

	return { diaryEntries, goals };
};
