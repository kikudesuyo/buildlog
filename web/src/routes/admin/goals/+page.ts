import type { PageLoad } from './$types';
import { fetchGoalHistory } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => ({
	goalHistory: await fetchGoalHistory(fetch)
});
