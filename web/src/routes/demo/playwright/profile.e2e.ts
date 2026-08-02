import { test, expect } from '@playwright/test';

test('capture profile screenshots', async ({ page }) => {
	await page.goto('http://localhost:5173/profile');
	await page.waitForLoadState('networkidle');

	// メールアドレスのコピー用ボタンが存在することを確認 (テキストに @ が含まれているはず)
	const emailButton = page.locator('main button:has-text("@")');
	await expect(emailButton).toBeVisible();

	// スクリーンショットを撮影
	await page.screenshot({ path: 'static/screenshots/profile.png', fullPage: true });
});
