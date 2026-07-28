import { test, expect } from '@playwright/test';

test('capture dark mode screenshots', async ({ page }) => {
	await page.goto('http://localhost:5173/');
	await page.waitForLoadState('networkidle');

	await expect(page.locator('html')).not.toHaveClass(/dark/);

	const themeToggle = page.locator('button[aria-label="Toggle Theme"]');
	await themeToggle.click();

	await expect(page.locator('html')).toHaveClass(/dark/);

	await page.waitForTimeout(500);

	await page.screenshot({ path: 'static/screenshots/dark-mode.png', fullPage: true });

	await themeToggle.click();
	await expect(page.locator('html')).not.toHaveClass(/dark/);
});
