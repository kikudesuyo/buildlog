import { test } from '@playwright/test';

test('capture profile screenshots', async ({ page }) => {
	// 管理者プロフィール編集画面のスクリーンショットを撮影
	await page.goto('http://localhost:5173/admin/profile');
	await page.waitForLoadState('networkidle');
	await page.screenshot({ path: 'static/screenshots/admin-profile.png', fullPage: true });

	// 公開プロフィール画面のスクリーンショットを撮影
	await page.goto('http://localhost:5173/profile');
	await page.waitForLoadState('networkidle');
	await page.screenshot({ path: 'static/screenshots/public-profile.png', fullPage: true });
});
