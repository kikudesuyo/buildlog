import type { PageLoad } from './$types';
import { fetchTrashEntries } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => {
	const trashEntries = await fetchTrashEntries(fetch);
	return {
		trashEntries
	};
};
