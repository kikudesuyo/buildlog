import { test, expect } from '@playwright/test';

test('capture profile screenshots', async ({ page }) => {
	await page.goto('http://localhost:5173/profile');
	await page.waitForLoadState('networkidle');

	await expect(page.getByRole('heading', { name: '自己紹介 / self-introduction', exact: true })).toBeVisible();
	await expect(page.getByRole('heading', { name: '連絡先と繋がり / Contact & Socials', exact: true })).toHaveCount(0);
	await expect(page.locator('#contact')).toHaveCount(0);
	await expect(page.locator('details')).toHaveCount(0);
	await expect(page.getByText('実績を表示', { exact: true })).toHaveCount(0);
	await expect(page.getByText('専門領域を表示', { exact: true })).toHaveCount(0);

	// スクリーンショットを撮影
	await page.screenshot({ path: 'static/screenshots/profile.png', fullPage: true });
});
