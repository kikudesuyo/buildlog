import { fetchAdminComments } from '$lib/api/client';

export const load = async () => ({ comments: await fetchAdminComments() });
