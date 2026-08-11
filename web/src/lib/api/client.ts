import type { LoadEvent } from '@sveltejs/kit';
import type {
	DiaryEntry,
	FeaturedTechArticle,
	TechArticle,
	TrashEntry,
	AppProject,
	ProfileData,
	AnalyticsData,
	HistoryItem,
	CommentEntry,
	GoalPeriod
} from '$lib/api/types';

export type ApiFetch = LoadEvent['fetch'];

type ApiListResponse<T> = {
	data_list: T[];
};

type ApiObjectResponse<T> = {
	data: T;
};




const apiBaseUrl = '/api/v1';

async function get<T>(fetchFn: ApiFetch, path: string): Promise<T> {
	const response = await fetchFn(`${apiBaseUrl}${path}`);
	if (!response.ok) {
		throw new Error(`API request failed: ${response.status} ${response.statusText}`);
	}
	return response.json() as Promise<T>;
}

export type ApiPost = {
	key?: string;
	id: number;
	type: string;
	title: string;
	content: string;
	excerpt: string;
	category: string;
	views: number;
	status: 'draft' | 'published';
	created_at: string;
	updated_at: string;
	likes_count: number;
	comments_count: number;
	has_liked: boolean;
	external?: {
		provider: string;
		url: string;
		thumbnail_url: string;
	};
};

export type DiarySort = 'newest' | 'likes';
export type DiarySortOrder = 'asc' | 'desc';
export type TechSortOrder = 'asc' | 'desc';

export async function fetchDiaryEntries(
	fetchFn: ApiFetch,
	all = false,
	offset = 0,
	limit = 0,
	sort: DiarySort = 'newest',
	order: DiarySortOrder = 'desc'
): Promise<DiaryEntry[]> {
	const params = new URLSearchParams();
	if (all) params.set('all', 'true');
	if (!all && offset > 0) params.set('offset', String(offset));
	if (!all && limit > 0) params.set('limit', String(limit));
	if (sort !== 'newest') params.set('sort', sort);
	if (order !== 'desc') params.set('order', order);
	const query = params.toString();
	const url = query ? `/diaries?${query}` : '/diaries';
	const response = await get<ApiListResponse<ApiPost>>(fetchFn, url);
	return response.data_list.map((post) => ({
		id: post.id,
		title: post.title,
		content: post.content,
		status: post.status,
		createdAt: post.created_at,
		updatedAt: post.updated_at,
		likesCount: post.likes_count,
		commentsCount: post.comments_count,
		hasLiked: post.has_liked
	}));
}

export async function fetchTechFeed(fetchFn: ApiFetch, all = false, offset = 0, limit = 0, order: TechSortOrder = 'desc'): Promise<{
	featuredArticle: FeaturedTechArticle | null;
	techArticles: TechArticle[];
	hasMore: boolean;
}> {
	const params = new URLSearchParams();
	if (all) params.set('all', 'true');
	if (!all && offset && offset > 0) params.set('offset', String(offset));
	if (!all && limit && limit > 0) params.set('limit', String(limit));
	if (order !== 'desc') params.set('order', order);
	const query = params.toString();
	const url = query ? `/techs?${query}` : '/techs';
	const response = await get<ApiListResponse<ApiPost>>(fetchFn, url);
	const hasMore = !all && !!limit && response.data_list.length > limit;
	const page = limit && limit > 0 ? response.data_list.slice(0, limit) : response.data_list;
	
	const allArticles: TechArticle[] = page.map((post) => ({
		key: post.key ?? `post:${post.id}`,
		id: post.id,
		title: post.title,
		content: post.content,
		views: post.views,
		status: post.status,
		createdAt: post.created_at,
		updatedAt: post.updated_at,
		likesCount: post.likes_count,
		commentsCount: post.comments_count,
		hasLiked: post.has_liked,
		external: post.external
			? { provider: post.external.provider, url: post.external.url, thumbnailUrl: post.external.thumbnail_url }
			: undefined
	}));

	const featured = !offset && allArticles.length > 0 ? allArticles[0] : null;
	const fallbackFeatured: FeaturedTechArticle = {
		key: 'fallback',
		id: 0,
		title: '',
		content: '',
		views: 0,
		status: 'draft' as const,
		createdAt: '',
		updatedAt: '',
		likesCount: 0,
		commentsCount: 0,
		hasLiked: false
	};
	const remaining = featured ? allArticles.slice(1) : allArticles;

	return {
		featuredArticle: featured ?? (offset ? null : fallbackFeatured),
		techArticles: remaining,
		hasMore
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
		updatedAt: response.data.updated_at,
		likesCount: response.data.likes_count,
		commentsCount: response.data.comments_count ?? 0,
		hasLiked: response.data.has_liked
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
		updatedAt: response.data.updated_at,
		likesCount: response.data.likes_count,
		commentsCount: response.data.comments_count ?? 0,
		hasLiked: response.data.has_liked
	};
}

export async function deleteDiary(id: number): Promise<void> {
	await sendRequest<void>('DELETE', `/diaries/${id}`);
}

export async function syncQiitaArticles(): Promise<number> {
	const response = await sendRequest<ApiObjectResponse<{ synced: number }>>('POST', '/admin/tech/qiita/sync');
	return response.data.synced;
}

