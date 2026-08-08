import type { PageLoad } from './$types';
import { fetchAdminComments } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => ({ comments: await fetchAdminComments(fetch) });
