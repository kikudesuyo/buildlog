import type { PageLoad } from './$types';
import { fetchProfileData } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => ({
	profileData: await fetchProfileData(fetch)
});
