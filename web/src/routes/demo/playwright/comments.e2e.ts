import { test, expect } from '@playwright/test';

test('posts a comment and a one-level reply', async ({ page }) => {
	await page.goto('http://localhost:5173/tech/5');

	const commentContent = page.locator('#comment-content');
	await expect(commentContent).toBeVisible();
	await commentContent.fill('自動テストからのコメントです');
	await page.getByRole('button', { name: 'コメントを送信' }).click();

	const comment = page.getByText('自動テストからのコメントです');
	await expect(comment).toBeVisible();
	await comment.locator('..').getByRole('button', { name: '返信する' }).click();

	const replyContent = page.locator('textarea[id^="reply-"]');
	await replyContent.fill('自動テストからの返信です');
	await page.getByRole('button', { name: '返信を送信' }).click();

	await expect(page.getByText('自動テストからの返信です')).toBeVisible();
});
