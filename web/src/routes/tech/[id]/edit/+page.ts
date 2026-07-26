import type { PageLoad } from './$types';
import { fetchTech } from '$lib/api/client';

export const load: PageLoad = async ({ params, fetch }) => {
	const id = parseInt(params.id, 10);
	const tech = await fetchTech(fetch, id);
	return {
		tech
	};
};
