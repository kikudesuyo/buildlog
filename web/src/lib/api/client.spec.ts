import { afterEach, describe, expect, it, vi } from 'vitest';

import {
	createDiary,
	fetchDiary,
	fetchDiaryEntries,
	fetchTechFeed,
	saveCurrentGoals
} from '$lib/api/client';
import type { ApiFetch } from '$lib/api/client';

const diaryPost = {
	id: 1,
	type: 'diary',
	title: 'テスト日記',
	content: '本文',
	excerpt: '抜粋',
	category: 'daily',
	views: 3,
	status: 'published' as const,
	created_at: '2026-01-01T00:00:00Z',
	updated_at: '2026-01-02T00:00:00Z',
	likes_count: 2,
	comments_count: 1,
	has_liked: true
};

const techPost = {
	...diaryPost,
	id: 2,
	type: 'tech',
	key: 'tech-2',
	title: 'テック記事',
	views: 10,
	external: {
		provider: 'qiita',
		url: 'https://example.com/article',
		thumbnail_url: 'https://example.com/image.png'
	}
};

function apiFetch(data: unknown, status = 200): ApiFetch {
	return vi.fn().mockResolvedValue(new Response(JSON.stringify(data), { status })) as unknown as ApiFetch;
}

function jsonResponse(data: unknown, status = 200): Response {
	return new Response(JSON.stringify(data), {
		status,
		headers: { 'Content-Type': 'application/json' }
	});
}

afterEach(() => {
	vi.restoreAllMocks();
});

describe('fetchDiaryEntries', () => {
	it('maps API entries and sends pagination and sort parameters', async () => {
		const fetchFn = apiFetch({ data_list: [diaryPost] });

		const result = await fetchDiaryEntries(fetchFn, false, 5, 10, 'likes', 'asc');

		expect(result).toEqual([
			{
				id: 1,
				title: 'テスト日記',
				content: '本文',
				status: 'published',
				createdAt: '2026-01-01T00:00:00Z',
				updatedAt: '2026-01-02T00:00:00Z',
				likesCount: 2,
				commentsCount: 1,
				hasLiked: true
			}
		]);
		expect(fetchFn).toHaveBeenCalledWith('/api/v1/diaries?offset=5&limit=10&sort=likes&order=asc');
	});

	it('returns an empty list for an empty API response', async () => {
		const result = await fetchDiaryEntries(apiFetch({ data_list: [] }));

		expect(result).toEqual([]);
	});
});

describe('fetchTechFeed', () => {
	it('returns all articles in the same list and reports an additional page', async () => {
		const fetchFn = apiFetch({ data_list: [techPost, { ...techPost, id: 3, key: undefined }] });

		const result = await fetchTechFeed(fetchFn, false, 0, 1);

		expect(result.techArticles.map((article) => article.key)).toEqual(['tech-2']);
		expect(result.hasMore).toBe(true);
		expect(fetchFn).toHaveBeenCalledWith('/api/v1/techs?limit=1');
	});

	it('returns an empty list when there are no posts', async () => {
		const result = await fetchTechFeed(apiFetch({ data_list: [] }));

		expect(result.techArticles).toEqual([]);
		expect(result.hasMore).toBe(false);
	});
});

describe('client API errors and writes', () => {
	it('throws a status-aware error when a fetch request fails', async () => {
		const fetchFn = apiFetch({ message: 'failed' }, 503);

		await expect(fetchDiary(fetchFn, 42)).rejects.toThrow('API request failed: 503');
	});

	it('sends a diary payload and maps the created response', async () => {
		const fetchMock = vi.fn().mockResolvedValue(jsonResponse({ data: diaryPost }));
		vi.stubGlobal('fetch', fetchMock);

		const result = await createDiary('タイトル', '本文', 'published');

		expect(result.id).toBe(1);
		expect(fetchMock).toHaveBeenCalledWith('/api/v1/diaries', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ title: 'タイトル', content: '本文', status: 'published' })
		});
	});

	it('converts goal fields to the API format and maps the response', async () => {
		const fetchMock = vi.fn().mockResolvedValue(
			jsonResponse({
				data: {
					period_type: 'monthly',
					starts_at: '2026-01-01',
					ends_at: '2026-01-31',
					goals: [{ id: 1, title: '読む', target_value: 5, progress_value: 2 }]
				}
			})
		);
		vi.stubGlobal('fetch', fetchMock);

		const result = await saveCurrentGoals([{ title: '読む', targetValue: 5, progressValue: 2 }]);

		expect(result.goals[0]).toEqual({ id: 1, title: '読む', targetValue: 5, progressValue: 2 });
		expect(fetchMock).toHaveBeenCalledWith('/api/v1/goals/current', {
			method: 'PUT',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({
				period_type: 'monthly',
				goals: [{ title: '読む', target_value: 5, progress_value: 2 }]
			})
		});
	});
});
