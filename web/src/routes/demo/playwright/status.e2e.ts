import { test, expect } from '@playwright/test';

test('verify draft status logic and capture screenshots', async ({ page }) => {
	const uniqueTitle = 'E2Eテスト下書き_' + Date.now();

	await page.goto('http://localhost:5173/');
	await page.waitForLoadState('networkidle');

	const draftBadgeOnPublic = page.locator('text=下書き');
	await expect(draftBadgeOnPublic).not.toBeVisible();

	await page.goto('http://localhost:5173/admin');
	await page.waitForLoadState('networkidle');

	await page.click('text=つぶやく');
	await page.waitForURL('**/admin/diary/new');

	await page.fill('input[placeholder="タイトルを入力..."]', uniqueTitle);
	await page.fill('textarea[placeholder="物語を書き始めましょう..."]', 'これは下書き保存のテストです。一般画面には表示されません。');

	await page.click('text=下書き保存');
	await page.waitForURL('**/admin');

	const adminDraftPost = page.locator('article', { hasText: uniqueTitle }).first();
	await expect(adminDraftPost).toBeVisible();
	await expect(adminDraftPost.locator('span:text-is("下書き")')).toBeVisible();

	await page.screenshot({ path: 'static/screenshots/admin-draft.png', fullPage: true });

	await page.goto('http://localhost:5173/');
	await page.waitForLoadState('networkidle');

	const publicDraftPost = page.locator('article', { hasText: uniqueTitle });
	await expect(publicDraftPost).not.toBeVisible();
});
