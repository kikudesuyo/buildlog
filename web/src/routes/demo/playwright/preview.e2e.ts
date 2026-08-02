import { test, expect } from '@playwright/test';

test('verify tabbed Markdown preview in tech editor and capture screenshots', async ({ page }) => {
	await page.goto('http://localhost:5173/admin/tech/new');
	await page.waitForLoadState('networkidle');

	// 1. タイトルと Markdown 本文の入力
	await page.fill('input[placeholder="タイトルを入力..."]', 'Markdownプレビューテスト');
	await page.fill('textarea[placeholder="本文を書き始めましょう..."]', '# テスト大見出し\n\nこれはプレビュー機能の検証です。\n\n- リスト項目A\n- リスト項目B\n\n`code block`');

	// 2. プレビュータブをクリック
	await page.getByRole('tab', { name: 'プレビュー' }).click();

	// プレビュー表示の検証
	const previewArea = page.locator('text=リアルタイムプレビュー');
	await expect(previewArea).toBeVisible();
	await expect(page.locator('textarea[placeholder="本文を書き始めましょう..."]')).toBeHidden();

	// パースされた HTML 要素を確認
	const parsedHeader = page.locator('h1:text-is("テスト大見出し")');
	await expect(parsedHeader).toBeVisible();
	
	const parsedListItem = page.locator('li:text-is("リスト項目A")');
	await expect(parsedListItem).toBeVisible();

	// 3. 編集タブへ戻り、プレビューが非表示になることを確認
	await page.getByRole('tab', { name: '編集' }).click();
	await expect(page.locator('textarea[placeholder="本文を書き始めましょう..."]')).toBeVisible();
	await expect(previewArea).toBeHidden();

	// スクリーンショット撮影
	await page.screenshot({ path: 'static/screenshots/markdown-preview.png', fullPage: true });
});
