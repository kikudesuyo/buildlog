import { test, expect } from '@playwright/test';

test('capture trash and soft delete screenshots', async ({ page }) => {
	page.on('dialog', async dialog => {
		await dialog.accept();
	});

	await page.goto('http://localhost:5173/admin/tech');
	await page.waitForLoadState('networkidle');

	const articleTitleEl = page.locator('h3').first();
	const articleTitle = await articleTitleEl.innerText();

	const deleteButton = page.locator('button[title="削除"]').first();
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
