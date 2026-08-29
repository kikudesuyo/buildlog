import { page } from 'vitest/browser';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import { render } from 'vitest-browser-svelte';
import LikeButton from './LikeButton.svelte';

const { likePost, unlikePost } = vi.hoisted(() => ({
	likePost: vi.fn(),
	unlikePost: vi.fn()
}));
vi.mock('$lib/api/client', () => ({ likePost, unlikePost }));

describe('LikeButton.svelte', () => {
	beforeEach(() => {
		likePost.mockReset();
		unlikePost.mockReset();
		localStorage.clear();
	});

	it('初期いいね数を表示し、いいね成功後の状態を反映する', async () => {
		likePost.mockResolvedValue({ likes_count: 4, has_liked: true });
		render(LikeButton, { postId: 7, initialLikesCount: 3 });

		await expect.element(page.getByRole('button', { name: 'いいねボタン' })).toHaveTextContent('3');
		await page.getByRole('button', { name: 'いいねボタン' }).click();
		await expect.element(page.getByRole('button', { name: 'いいねボタン' })).toHaveTextContent('4');
		expect(likePost).toHaveBeenCalledWith(7);
	});

	it('いいね解除後の状態を反映する', async () => {
		unlikePost.mockResolvedValue({ likes_count: 1, has_liked: false });
		render(LikeButton, { postId: 7, initialLikesCount: 2, initialHasLiked: true });

		await page.getByRole('button', { name: 'いいねボタン' }).click();
		await expect.element(page.getByRole('button', { name: 'いいねボタン' })).toHaveTextContent('1');
		expect(unlikePost).toHaveBeenCalledWith(7);
	});

	it('APIエラー時は元の状態へ戻す', async () => {
		likePost.mockRejectedValue(new Error('network error'));
		render(LikeButton, { postId: 7, initialLikesCount: 3 });

		await page.getByRole('button', { name: 'いいねボタン' }).click();
		await expect.element(page.getByRole('button', { name: 'いいねボタン' })).toHaveTextContent('3');
	});
});
