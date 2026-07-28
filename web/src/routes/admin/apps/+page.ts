import type { PageLoad } from './$types';
import { fetchAppProjects } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => {
	const appProjects = await fetchAppProjects(fetch);
	return {
		appProjects
	};
};
