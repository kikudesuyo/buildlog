import type { PageServerLoad } from './$types';
import { fetchDiaryEntries } from '$lib/api/client';

export const load: PageServerLoad = async ({ fetch, setHeaders }) => {
	setHeaders({ 'cache-control': 'public, max-age=30, s-maxage=30, stale-while-revalidate=60' });
	const diaryEntries = await fetchDiaryEntries(fetch);
	return {
		diaryEntries: diaryEntries.map((entry) => ({ ...entry, hasLiked: false }))
	};
};
