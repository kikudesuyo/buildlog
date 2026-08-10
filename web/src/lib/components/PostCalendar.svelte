<script lang="ts">
import type { HistoryItem } from '$lib/api/types';
import { resolve } from '$app/paths';
import { SvelteMap } from 'svelte/reactivity';

	type Props = {
		history: HistoryItem[];
	};

	let { history }: Props = $props();

	// 履歴データを日付（YYYY-MM-DD）ごとにグループ化するマップ
	const historyMap = $derived.by(() => {
		const m = new SvelteMap<string, HistoryItem[]>();
		for (const item of history) {
			if (!item.createdAt) continue;
			const dateStr = item.createdAt.substring(0, 10); // YYYY-MM-DD
			if (!m.has(dateStr)) {
				m.set(dateStr, []);
			}
			m.get(dateStr)!.push(item);
		}
		return m;
	});

	// ポップオーバーの表示ステート
	let selectedDate = $state<string | null>(null);
	let selectedPosts = $derived(selectedDate ? historyMap.get(selectedDate) || [] : []);
	let popoverX = $state(0);
	let popoverY = $state(0);
	let isPopoverOpen = $state(false);

	// 日付クリック時のポップオーバー処理
	function handleDateClick(event: MouseEvent, dateStr: string) {
		const posts = historyMap.get(dateStr) || [];
		if (posts.length === 0) {
			isPopoverOpen = false;
			return;
		}

		selectedDate = dateStr;
		isPopoverOpen = true;

		// クリックした要素の近くにポップオーバーを配置
		const rect = (event.currentTarget as HTMLElement).getBoundingClientRect();
		const scrollX = window.scrollX;
		const scrollY = window.scrollY;

		popoverX = rect.left + scrollX + rect.width / 2;
		popoverY = rect.bottom + scrollY + 8;
	}

	function closePopover() {
		isPopoverOpen = false;
	}

	// 月間カレンダーのデータ生成
	let currentYear = $state(new Date().getFullYear());
	let currentMonth = $state(new Date().getMonth()); // 0-indexed

	function nextMonth() {
		if (currentMonth === 11) {
			currentMonth = 0;
			currentYear += 1;
		} else {
			currentMonth += 1;
		}
		closePopover();
	}

	function prevMonth() {
		if (currentMonth === 0) {
			currentMonth = 11;
			currentYear -= 1;
		} else {
			currentMonth -= 1;
		}
		closePopover();
	}

	const monthlyDays = $derived.by(() => {
		// 対象月の最初の日
		const firstDay = new Date(currentYear, currentMonth, 1);
		const startDayOfWeek = firstDay.getDay(); // 曜日 (0: 日曜日, 6: 土曜日)

		// 対象月の日数
		const lastDay = new Date(currentYear, currentMonth + 1, 0);
		const totalDays = lastDay.getDate();

		const days: { dateStr: string | null; dayNum: number | null; hasPost: boolean }[] = [];

		// 月最初の空白埋め
		for (let i = 0; i < startDayOfWeek; i++) {
			days.push({ dateStr: null, dayNum: null, hasPost: false });
		}

		// 日にちの追加
		for (let d = 1; d <= totalDays; d++) {
			const tempDate = new Date(currentYear, currentMonth, d);
			// タイムゾーンによるずれを防ぐためローカル日付を取得して整形
			const y = tempDate.getFullYear();
			const m = String(tempDate.getMonth() + 1).padStart(2, '0');
			const dayStr = String(tempDate.getDate()).padStart(2, '0');
			const dateStr = `${y}-${m}-${dayStr}`;

			const posts = historyMap.get(dateStr) || [];
			days.push({
				dateStr,
				dayNum: d,
				hasPost: posts.length > 0
			});
		}

		return days;
	});

	const monthNames = [
		'1月', '2月', '3月', '4月', '5月', '6月',
		'7月', '8月', '9月', '10月', '11月', '12月'
	];
</script>

