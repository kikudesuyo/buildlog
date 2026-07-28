import type { PageLoad } from './$types';
import { fetchDiaryEntries } from '$lib/api/client';

export const load: PageLoad = async ({ fetch, url }) => {
	const tag = url.searchParams.get('tag') || undefined;
	return {
		diaryEntries: await fetchDiaryEntries(fetch, tag),
		selectedTag: tag
	};
};
