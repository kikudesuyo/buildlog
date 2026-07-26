import type { PageLoad } from './$types';
import { fetchDiaryEntries } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => ({
	diaryEntries: await fetchDiaryEntries(fetch)
});