export async function fetchDiary(fetchFn: ApiFetch, id: number, countView = false): Promise<DiaryEntry> {
	const response = await get<ApiObjectResponse<ApiPost>>(fetchFn, `/diaries/${id}${countView ? '?count_view=true' : ''}`);
	return {
		id: response.data.id,
		title: response.data.title,
		content: response.data.content,
		status: response.data.status,
		createdAt: response.data.created_at,
		updatedAt: response.data.updated_at,
		likesCount: response.data.likes_count,
		commentsCount: response.data.comments_count ?? 0,
		hasLiked: response.data.has_liked
	};
}

export type ApiLikeStatus = {
	likes_count: number;
	has_liked: boolean;
};

type ApiComment = {
	id: number;
	post_id: number;
	content: string;
	created_at: string;
	updated_at: string;
};

function mapComment(comment: ApiComment): CommentEntry {
	return {
		id: comment.id,
		postId: comment.post_id,
		content: comment.content,
		createdAt: comment.created_at,
		updatedAt: comment.updated_at
	};
}

export async function fetchComments(fetchFn: ApiFetch, postId: number): Promise<CommentEntry[]> {
	const response = await get<ApiListResponse<ApiComment>>(fetchFn, `/posts/${postId}/comments`);
	return response.data_list.map(mapComment);
}

export async function createComment(
	postId: number,
	content: string
): Promise<CommentEntry> {
	const response = await sendRequest<ApiObjectResponse<ApiComment>>('POST', `/posts/${postId}/comments`, { content });
	return mapComment(response.data);
}

export async function likePost(id: number): Promise<ApiLikeStatus> {
	const response = await sendRequest<ApiObjectResponse<ApiLikeStatus>>('POST', `/posts/${id}/like`);
	return response.data;
}

export async function unlikePost(id: number): Promise<ApiLikeStatus> {
	const response = await sendRequest<ApiObjectResponse<ApiLikeStatus>>('DELETE', `/posts/${id}/like`);
	return response.data;
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

type ApiGoal = {
	id: number;
	title: string;
	target_value: number;
	progress_value: number;
};

type ApiGoalPeriod = {
	period_type: 'monthly';
	starts_at: string;
	ends_at: string;
	goals: ApiGoal[];
};

function mapGoalPeriod(period: ApiGoalPeriod): GoalPeriod {
	return {
		periodType: period.period_type,
		startsAt: period.starts_at,
		endsAt: period.ends_at,
		goals: period.goals.map((goal) => ({
			id: goal.id,
			title: goal.title,
			targetValue: goal.target_value,
			progressValue: goal.progress_value
		}))
	};
}

export async function fetchCurrentGoals(fetchFn: ApiFetch): Promise<GoalPeriod> {
	const response = await get<ApiObjectResponse<ApiGoalPeriod>>(fetchFn, '/goals/current');
	return mapGoalPeriod(response.data);
}

export async function saveCurrentGoals(goals: Array<{ title: string; targetValue: number; progressValue: number }>): Promise<GoalPeriod> {
	const response = await sendRequest<ApiObjectResponse<ApiGoalPeriod>>('PUT', '/goals/current', {
		period_type: 'monthly',
		goals: goals.map((goal) => ({
			title: goal.title,
			target_value: goal.targetValue,
			progress_value: goal.progressValue
		}))
	});
	return mapGoalPeriod(response.data);
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
		title: string;
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
		title: response.data.title,
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
		title: string;
		quote: string;
		bio: string[];
		highlights: { title: string; period: string; description: string }[];
		award: string;
		expertise: string[];
		contact_email: string;
		final_quote: string;
	}>>('PUT', '/profile', {
		name: profile.name,
		title: profile.title,
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
		title: response.data.title,
		quote: response.data.quote,
		bio: response.data.bio,
		highlights: response.data.highlights,
		award: response.data.award,
		expertise: response.data.expertise,
		contactEmail: response.data.contact_email,
		finalQuote: response.data.final_quote
	};
}

export async function fetchAnalytics(fetchFn?: ApiFetch): Promise<AnalyticsData> {
	const response = await get<ApiObjectResponse<{
		total_views: number;
		total_likes: number;
		total_posts: number;
		diary_count: number;
		tech_count: number;
		top_views_articles: {
			id: number;
			type: string;
			title: string;
			views: number;
			likes: number;
		}[];
		top_likes_articles: {
			id: number;
			type: string;
			title: string;
			views: number;
			likes: number;
		}[];
		monthly_activities: {
			month: string;
			count: number;
		}[];
	}>>(fetchFn || fetch, '/admin/analytics');

	return {
		totalViews: response.data.total_views,
		totalLikes: response.data.total_likes,
		totalPosts: response.data.total_posts,
		diaryCount: response.data.diary_count,
		techCount: response.data.tech_count,
		topViewsArticles: response.data.top_views_articles.map((item) => ({
			id: item.id,
			type: item.type,
			title: item.title,
			views: item.views,
			likes: item.likes
		})),
		topLikesArticles: response.data.top_likes_articles.map((item) => ({
			id: item.id,
			type: item.type,
			title: item.title,
			views: item.views,
			likes: item.likes
		})),
		monthlyActivities: response.data.monthly_activities.map((item) => ({
			month: item.month,
			count: item.count
		}))
	};
}

export async function fetchPostHistory(fetchFn?: ApiFetch): Promise<HistoryItem[]> {
	const response = await get<ApiListResponse<{
		id: number;
		type: string;
		title: string;
		created_at: string;
		url?: string;
	}>>(fetchFn || fetch, '/posts/history');

	return response.data_list.map((item) => ({
		id: item.id,
		type: item.type as 'diary' | 'tech',
		title: item.title,
		createdAt: item.created_at,
		url: item.url
	}));
}
