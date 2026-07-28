import * as env from '$env/static/public';
import type { LoadEvent } from '@sveltejs/kit';
import type {
	DiaryEntry,
	FeaturedTechArticle,
	TechArticle,
	TrashEntry,
	AppProject,
	ProfileData
} from '$lib/api/types';

type ApiFetch = LoadEvent['fetch'];

type ApiListResponse<T> = {
	data_list: T[];
};

type ApiObjectResponse<T> = {
	data: T;
};




const apiBaseUrl = (() => {
	const rawUrl = import.meta.env.PUBLIC_API_BASE_URL || 'http://localhost:8081';
	return rawUrl.endsWith('/api/v1') ? rawUrl : `${rawUrl}/api/v1`;
})();

async function get<T>(fetchFn: ApiFetch, path: string): Promise<T> {
	const response = await fetchFn(`${apiBaseUrl}${path}`);
	if (!response.ok) {
		throw new Error(`API request failed: ${response.status} ${response.statusText}`);
	}
	return response.json() as Promise<T>;
}

export type ApiPost = {
	id: number;
	type: string;
	title: string;
	content: string;
	excerpt: string;
	category: string;
	views: string;
	created_at: string;
	updated_at: string;
};

export async function fetchDiaryEntries(fetchFn: ApiFetch): Promise<DiaryEntry[]> {
	const response = await get<ApiListResponse<ApiPost>>(fetchFn, '/diaries');
	return response.data_list.map((post) => ({
		id: post.id,
		title: post.title,
		content: post.content,
		createdAt: post.created_at,
		updatedAt: post.updated_at
	}));
}

export async function fetchTechFeed(fetchFn: ApiFetch): Promise<{
	featuredArticle: FeaturedTechArticle;
	techArticles: TechArticle[];
}> {
	const response = await get<ApiListResponse<ApiPost>>(fetchFn, '/techs');
	
	const allArticles: TechArticle[] = response.data_list.map((post) => ({
		id: post.id,
		title: post.title,
		content: post.content,
		category: post.category,
		views: post.views,
		createdAt: post.created_at,
		updatedAt: post.updated_at
	}));

	const featured = allArticles.length > 0 ? allArticles[0] : {
		id: 0,
		title: '',
		content: '',
		category: '',
		views: '',
		createdAt: '',
		updatedAt: ''
	};
	const remaining = allArticles.length > 1 ? allArticles.slice(1) : [];

	return {
		featuredArticle: featured,
		techArticles: remaining
	};
}

async function sendRequest<T>(method: string, path: string, body?: unknown): Promise<T> {
	const response = await fetch(`${apiBaseUrl}${path}`, {
		method,
		headers: body ? { 'Content-Type': 'application/json' } : undefined,
		body: body ? JSON.stringify(body) : undefined
	});
	if (!response.ok) {
		throw new Error(`API request failed: ${response.status} ${response.statusText}`);
	}
	return response.json() as Promise<T>;
}

export async function createDiary(title: string, content: string): Promise<DiaryEntry> {
	const response = await sendRequest<ApiObjectResponse<ApiPost>>('POST', '/diaries', { title, content });
	return {
		id: response.data.id,
		title: response.data.title,
		content: response.data.content,
		createdAt: response.data.created_at,
		updatedAt: response.data.updated_at
	};
}

export async function updateDiary(id: number, title: string, content: string): Promise<DiaryEntry> {
	const response = await sendRequest<ApiObjectResponse<ApiPost>>('PUT', `/diaries/${id}`, { title, content });
	return {
		id: response.data.id,
		title: response.data.title,
		content: response.data.content,
		createdAt: response.data.created_at,
		updatedAt: response.data.updated_at
	};
}

export async function deleteDiary(id: number): Promise<void> {
	await sendRequest<void>('DELETE', `/diaries/${id}`);
}

