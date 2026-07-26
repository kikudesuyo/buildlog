import type { PageLoad } from './$types';
import { fetchTech } from '$lib/api/client';

export const load: PageLoad = async ({ params, fetch }) => ({
	tech: await fetchTech(fetch, parseInt(params.id, 10))
});
