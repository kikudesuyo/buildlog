import type { PageLoad } from './$types';
import { fetchTechFeed } from '$lib/api/client';

export const load: PageLoad = async ({ fetch }) => {
		try {
			return await fetchTechFeed(fetch);
		} catch {
			return { featuredArticle: null, techArticles: [], loadError: true };
		}
};