<div id="post-history" class="flex flex-col gap-10 mt-10 p-6 rounded-2xl border border-outline-variant/30 bg-surface-container-lowest">
	<!-- ヘッダー -->
	<header class="flex flex-col gap-2">
		<h2 class="font-headline-md text-headline-md font-bold text-primary flex items-center gap-2">
			<span class="material-symbols-outlined text-2xl">calendar_month</span>
			投稿履歴とカレンダー
		</h2>
		<p class="font-body-sm text-body-sm text-outline">
			これまでの執筆活動の記録です。緑色のセルやドットマークのある日をクリックすると、その日の記事にアクセスできます。
		</p>
	</header>

	<section class="flex flex-col gap-4">
			<header class="flex items-center justify-between">
				<h3 class="font-label-md text-label-md tracking-wider text-outline uppercase">
					カレンダー / Calendar
				</h3>
				<div class="flex items-center gap-1">
					<button
						type="button"
						onclick={prevMonth}
						class="p-1 rounded-md text-outline hover:bg-surface-container hover:text-primary transition-all cursor-pointer flex items-center justify-center"
					>
						<span class="material-symbols-outlined text-lg">chevron_left</span>
					</button>
					<span class="font-label-md text-label-md font-semibold text-primary px-2">
						{currentYear}年 {monthNames[currentMonth]}
					</span>
					<button
						type="button"
						onclick={nextMonth}
						class="p-1 rounded-md text-outline hover:bg-surface-container hover:text-primary transition-all cursor-pointer flex items-center justify-center"
					>
						<span class="material-symbols-outlined text-lg">chevron_right</span>
					</button>
				</div>
			</header>

			<!-- カレンダーグリッド -->
			<div class="grid grid-cols-7 gap-y-2 text-center text-body-sm font-medium">
				<!-- 曜日ヘッダー -->
				<span class="text-error/85 py-1 text-[11px]">日</span>
				<span class="text-outline py-1 text-[11px]">月</span>
				<span class="text-outline py-1 text-[11px]">火</span>
				<span class="text-outline py-1 text-[11px]">水</span>
				<span class="text-outline py-1 text-[11px]">木</span>
				<span class="text-outline py-1 text-[11px]">金</span>
				<span class="text-primary py-1 text-[11px]">土</span>

				<!-- 日にちセル -->
				{#each monthlyDays as day, index (day.dateStr ?? index)}
					{#if day.dayNum}
						<!-- svelte-ignore a11y_click_events_have_key_events -->
						<!-- svelte-ignore a11y_no_static_element_interactions -->
						<div
							onclick={(e) => day.dateStr && handleDateClick(e, day.dateStr)}
							class="relative flex flex-col items-center justify-center h-8 w-8 mx-auto rounded-full transition-all duration-150 {day.hasPost ? 'bg-primary/10 hover:bg-primary/20 text-primary font-bold cursor-pointer' : 'text-on-surface-variant'}"
						>
							<span>{day.dayNum}</span>
							{#if day.hasPost}
								<span class="absolute bottom-1 h-1 w-1 rounded-full bg-primary"></span>
							{/if}
						</div>
					{:else}
						<div class="h-8 w-8"></div>
					{/if}
				{/each}
			</div>
	</section>
</div>

<!-- 記事リストポップオーバー -->
{#if isPopoverOpen && selectedDate}
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		onclick={closePopover}
		class="fixed inset-0 z-40 bg-transparent"
	></div>

	<div
		class="absolute z-50 min-w-[280px] max-w-[340px] rounded-xl border border-outline-variant/30 bg-surface-container-lowest p-4 shadow-xl -translate-x-1/2 flex flex-col gap-3"
		style="left: {popoverX}px; top: {popoverY}px;"
	>
		<header class="flex items-center justify-between border-b border-outline-variant/20 pb-2">
			<span class="font-label-md text-label-md font-bold text-primary">
				{selectedDate} の投稿
			</span>
			<button
				type="button"
				onclick={closePopover}
				class="text-outline hover:text-primary transition-colors text-lg font-bold"
			>
				×
			</button>
		</header>
		<ul class="flex flex-col gap-2.5 max-h-[220px] overflow-y-auto pr-1">
			{#each selectedPosts as post (post.id)}
				<li class="flex flex-col gap-1 text-body-md">
					<a
						href={post.type === 'tech' ? resolve(`/tech/${post.id}?from=calendar`) : resolve(`/diary/${post.id}?from=calendar`)}
						class="text-on-surface hover:text-primary hover:underline transition-colors font-medium leading-snug"
						onclick={closePopover}
					>
						{post.title}
					</a>
					<div class="flex items-center gap-2">
						<span
							class="text-[10px] font-bold px-1.5 py-0.5 rounded uppercase tracking-wider {post.type === 'tech' ? 'bg-primary-fixed text-primary' : 'bg-secondary-fixed text-secondary'}"
						>
							{post.type}
						</span>
					</div>
				</li>
			{/each}
		</ul>
	</div>
{/if}
