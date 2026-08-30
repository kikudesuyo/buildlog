<script lang="ts">
	import { createLearningSummary } from '$lib/api/client';
	import type { Learning, LearningPeriodType } from '$lib/api/types';

	let { data } = $props();
	let activePeriod = $state<LearningPeriodType>('daily');
	let learnings = $state<Record<LearningPeriodType, Learning[]>>({
		daily: data.daily,
		weekly: data.weekly,
		monthly: data.monthly
	});
	let errorMessage = $state('');
	let isSubmitting = $state(false);

	const labels = { daily: '今日', weekly: '今週', monthly: '今月' } as const;

	async function generate() {
		if (activePeriod === 'daily') return;
		isSubmitting = true;
		errorMessage = '';
		try {
			const item = await createLearningSummary(activePeriod);
			learnings[activePeriod] = [item, ...learnings[activePeriod]];
		} catch {
			errorMessage = '元となる学びがないため、まとめを生成できませんでした。';
		} finally {
			isSubmitting = false;
		}
	}
</script>

<svelte:head><title>成長ログ — Buildlog</title></svelte:head>

<section class="mx-auto max-w-4xl px-gutter" aria-labelledby="learning-heading">
	<header class="mb-8 flex flex-wrap items-end justify-between gap-4">
		<div>
			<p class="font-label-sm text-label-sm mb-2 tracking-widest text-outline uppercase">Growth Log</p>
			<h1 id="learning-heading" class="font-headline-lg text-headline-lg text-primary">エンジニア成長ログ</h1>
			<p class="font-body-md text-body-md mt-2 text-on-surface-variant">日々の学びを残し、週・月単位で振り返ります。</p>
		</div>
		{#if activePeriod !== 'daily'}
			<button type="button" onclick={generate} disabled={isSubmitting} class="font-label-md text-label-md min-h-11 rounded-lg bg-primary px-4 py-2 text-on-primary">{isSubmitting ? '生成中…' : `${labels[activePeriod]}をまとめる`}</button>
		{/if}
	</header>

	<div class="mb-6 flex gap-2 border-b border-outline-variant/30" role="tablist">
		{#each Object.entries(labels) as [period, label] (period)}
			<button type="button" role="tab" aria-selected={activePeriod === period} onclick={() => (activePeriod = period as LearningPeriodType)} class="font-label-md text-label-md border-b-2 px-4 py-3 {activePeriod === period ? 'border-primary text-primary' : 'border-transparent text-outline'}">{label}</button>
		{/each}
	</div>

	{#if activePeriod === 'daily'}
		<div class="mb-8 rounded-xl border border-outline-variant/30 bg-surface-container-lowest p-5">
			<p class="font-label-md text-label-md text-primary">Daily は自動取り込みです</p>
			<p class="font-body-sm text-body-sm mt-2 text-on-surface-variant">Notion に記録したログをローカルスケジューラが取り込みます。この画面では取り込まれた今日の学びを確認できます。</p>
		</div>
	{/if}

	{#if errorMessage}<p class="font-body-sm text-body-sm mb-4 text-error" role="alert">{errorMessage}</p>{/if}
	<div class="space-y-4">
		{#each learnings[activePeriod] as item (item.id)}
			<article class="rounded-xl border border-outline-variant/30 bg-surface-container-lowest p-5">
				<div class="mb-3 flex flex-wrap items-center justify-between gap-2"><span class="font-label-sm text-label-sm text-outline">{item.periodStart} — {item.periodEnd}</span>{#if item.level}<span class="font-label-sm text-label-sm rounded-full bg-primary/10 px-2 py-1 text-primary">{item.level}</span>{/if}</div>
				<p class="font-body-md text-body-md whitespace-pre-wrap text-primary">{item.content}</p>
			</article>
		{:else}
			<p class="font-body-md text-body-md rounded-xl border border-dashed border-outline-variant/40 p-6 text-on-surface-variant">まだ記録がありません。</p>
		{/each}
	</div>
</section>
