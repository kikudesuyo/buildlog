import type { PageLoad } from './$types';
import { fetchProfile, fetchPostHistory } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => {
	const profileData = await fetchProfile(fetch);
	const postHistory = await fetchPostHistory(fetch);
	return {
		profileData,
		postHistory
	};
};
