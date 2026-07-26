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
	created_at: string;
	updated_at: string;
	likes_count: number;
	has_liked: boolean;
};

export async function fetchDiaryEntries(fetchFn: ApiFetch): Promise<DiaryEntry[]> {
	const response = await get<ApiListResponse<ApiPost>>(fetchFn, '/diaries');
	return response.data_list.map((post) => ({
		id: post.id,
		title: post.title,
		content: post.content,
		createdAt: post.created_at,
		updatedAt: post.updated_at,
		likesCount: post.likes_count,
		hasLiked: post.has_liked
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
		updatedAt: post.updated_at,
		likesCount: post.likes_count,
		hasLiked: post.has_liked
	}));

	const featured = allArticles.length > 0 ? allArticles[0] : {
		id: 0,
		title: '',
		content: '',
		category: '',
		views: '',
		createdAt: '',
		updatedAt: '',
		likesCount: 0,
		hasLiked: false
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
		updatedAt: response.data.updated_at,
		likesCount: response.data.likes_count,
		hasLiked: response.data.has_liked
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
		updatedAt: response.data.updated_at,
		likesCount: response.data.likes_count,
		hasLiked: response.data.has_liked
	};
}

export type ApiLikeStatus = {
	likes_count: number;
	has_liked: boolean;
};

export async function likePost(id: number): Promise<ApiLikeStatus> {
	const response = await sendRequest<ApiObjectResponse<ApiLikeStatus>>('POST', `/posts/${id}/like`);
	return response.data;
}

export async function unlikePost(id: number): Promise<ApiLikeStatus> {
	const response = await sendRequest<ApiObjectResponse<ApiLikeStatus>>('DELETE', `/posts/${id}/like`);
	return response.data;
}
