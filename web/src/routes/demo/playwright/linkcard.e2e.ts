import { test, expect } from '@playwright/test';

test('verify link card rendering and OGP fetching', async ({ page }) => {
	// テスト用の記事を作成するため、管理者APIを直接呼び出して作成する
	const uniqueTitle = 'E2Eテスト_リンクカード_' + Date.now();
	
	// APIを叩いて、単独URLを含む記事を投稿
	const apiContext = page.request;
	const createRes = await apiContext.post('http://localhost:8081/api/v1/techs', {
		data: {
			title: uniqueTitle,
			content: 'この記事には外部リンクが含まれます。\n\nhttps://github.com/kikudesuyo\n\n確認してください。',
			status: 'published'
		}
	});
	expect(createRes.ok()).toBeTruthy();
	const createResult = await createRes.json();
	const techId = createResult.data.id;

	// 作成した記事のページへ移動
	await page.goto(`http://localhost:5173/tech/${techId}`);
	await page.waitForLoadState('networkidle');

	// フッター等のリンクと競合しないよう、プレースホルダー内の LinkCard を特定して待機
	const linkCard = page.locator('.link-card-placeholder a[href="https://github.com/kikudesuyo"]').first();
	await linkCard.waitFor({ state: 'visible', timeout: 8000 });

	// ブログカードの内容を検証
	await expect(linkCard).toBeVisible();
	
	// github という文字がいずれかの要素（タイトルやサイト名など）に含まれているはず
	await expect(linkCard.locator('span:has-text("github")').first()).toBeVisible();

	// スクリーンショット撮影
	await page.screenshot({ path: 'static/screenshots/linkcard-verification.png', fullPage: true });

	// 後片付けとして削除
	const deleteRes = await apiContext.delete(`http://localhost:8081/api/v1/techs/${techId}`);
	expect(deleteRes.ok()).toBeTruthy();
});
