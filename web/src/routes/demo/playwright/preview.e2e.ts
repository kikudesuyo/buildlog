import { test, expect } from '@playwright/test';

test('verify real-time Markdown preview in tech editor and capture screenshots', async ({ page }) => {
	await page.goto('http://localhost:5173/admin/tech/new');
	await page.waitForLoadState('networkidle');

	// 1. タイトルと Markdown 本文の入力
	await page.fill('input[placeholder="タイトルを入力..."]', 'Markdownプレビューテスト');
	await page.fill('textarea[placeholder="本文を書き始めましょう..."]', '# テスト大見出し\n\nこれはプレビュー機能の検証です。\n\n- リスト項目A\n- リスト項目B\n\n`code block`');

	// 2. プレビュータブをクリック
	await page.click('text=👁️ プレビュー');

	// プレビュー表示の検証
	const previewArea = page.locator('text=リアルタイムプレビュー');
	await expect(previewArea).toBeVisible();

	// パースされた HTML 要素を確認
	const parsedHeader = page.locator('h1:text-is("テスト大見出し")');
	await expect(parsedHeader).toBeVisible();
	
	const parsedListItem = page.locator('li:text-is("リスト項目A")');
	await expect(parsedListItem).toBeVisible();

	// 3. 2カラム（分割表示）をクリック
	await page.click('text=📖 2カラム');

	// エディタとプレビューの両方が見えることを確認
	const editorTextarea = page.locator('textarea[placeholder="本文を書き始めましょう..."]');
	await expect(editorTextarea).toBeVisible();
	await expect(parsedHeader).toBeVisible();

	// スクリーンショット撮影
	await page.screenshot({ path: 'static/screenshots/markdown-preview.png', fullPage: true });
});
