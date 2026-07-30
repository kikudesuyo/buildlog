import type { PageLoad } from './$types';
import { fetchProfile } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => {
	const profileData = await fetchProfile(fetch);
	return {
		profileData
	};
};
