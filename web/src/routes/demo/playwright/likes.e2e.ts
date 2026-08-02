import { test, expect } from '@playwright/test';

test('capture likes screenshots', async ({ page }) => {
	await page.goto('http://localhost:5173/');
	await page.waitForLoadState('networkidle');

	const likeButtons = page.locator('button[aria-label="いいねボタン"]');
	await expect(likeButtons.first()).toBeVisible();

	await page.screenshot({ path: 'static/screenshots/likes-before.png' });

	await likeButtons.first().click();
	await page.waitForTimeout(500);

	await page.screenshot({ path: 'static/screenshots/likes-after.png' });
});
