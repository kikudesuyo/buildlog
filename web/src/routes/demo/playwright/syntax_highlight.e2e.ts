import { test, expect } from '@playwright/test';

test('verify syntax highlighting is applied to code blocks in article view', async ({ page }) => {
	// API を使ってテスト用の技術記事を投稿
	const response = await page.request.post('http://localhost:8081/api/v1/techs', {
		data: {
			title: 'Test Code Highlight',
			content: '以下はJavaScriptのテストコードです。\n\n```javascript\nconst hello = "world";\nconsole.log(hello);\n```',
			status: 'published'
		}
	});
	expect(response.ok()).toBeTruthy();
	const result = await response.json();
	const postId = result.data.id;

	// 作成された記事の詳細画面へ移動
	await page.goto(`http://localhost:5173/tech/${postId}`);
	await page.waitForLoadState('networkidle');

	// コードブロックの存在を確認
	const codeBlock = page.locator('pre code');
	await expect(codeBlock).toBeVisible();

	// highlight.js によって適用されたクラス (hljs または language-) を確認
	const classAttr = await codeBlock.getAttribute('class');
	expect(classAttr).toContain('hljs');
	expect(classAttr).toContain('language-javascript');

	// 変化後のスクリーンショットを撮影
	await page.screenshot({ path: 'static/screenshots/after_highlight.png', fullPage: true });
});
