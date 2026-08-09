import type { PageLoad } from './$types';
import { fetchProfile } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => ({
	profileData: await fetchProfile(fetch)
});
