import { page } from 'vitest/browser';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import type { DiaryEntry } from '$lib/api/types';
import { render } from 'vitest-browser-svelte';
import DiaryFeed from './DiaryFeed.svelte';

const { fetchDiaryEntries, likePost, unlikePost } = vi.hoisted(() => ({
	fetchDiaryEntries: vi.fn(),
	likePost: vi.fn(),
	unlikePost: vi.fn()
}));
vi.mock('$lib/api/client', () => ({ fetchDiaryEntries, likePost, unlikePost }));
vi.mock('$app/navigation', () => ({ replaceState: vi.fn() }));
vi.mock('$app/paths', () => ({ resolve: (path: string) => path }));
vi.mock('$app/stores', async () => {
	const { writable } = await import('svelte/store');
	return { page: writable({ url: new URL('http://localhost/diary') }) };
});

const makeEntry = (id: number): DiaryEntry => ({
	id,
	title: '日記' + id,
	content: '本文' + id,
	createdAt: '2026-08-12T00:00:00.000Z',
	updatedAt: '2026-08-12T00:00:00.000Z',
	likesCount: 0,
	commentsCount: 0,
	status: 'published'
});

describe('DiaryFeed.svelte', () => {
	beforeEach(() => {
		fetchDiaryEntries.mockReset();
		likePost.mockResolvedValue({ likes_count: 1, has_liked: true });
		unlikePost.mockResolvedValue({ likes_count: 0, has_liked: false });
		sessionStorage.clear();
	});

	it('日記一覧を表示する', async () => {
		render(DiaryFeed, { entries: [makeEntry(1)], isAdmin: true });

		await expect.element(page.getByRole('heading', { name: 'つぶやき管理' })).toBeInTheDocument();
		await expect.element(page.getByRole('heading', { name: '日記1' })).toBeInTheDocument();
	});

	it('空の公開一覧では完了状態を表示する', async () => {
		render(DiaryFeed, { entries: [] });

		await expect.element(page.getByRole('status')).toHaveTextContent('すべての記録を表示しました。');
	});

	it('追加読み込みエラーを表示する', async () => {
		fetchDiaryEntries.mockRejectedValue(new Error('network error'));
		render(DiaryFeed, { entries: [makeEntry(1), makeEntry(2), makeEntry(3)] });

		await page.getByRole('button', { name: '過去の記録を見る' }).click();
		await expect.element(page.getByRole('alert')).toHaveTextContent('過去の記録を読み込めませんでした。');
		expect(fetchDiaryEntries).toHaveBeenCalled();
	});
});
