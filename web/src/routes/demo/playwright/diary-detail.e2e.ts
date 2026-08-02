import { expect, test } from '@playwright/test';

test('opens a diary entry and shows its comments', async ({ page }) => {
	await page.goto('http://localhost:5173/');

	const diaryLink = page.locator('a[href^="/diary/"]').first();
	await expect(diaryLink).toBeVisible();
	await diaryLink.click();

	await expect(page).toHaveURL(/\/diary\/\d+$/);
	await expect(page.getByRole('heading', { name: 'コメント' })).toBeVisible();
});
