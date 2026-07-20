import { env } from '$env/dynamic/public';
import type { LoadEvent } from '@sveltejs/kit';
import type {
	AppProject,
	DiaryEntry,
	FeaturedTechArticle,
	ProfileData,
	TechArticle
} from '$lib/api/types';

type ApiFetch = LoadEvent['fetch'];

type ApiListResponse<T> = {
	data_list: T[];
};

type ApiObjectResponse<T> = {
	data: T;
};

type ApiDiaryEntry = Omit<DiaryEntry, 'imageAlt' | 'categoryColorClass'> & {
	image_alt?: string;
	category_color_class?: string;
};

type ApiTechFeed = {
	featured_article: ApiFeaturedTechArticle;
	articles: ApiTechArticle[];
};

type ApiTechArticle = Omit<TechArticle, 'readTime' | 'isNewsletter'> & {
	read_time: string;
	is_newsletter?: boolean;
};

type ApiFeaturedTechArticle = Omit<FeaturedTechArticle, 'readTime'> & {
	read_time: string;
};

type ApiAppProject = Omit<AppProject, 'demoUrl' | 'codeUrl'> & {
	demo_url?: string;
	code_url?: string;
};

type ApiProfileData = Omit<ProfileData, 'avatarUrl' | 'contactEmail' | 'finalQuote'> & {
	avatar_url: string;
	contact_email: string;
	final_quote: string;
};

const apiBaseUrl = env.PUBLIC_API_BASE_URL || 'http://localhost:8081/api/v1';

async function get<T>(fetchFn: ApiFetch, path: string): Promise<T> {
	const response = await fetchFn(`${apiBaseUrl}${path}`);
	if (!response.ok) {
		throw new Error(`API request failed: ${response.status} ${response.statusText}`);
	}
	return response.json() as Promise<T>;
}

export async function fetchDiaryEntries(fetchFn: ApiFetch): Promise<DiaryEntry[]> {
	const response = await get<ApiListResponse<ApiDiaryEntry>>(fetchFn, '/diary');
	return response.data_list.map(({ image_alt, category_color_class, ...entry }) => ({
		...entry,
		imageAlt: image_alt,
		categoryColorClass: category_color_class
	}));
}

export async function fetchTechFeed(fetchFn: ApiFetch): Promise<{
	featuredArticle: FeaturedTechArticle;
	techArticles: TechArticle[];
}> {
	const response = await get<ApiObjectResponse<ApiTechFeed>>(fetchFn, '/tech');
	return {
		featuredArticle: toFeaturedTechArticle(response.data.featured_article),
		techArticles: response.data.articles.map(toTechArticle)
	};
}

export async function fetchAppProjects(fetchFn: ApiFetch): Promise<AppProject[]> {
	const response = await get<ApiListResponse<ApiAppProject>>(fetchFn, '/apps');
	return response.data_list.map(({ demo_url, code_url, ...project }) => ({
		...project,
		demoUrl: demo_url,
		codeUrl: code_url
	}));
}

export async function fetchProfileData(fetchFn: ApiFetch): Promise<ProfileData> {
	const response = await get<ApiObjectResponse<ApiProfileData>>(fetchFn, '/profile');
	const { avatar_url, contact_email, final_quote, ...profile } = response.data;
	return {
		...profile,
		avatarUrl: avatar_url,
		contactEmail: contact_email,
		finalQuote: final_quote
	};
}

function toTechArticle({
	read_time,
	is_newsletter,
	...article
}: ApiTechArticle): TechArticle {
	return {
		...article,
		readTime: read_time,
		isNewsletter: is_newsletter
	};
}

function toFeaturedTechArticle({ read_time, ...article }: ApiFeaturedTechArticle): FeaturedTechArticle {
	return {
		...article,
		readTime: read_time
	};
}
