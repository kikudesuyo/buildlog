import type { PageLoad } from './$types';
import { fetchAnalytics } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => {
	const analyticsData = await fetchAnalytics(fetch);
	return {
		analyticsData
	};
};
