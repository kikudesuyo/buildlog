import { page } from 'vitest/browser';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import { render } from 'vitest-browser-svelte';
import type { TechArticle } from '$lib/api/types';
import TechFeed from './TechFeed.svelte';

const { fetchTechFeed } = vi.hoisted(() => ({ fetchTechFeed: vi.fn() }));

vi.mock('$lib/api/client', () => ({ fetchTechFeed }));
vi.mock('$app/navigation', () => ({ invalidateAll: vi.fn(), replaceState: vi.fn() }));
vi.mock('$app/paths', () => ({ resolve: (path: string) => path }));
vi.mock('$app/stores', async () => {
	const { writable } = await import('svelte/store');
	return { page: writable({ url: new URL('http://localhost/tech') }) };
});

const makeArticle = (overrides: Partial<TechArticle> = {}): TechArticle => ({
	key: 'article-1',
	id: 1,
	title: '技術記事タイトル',
	content: '技術記事の本文です。',
	createdAt: '2026-08-12T00:00:00.000Z',
	updatedAt: '2026-08-12T00:00:00.000Z',
	likesCount: 0,
	commentsCount: 0,
	status: 'published',
	...overrides
});

describe('TechFeed.svelte', () => {
	beforeEach(() => {
		fetchTechFeed.mockReset();
		window.sessionStorage.clear();
	});

	it('記事と並び順を表示する', async () => {
		const firstArticle: TechArticle = makeArticle({
			key: 'featured-article',
			title: '注目の記事'
		});
		const article = makeArticle({ key: 'article-2', id: 2, title: '一覧の記事' });

		render(TechFeed, { techArticles: [firstArticle, article] });

		await expect.element(page.getByRole('heading', { name: '技術ブログ' })).toBeInTheDocument();
		await expect.element(page.getByRole('heading', { name: '注目の記事' })).toBeInTheDocument();
		await expect.element(page.getByRole('heading', { name: '一覧の記事' })).toBeInTheDocument();
		await expect.element(page.getByRole('combobox', { name: '技術記事の並び順' })).toHaveValue('desc');
	});

	it('記事がない場合は空状態を表示する', async () => {
		render(TechFeed, { techArticles: [], isAdmin: true });

		await expect.element(page.getByRole('heading', { name: '技術記事はまだありません' })).toBeInTheDocument();
		await expect.element(page.getByText('新しい記事が公開されるまでお待ちください。')).toBeInTheDocument();
	});

	it('初期読み込みエラー時は再試行案内を表示する', async () => {
		render(TechFeed, { loadError: true, isAdmin: true });

		await expect.element(page.getByRole('alert')).toHaveTextContent('記事を読み込めませんでした');
		await expect.element(page.getByRole('button', { name: '再試行' })).toBeInTheDocument();
	});
});
