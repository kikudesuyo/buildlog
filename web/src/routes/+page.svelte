<script lang="ts">
	import DiaryFeed from '$lib/components/DiaryFeed.svelte';
	import { resolve } from '$app/paths';
	let { data } = $props();
</script>

<svelte:head><title>Buildlog — Diary</title></svelte:head>

<div class="mx-auto max-w-container-max px-gutter">
	{#if data.selectedTag}
		<div class="mb-6 flex w-fit items-center gap-3 rounded-lg border border-outline-variant/30 bg-surface-container-high px-4 py-2.5">
			<span class="text-body-sm font-medium text-on-surface-variant">タグフィルタ:</span>
			<span class="font-label-sm text-label-sm rounded-full bg-primary px-2.5 py-1 font-bold text-on-primary">{data.selectedTag}</span>
			<a href={resolve('/')} class="ml-1 cursor-pointer text-[18px] font-bold text-outline transition-colors hover:text-primary" title="フィルタ解除">×</a>
		</div>
	{/if}
	<DiaryFeed entries={data.diaryEntries} />
</div>
