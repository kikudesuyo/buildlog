import { page } from 'vitest/browser';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import { render } from 'vitest-browser-svelte';
import type { CommentEntry } from '$lib/api/types';
import CommentSection from './CommentSection.svelte';

const { fetchComments, createComment } = vi.hoisted(() => ({
	fetchComments: vi.fn(),
	createComment: vi.fn()
}));
vi.mock('$lib/api/client', () => ({ fetchComments, createComment }));

const comment: CommentEntry = {
	id: 1,
	postId: 42,
	content: '参考になる記事でした。',
	createdAt: '2026-08-12T09:00:00.000Z',
	updatedAt: '2026-08-12T09:00:00.000Z'
};

describe('CommentSection.svelte', () => {
	beforeEach(() => {
		fetchComments.mockReset();
		createComment.mockReset();
	});

	it('コメント一覧を表示し、投稿したコメントを追加する', async () => {
		fetchComments.mockResolvedValue([comment]);
		createComment.mockResolvedValue({ ...comment, id: 2, content: '追加コメント' });
		render(CommentSection, { postId: 42 });

		await expect.element(page.getByText('参考になる記事でした。')).toBeInTheDocument();
		const input = page.getByLabelText('コメント内容');
		await input.fill('追加コメント');
		await page.getByRole('button', { name: 'コメントを送信' }).click();
		await expect.element(page.getByText('追加コメント')).toBeInTheDocument();
		expect(createComment).toHaveBeenCalledWith(42, '追加コメント');
	});

	it('コメントがない場合は空状態を表示する', async () => {
		fetchComments.mockResolvedValue([]);
		render(CommentSection, { postId: 42 });

		await expect.element(page.getByText('最初のコメントを投稿しましょう。')).toBeInTheDocument();
	});

	it('コメント取得エラーを表示する', async () => {
		fetchComments.mockRejectedValue(new Error('network error'));
		render(CommentSection, { postId: 42 });

		await expect.element(page.getByRole('alert')).toHaveTextContent('コメントを読み込めませんでした。');
	});
});
