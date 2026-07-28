import { test, expect } from '@playwright/test';

test('verify public Navbar does not contain admin or settings icons and capture screenshots', async ({ page }) => {
	await page.goto('http://localhost:5173/');
	await page.waitForLoadState('networkidle');

	// 一般向け Navbar の検証
	const navbar = page.locator('nav');
	await expect(navbar).toBeVisible();

	// settings (歯車) アイコンが存在しないことを確認
	const settingsIcon = page.locator('button[aria-label="Settings"]');
	await expect(settingsIcon).not.toBeVisible();

	// admin_panel_settings アイコンが存在しないことを確認
	const adminIcon = page.locator('a[title="管理画面に切り替え"]');
	await expect(adminIcon).not.toBeVisible();

	// 検索 (search) アイコンは存在することを確認
	const searchIcon = page.locator('button[aria-label="Search"]');
	await expect(searchIcon).toBeVisible();

	// スクリーンショット撮影
	await page.screenshot({ path: 'static/screenshots/navbar-cleaned.png', fullPage: false });
});
