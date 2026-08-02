import type { PageLoad } from './$types';
import { fetchTechFeed } from '$lib/api/client';

export const load: PageLoad = async ({ fetch, url }) => {
	const tag = url.searchParams.get('tag') || undefined;
		try {
			const feed = await fetchTechFeed(fetch, tag);
			return { ...feed, selectedTag: tag };
		} catch {
			return { featuredArticle: null, techArticles: [], loadError: true, selectedTag: tag };
		}
};
