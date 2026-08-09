import type { PageLoad } from './$types';
import { fetchCurrentLearnings } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => ({
	daily: await fetchCurrentLearnings(fetch, 'daily'),
	weekly: await fetchCurrentLearnings(fetch, 'weekly'),
	monthly: await fetchCurrentLearnings(fetch, 'monthly')
});
