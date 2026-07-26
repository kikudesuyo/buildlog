import type { PageLoad } from './$types';
import { fetchDiary } from '$lib/api/client';

export const load: PageLoad = async ({ params, fetch }) => ({
	diary: await fetchDiary(fetch, parseInt(params.id, 10))
});
