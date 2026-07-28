import { test, expect } from '@playwright/test';

test('capture trash and soft delete screenshots', async ({ page }) => {
	page.on('dialog', async dialog => {
		await dialog.accept();
	});

	await page.goto('http://localhost:5173/admin/tech');
	await page.waitForLoadState('networkidle');

	// 最初の記事要素を取得
	const firstArticle = page.locator('article').first();
	
	// その記事要素の中の見出し (h2 または h3) を取得
	const articleTitleEl = firstArticle.locator('h2, h3').first();
	const articleTitle = (await articleTitleEl.innerText()).trim();

	// その記事要素の中の削除ボタンをクリック
	const deleteButton = firstArticle.locator('button[title="削除"]').first();
	await deleteButton.click();

	await page.waitForTimeout(500);

	await page.goto('http://localhost:5173/admin/trash');
	await page.waitForLoadState('networkidle');

	await expect(page.locator(`h3:has-text("${articleTitle}")`)).toBeVisible();

	await page.screenshot({ path: 'static/screenshots/admin-trash.png', fullPage: true });

	const restoreButton = page.locator('button:has-text("復元する")').first();
	await restoreButton.click();

	await page.waitForTimeout(500);

	await expect(page.locator(`h3:has-text("${articleTitle}")`)).not.toBeVisible();
});
