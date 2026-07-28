import { test, expect } from '@playwright/test';

test('capture profile screenshots', async ({ page }) => {
	await page.goto('http://localhost:5173/profile');
	await page.waitForLoadState('networkidle');

	const emailLink = page.locator('main a[href^="mailto:"]');
	await expect(emailLink).toBeVisible();

	const githubLink = page.locator('main a[href*="github.com"]');
	await expect(githubLink).toBeVisible();

	await page.screenshot({ path: 'static/screenshots/profile.png', fullPage: true });
});
