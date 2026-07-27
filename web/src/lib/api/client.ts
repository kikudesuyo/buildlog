import * as env from '$env/static/public';
import type { LoadEvent } from '@sveltejs/kit';
import type {
	DiaryEntry,
	FeaturedTechArticle,
	TechArticle
} from '$lib/api/types';

type ApiFetch = LoadEvent['fetch'];

type ApiListResponse<T> = {
	data_list: T[];
};

type ApiObjectResponse<T> = {
	data: T;
};




const apiBaseUrl = (() => {
	const rawUrl = env.PUBLIC_API_BASE_URL || 'http://localhost:8081';
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
	status: 'draft' | 'published';
	created_at: string;
	updated_at: string;
};

export async function fetchDiaryEntries(fetchFn: ApiFetch, all = false): Promise<DiaryEntry[]> {
	const url = all ? '/diaries?all=true' : '/diaries';
	const response = await get<ApiListResponse<ApiPost>>(fetchFn, url);
	return response.data_list.map((post) => ({
		id: post.id,
		title: post.title,
		content: post.content,
		status: post.status,
		createdAt: post.created_at,
		updatedAt: post.updated_at
	}));
}

export async function fetchTechFeed(fetchFn: ApiFetch, all = false): Promise<{
	featuredArticle: FeaturedTechArticle;
	techArticles: TechArticle[];
}> {
	const url = all ? '/techs?all=true' : '/techs';
	const response = await get<ApiListResponse<ApiPost>>(fetchFn, url);
	
	const allArticles: TechArticle[] = response.data_list.map((post) => ({
		id: post.id,
		title: post.title,
		content: post.content,
		category: post.category,
		views: post.views,
		status: post.status,
		createdAt: post.created_at,
		updatedAt: post.updated_at
	}));

	const featured = allArticles.length > 0 ? allArticles[0] : {
		id: 0,
		title: '',
		content: '',
		category: '',
		views: '',
		status: 'draft' as const,
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

export async function createDiary(title: string, content: string, status?: 'draft' | 'published'): Promise<DiaryEntry> {
	const response = await sendRequest<ApiObjectResponse<ApiPost>>('POST', '/diaries', { title, content, status: status || 'draft' });
	return {
		id: response.data.id,
		title: response.data.title,
		content: response.data.content,
		status: response.data.status,
		createdAt: response.data.created_at,
		updatedAt: response.data.updated_at
	};
}

export async function updateDiary(id: number, title: string, content: string, status?: 'draft' | 'published'): Promise<DiaryEntry> {
	const response = await sendRequest<ApiObjectResponse<ApiPost>>('PUT', `/diaries/${id}`, { title, content, status: status || 'draft' });
	return {
		id: response.data.id,
		title: response.data.title,
		content: response.data.content,
		status: response.data.status,
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
	status?: 'draft' | 'published';
}): Promise<TechArticle> {
	const response = await sendRequest<ApiObjectResponse<ApiPost>>('POST', '/techs', {
		title: req.title,
		content: req.content,
		category: req.category,
		views: req.views || '',
		status: req.status || 'draft'
	});
	return {
		id: response.data.id,
		title: response.data.title,
		content: response.data.content,
		category: response.data.category,
		views: response.data.views,
		status: response.data.status,
		createdAt: response.data.created_at,
		updatedAt: response.data.updated_at
	};
}

export async function updateTech(id: number, req: {
	title: string;
	content: string;
	category: string;
	views?: string;
	status?: 'draft' | 'published';
}): Promise<TechArticle> {
	const response = await sendRequest<ApiObjectResponse<ApiPost>>('PUT', `/techs/${id}`, {
		title: req.title,
		content: req.content,
		category: req.category,
		views: req.views || '',
		status: req.status || 'draft'
	});
	return {
		id: response.data.id,
		title: response.data.title,
		content: response.data.content,
		category: response.data.category,
		views: response.data.views,
		status: response.data.status,
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
		status: response.data.status,
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
		status: response.data.status,
		createdAt: response.data.created_at,
		updatedAt: response.data.updated_at
	};
}
