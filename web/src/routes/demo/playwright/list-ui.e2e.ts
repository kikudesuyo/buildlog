import { test } from '@playwright/test';

test('capture list UI screenshots', async ({ page }) => {
	await page.goto('http://localhost:5173/');
	await page.waitForLoadState('networkidle');
	await page.screenshot({ path: 'static/screenshots/diary-feed.png', fullPage: true });

	await page.goto('http://localhost:5173/tech');
	await page.waitForLoadState('networkidle');
	await page.screenshot({ path: 'static/screenshots/tech-feed.png', fullPage: true });
});
