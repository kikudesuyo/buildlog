import type { PageLoad } from './$types';
import { fetchCurrentGoals, fetchDiaryEntries } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => {
	const [diaryEntries, goals] = await Promise.all([
		fetchDiaryEntries(fetch, true),
		fetchCurrentGoals(fetch)
	]);
	return { diaryEntries, goals };
};
