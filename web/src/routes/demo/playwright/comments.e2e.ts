import { test, expect } from '@playwright/test';

test('posts a comment', async ({ page }) => {
	await page.goto('http://localhost:5173/tech/5');

	const commentContent = page.locator('#comment-content');
	await expect(commentContent).toBeVisible();
	await commentContent.fill('自動テストからのコメントです');
	await page.getByRole('button', { name: 'コメントを送信' }).click();

	const comment = page.getByText('自動テストからのコメントです');
	await expect(comment).toBeVisible();
	await expect(page.getByRole('button', { name: '返信する' })).toHaveCount(0);
});
