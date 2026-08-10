import { test, expect } from '@playwright/test';

test('capture profile screenshots', async ({ page }) => {
	await page.goto('http://localhost:5173/profile');
	await page.waitForLoadState('networkidle');

	// メールアドレスのコピー用ボタンが存在することを確認 (テキストに @ が含まれているはず)
	const emailButton = page.locator('main button:has-text("@")');
	await expect(emailButton).toBeVisible();
	await expect(page.locator('details')).toHaveCount(0);
	await expect(page.getByText('実績を表示', { exact: true })).toHaveCount(0);
	await expect(page.getByText('専門領域を表示', { exact: true })).toHaveCount(0);

	// スクリーンショットを撮影
	await page.screenshot({ path: 'static/screenshots/profile.png', fullPage: true });
});
