import { test, expect } from '@playwright/test';

const pages = [
	{ name: 'home', path: '/' },
	{ name: 'tech_list', path: '/tech' },
	{ name: 'tech_detail', path: '/tech/10' },
	{ name: 'profile', path: '/profile' },
	{ name: 'apps', path: '/apps' },
	{ name: 'admin', path: '/admin' }
];

test('capture debug screenshots for PC', async ({ page }) => {
	await page.setViewportSize({ width: 1200, height: 800 });
	for (const p of pages) {
		try {
			await page.goto(`http://localhost:5173${p.path}`);
			await page.waitForLoadState('networkidle');
			await page.waitForTimeout(500);
			await page.screenshot({ path: `static/screenshots/ui_debug/pc_${p.name}.png`, fullPage: true });
		} catch (e) {
			console.error(`Failed to capture PC ${p.name}:`, e);
		}
	}
});

test('capture debug screenshots for Mobile', async ({ page }) => {
	await page.setViewportSize({ width: 390, height: 844 });
	for (const p of pages) {
		try {
			await page.goto(`http://localhost:5173${p.path}`);
			await page.waitForLoadState('networkidle');
			await page.waitForTimeout(500);
			await page.screenshot({ path: `static/screenshots/ui_debug/mobile_${p.name}.png`, fullPage: true });
		} catch (e) {
			console.error(`Failed to capture Mobile ${p.name}:`, e);
		}
	}
});
