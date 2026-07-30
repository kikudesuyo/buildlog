import { test, expect } from '@playwright/test';

test('verify rss feed generation', async ({ request }) => {
	const response = await request.get('http://localhost:5173/rss.xml');
	expect(response.ok()).toBeTruthy();

	const contentType = response.headers()['content-type'];
	expect(contentType).toContain('application/xml');

	const text = await response.text();
	expect(text).toContain('<?xml version="1.0" encoding="UTF-8" ?>');
	expect(text).toContain('<rss version="2.0"');
	expect(text).toContain('<channel>');
	expect(text).toContain('<title>Buildlog</title>');
	
	// DBに初期データが存在するため、少なくとも一つの <item> が含まれているはず
	expect(text).toContain('<item>');
});