export async function createTech(req: {
	title: string;
	content: string;
	category: string;
	views?: string;
}): Promise<TechArticle> {
	const response = await sendRequest<ApiObjectResponse<ApiPost>>('POST', '/techs', {
		title: req.title,
		content: req.content,
		category: req.category,
		views: req.views || ''
	});
	return {
		id: response.data.id,
		title: response.data.title,
		content: response.data.content,
		category: response.data.category,
		views: response.data.views,
		createdAt: response.data.created_at,
		updatedAt: response.data.updated_at
	};
}

export async function updateTech(id: number, req: {
	title: string;
	content: string;
	category: string;
	views?: string;
}): Promise<TechArticle> {
	const response = await sendRequest<ApiObjectResponse<ApiPost>>('PUT', `/techs/${id}`, {
		title: req.title,
		content: req.content,
		category: req.category,
		views: req.views || ''
	});
	return {
		id: response.data.id,
		title: response.data.title,
		content: response.data.content,
		category: response.data.category,
		views: response.data.views,
		createdAt: response.data.created_at,
		updatedAt: response.data.updated_at
	};
}

export async function deleteTech(id: number): Promise<void> {
	await sendRequest<void>('DELETE', `/techs/${id}`);
}

export async function fetchDiary(fetchFn: ApiFetch, id: number): Promise<DiaryEntry> {
	const response = await get<ApiObjectResponse<ApiPost>>(fetchFn, `/diaries/${id}`);
	return {
		id: response.data.id,
		title: response.data.title,
		content: response.data.content,
		createdAt: response.data.created_at,
		updatedAt: response.data.updated_at
	};
}

export async function fetchTech(fetchFn: ApiFetch, id: number): Promise<TechArticle> {
	const response = await get<ApiObjectResponse<ApiPost>>(fetchFn, `/techs/${id}`);
	return {
		id: response.data.id,
		title: response.data.title,
		content: response.data.content,
		category: response.data.category,
		views: response.data.views,
		createdAt: response.data.created_at,
		updatedAt: response.data.updated_at
	};
}

export type ApiTrashPost = {
	id: number;
	type: string;
	title: string;
	content: string;
	category: string;
	created_at: string;
	deleted_at: string;
};

export async function fetchTrashEntries(fetchFn: ApiFetch): Promise<TrashEntry[]> {
	const response = await get<ApiListResponse<ApiTrashPost>>(fetchFn, '/trash');
	return response.data_list.map((post) => ({
		id: post.id,
		type: post.type,
		title: post.title,
		content: post.content,
		category: post.category,
		createdAt: post.created_at,
		deletedAt: post.deleted_at
	}));
}

export async function restoreEntry(id: number): Promise<void> {
	await sendRequest<void>('PUT', `/trash/${id}/restore`);
}
export type ApiApp = {
	id: number;
	slug: string;
	name: string;
	category: string;
	tags: string[];
	description: string;
	icon: string;
	icon_url: string;
	demo_url: string;
	code_url: string;
	created_at: string;
	updated_at: string;
};

export async function fetchAppProjects(fetchFn: ApiFetch): Promise<AppProject[]> {
	const response = await get<ApiListResponse<ApiApp>>(fetchFn, '/apps');
	return response.data_list.map((app) => ({
		id: app.id.toString(),
		slug: app.slug,
		name: app.name,
		category: app.category,
		tags: app.tags,
		description: app.description,
		icon: app.icon,
		iconUrl: app.icon_url,
		demoUrl: app.demo_url,
		codeUrl: app.code_url
	}));
}

export async function fetchApp(fetchFn: ApiFetch, id: number): Promise<AppProject> {
	const response = await get<ApiObjectResponse<ApiApp>>(fetchFn, `/apps/${id}`);
	return {
		id: response.data.id.toString(),
		slug: response.data.slug,
		name: response.data.name,
		category: response.data.category,
		tags: response.data.tags,
		description: response.data.description,
		icon: response.data.icon,
		iconUrl: response.data.icon_url,
		demoUrl: response.data.demo_url,
		codeUrl: response.data.code_url
	};
}

