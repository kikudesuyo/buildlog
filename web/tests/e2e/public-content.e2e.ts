import { expect, test } from '@playwright/test';

test.describe('公開コンテンツページ', () => {
	test('Appsの正常データを表示する', async ({ page }) => {
		await page.goto('http://localhost:4174/apps');
		await expect(page.getByRole('heading', { name: 'E2Eアプリ' })).toBeVisible();
	});

	test('Techの空状態を表示する', async ({ page }) => {
		await page.goto('http://localhost:4174/tech');
		await expect(page.getByRole('heading', { name: '技術記事はまだありません' })).toBeVisible();
	});

	test('モバイル・PCでコンテンツが画面内に収まる', async ({ page }) => {
		for (const viewport of [{ width: 375, height: 812 }, { width: 1280, height: 800 }]) {
			await page.setViewportSize(viewport);
			await page.goto('http://localhost:4174/profile');
			const overflow = await page.evaluate(() => document.documentElement.scrollWidth > document.documentElement.clientWidth);
			expect(overflow).toBe(false);
		}
	});
});
