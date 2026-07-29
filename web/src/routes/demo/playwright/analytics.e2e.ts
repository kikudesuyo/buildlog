import { test, expect } from '@playwright/test';

test('verify admin analytics dashboard UI elements load correctly', async ({ page }) => {
	// 管理画面ダッシュボードへ移動
	await page.goto('http://localhost:5173/admin/analytics');
	await page.waitForLoadState('networkidle');

	// ヘッダーの確認
	const header = page.locator('text=Analytics Dashboard');
	await expect(header).toBeVisible();

	// サマリーカードの確認
	const viewsCard = page.locator('text=総閲覧数');
	const likesCard = page.locator('text=総いいね数');
	const postsCard = page.locator('text=総コンテンツ数');
	await expect(viewsCard).toBeVisible();
	await expect(likesCard).toBeVisible();
	await expect(postsCard).toBeVisible();

	// グラフの存在を確認
	const graphTitle = page.locator('text=過去12ヶ月の活動推移 / Monthly Activity');
	await expect(graphTitle).toBeVisible();

	// ランキングセクションの確認
	const rankingTitle = page.locator('text=人気コンテンツ / Content Ranking');
	await expect(rankingTitle).toBeVisible();

	// いいね数タブへの切り替えテスト
	const likesTab = page.locator('button:has-text("いいね数")');
	await expect(likesTab).toBeVisible();
	await likesTab.click();

	// 変化後のスクリーンショットを撮影
	await page.screenshot({ path: 'static/screenshots/after_admin.png', fullPage: true });
});