export async function createApp(req: {
	slug: string;
	name: string;
	category: string;
	tags: string[];
	description: string;
	icon: string;
	iconUrl?: string;
	demoUrl?: string;
	codeUrl?: string;
}): Promise<AppProject> {
	const response = await sendRequest<ApiObjectResponse<ApiApp>>('POST', '/apps', {
		slug: req.slug,
		name: req.name,
		category: req.category,
		tags: req.tags,
		description: req.description,
		icon: req.icon,
		icon_url: req.iconUrl || '',
		demo_url: req.demoUrl || '',
		code_url: req.codeUrl || ''
	});

	return {
		id: response.data.id.toString(),
		slug: response.data.slug,
		name: response.data.name,
		category: response.data.category,
		tags: response.data.tags,
		description: response.data.description,
		icon: response.data.icon,
		iconUrl: response.data.icon_url,
		demoUrl: response.data.demo_url,
		codeUrl: response.data.code_url
	};
}

export async function updateApp(id: number, req: {
	slug: string;
	name: string;
	category: string;
	tags: string[];
	description: string;
	icon: string;
	iconUrl?: string;
	demoUrl?: string;
	codeUrl?: string;
}): Promise<AppProject> {
	const response = await sendRequest<ApiObjectResponse<ApiApp>>('PUT', `/apps/${id}`, {
		slug: req.slug,
		name: req.name,
		category: req.category,
		tags: req.tags,
		description: req.description,
		icon: req.icon,
		icon_url: req.iconUrl || '',
		demo_url: req.demoUrl || '',
		code_url: req.codeUrl || ''
	});

	return {
		id: response.data.id.toString(),
		slug: response.data.slug,
		name: response.data.name,
		category: response.data.category,
		tags: response.data.tags,
		description: response.data.description,
		icon: response.data.icon,
		iconUrl: response.data.icon_url,
		demoUrl: response.data.demo_url,
		codeUrl: response.data.code_url
	};
}

export async function deleteApp(id: number): Promise<void> {
	await sendRequest<void>('DELETE', `/apps/${id}`);
}

export async function fetchProfile(fetchFn: ApiFetch): Promise<ProfileData> {
	const response = await get<ApiObjectResponse<{
		id: number;
		name: string;
		subtitle: string;
		title: string;
		avatar_url: string;
		quote: string;
		bio: string[];
		highlights: { title: string; period: string; description: string }[];
		award: string;
		expertise: string[];
		contact_email: string;
		final_quote: string;
	}>>(fetchFn, '/profile');

	return {
		name: response.data.name,
		subtitle: response.data.subtitle,
		title: response.data.title,
		avatarUrl: response.data.avatar_url,
		quote: response.data.quote,
		bio: response.data.bio,
		highlights: response.data.highlights,
		award: response.data.award,
		expertise: response.data.expertise,
		contactEmail: response.data.contact_email,
		finalQuote: response.data.final_quote
	};
}

export async function updateProfile(profile: ProfileData): Promise<ProfileData> {
	const response = await sendRequest<ApiObjectResponse<{
		id: number;
		name: string;
		subtitle: string;
		title: string;
		avatar_url: string;
		quote: string;
		bio: string[];
		highlights: { title: string; period: string; description: string }[];
		award: string;
		expertise: string[];
		contact_email: string;
		final_quote: string;
	}>>('PUT', '/profile', {
		name: profile.name,
		subtitle: profile.subtitle,
		title: profile.title,
		avatar_url: profile.avatarUrl,
		quote: profile.quote,
		bio: profile.bio,
		highlights: profile.highlights,
		award: profile.award || '',
		expertise: profile.expertise,
		contact_email: profile.contactEmail,
		final_quote: profile.finalQuote
	});

	return {
		name: response.data.name,
		subtitle: response.data.subtitle,
		title: response.data.title,
		avatarUrl: response.data.avatar_url,
		quote: response.data.quote,
		bio: response.data.bio,
		highlights: response.data.highlights,
		award: response.data.award,
		expertise: response.data.expertise,
		contactEmail: response.data.contact_email,
		finalQuote: response.data.final_quote
	};
}

