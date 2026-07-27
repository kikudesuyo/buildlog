import type { PageLoad } from './$types';
import { fetchTechFeed } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => fetchTechFeed(fetch, true);
