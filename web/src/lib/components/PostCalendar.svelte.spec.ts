import { page } from 'vitest/browser';
import { describe, expect, it } from 'vitest';
import type { HistoryItem } from '$lib/api/types';
import { render } from 'vitest-browser-svelte';
import PostCalendar from './PostCalendar.svelte';

const date = new Date();
const dateKey = date.getFullYear() + '-' + String(date.getMonth() + 1).padStart(2, '0') + '-15';
const history: HistoryItem[] = [
	{ id: 10, type: 'diary', title: '今日の日記', createdAt: dateKey + 'T09:00:00.000Z' },
	{ id: 11, type: 'tech', title: '今日のTech記事', createdAt: dateKey + 'T10:00:00.000Z' }
];

describe('PostCalendar.svelte', () => {
	it('カレンダーと投稿履歴を表示する', async () => {
		render(PostCalendar, { history });

		await expect.element(page.getByRole('heading', { name: '投稿履歴' })).toBeInTheDocument();
		await expect.element(page.getByText(date.getFullYear() + '年 ' + (date.getMonth() + 1) + '月')).toBeInTheDocument();
	});

	it('投稿のある日を選ぶと、その日の投稿を表示する', async () => {
		render(PostCalendar, { history });

		await page.getByText('15', { exact: true }).click();
		await expect.element(page.getByText(dateKey + ' の投稿')).toBeInTheDocument();
		await expect.element(page.getByText('今日の日記')).toBeInTheDocument();
		await expect.element(page.getByText('今日のTech記事')).toBeInTheDocument();
	});

	it('月移動ボタンで表示月を変更する', async () => {
		render(PostCalendar, { history });
		await page.getByRole('button', { name: '次の月' }).click();

		const next = new Date(date.getFullYear(), date.getMonth() + 1, 1);
		await expect.element(page.getByText(next.getFullYear() + '年 ' + (next.getMonth() + 1) + '月')).toBeInTheDocument();
	});
});
