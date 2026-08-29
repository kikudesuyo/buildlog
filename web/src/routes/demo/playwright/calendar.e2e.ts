import { test, expect } from '@playwright/test';

test('verify calendar UI displays correctly and interactive click actions work', async ({ page }) => {
	// プロフィール画面へ移動
	await page.goto('http://localhost:5173/profile');
	await page.waitForLoadState('networkidle');

	// 見出しの確認
	const header = page.getByRole('heading', { name: '投稿履歴', exact: true });
	await expect(header).toBeVisible();
	await expect(page.getByText('これまでの執筆活動の記録です。投稿のある日をクリックすると、その記事にアクセスできます。')).toBeVisible();

	// カレンダー内の投稿日を示すドットマークや primary カラーのセルがあるかを確認
	// 投稿が存在する日付セルをクリックしてみる
	const activeCell = page.locator('.relative.h-8.w-8.mx-auto.rounded-full.text-primary').first();
	
	// もし投稿がある日付が存在する場合
	if (await activeCell.count() > 0) {
		await activeCell.click();
		
		// ポップオーバーが表示されていることを確認
		const popoverHeader = page.locator('text=の投稿');
		await expect(popoverHeader).toBeVisible();
	}

	// 変化後のスクリーンショットを撮影
	await page.screenshot({ path: 'static/screenshots/after_profile.png', fullPage: true });
});
