import type { PageLoad } from './$types';
import { fetchDiary } from '$lib/api/client';

export const load: PageLoad = async ({ params, fetch }) => {
	const id = parseInt(params.id, 10);
	const diary = await fetchDiary(fetch, id);
	return {
		diary
	};
};
