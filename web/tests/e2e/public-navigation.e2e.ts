import { expect, test } from '@playwright/test';

test.describe('公開ページのナビゲーション', () => {
	for (const path of ['/', '/profile', '/apps', '/contact', '/tech']) {
		test(`${path}へ遷移できる`, async ({ page }) => {
			await page.goto(`http://localhost:5173${path}`);
			await expect(page.locator('main')).toBeVisible();
			await expect(page).not.toHaveTitle(/500|Error/);
		});
	}

	test('モバイル・PCの主要viewportでナビゲーションが表示される', async ({ page }) => {
		for (const viewport of [{ width: 375, height: 812 }, { width: 1280, height: 800 }]) {
			await page.setViewportSize(viewport);
			await page.goto('http://localhost:5173/');
			await expect(page.locator('nav').first()).toBeVisible();
		}
	});
});
