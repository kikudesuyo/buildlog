import type { PageLoad } from './$types';
import { fetchApp } from '$lib/api/client';

export const load: PageLoad = async ({ fetch, params }) => {
	const id = Number(params.id);
	const appProject = await fetchApp(fetch, id);
	return {
		appProject
	};
};
