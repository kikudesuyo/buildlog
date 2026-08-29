import { expect, test } from '@playwright/test';

test.describe('公開ページの操作', () => {
	test('いいね・コメント表示を操作できる', async ({ page }) => {
		await page.goto('http://localhost:4176/diary/1');
		await expect(page.getByText('既存コメント')).toBeVisible();
		const like = page.getByRole('button', { name: 'いいねボタン' });
		await like.click();
		await expect(like).toContainText('2');
	});

	test('日記一覧の追加読み込みを検証できる', async ({ page }) => {
		await page.goto('http://localhost:4176/');
		await expect(page.getByRole('heading', { name: '操作対象の日記' })).toBeVisible();
		const more = page.getByRole('button', { name: '過去の記録を見る' });
		await expect(more).toBeVisible();
		await more.click();
		await expect(page.getByRole('heading', { name: '追加日記4' })).toBeVisible();
	});
});
