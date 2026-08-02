import { test } from '@playwright/test';

test('capture apps screenshots', async ({ page }) => {
	// 管理画面 (一覧)
	await page.goto('http://localhost:5173/admin/apps');
	await page.waitForLoadState('networkidle');
	await page.screenshot({ path: 'static/screenshots/admin-apps-list.png', fullPage: true });

	// 管理画面 (新規作成フォーム)
	await page.goto('http://localhost:5173/admin/apps/new');
	await page.waitForLoadState('networkidle');
	await page.screenshot({ path: 'static/screenshots/admin-apps-new.png', fullPage: true });

	// 公開の Showcase 画面
	await page.goto('http://localhost:5173/apps');
	await page.waitForLoadState('networkidle');
	await page.screenshot({ path: 'static/screenshots/public-apps.png', fullPage: true });
});
