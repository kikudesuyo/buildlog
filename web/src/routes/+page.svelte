<script lang="ts">
	import DiaryFeed from '$lib/components/DiaryFeed.svelte';
	import { resolve } from '$app/paths';
	let { data } = $props();
</script>

<svelte:head><title>Buildlog — Diary</title></svelte:head>

<div class="mx-auto max-w-container-max px-gutter">
	{#if data.selectedTag}
		<div class="mb-6 flex items-center gap-3 bg-surface-container-high px-4 py-2.5 rounded-lg border border-outline-variant/30 w-fit">
			<span class="text-body-sm font-medium text-on-surface-variant">タグフィルタ:</span>
			<span class="font-label-sm text-label-sm px-2.5 py-1 rounded-full bg-primary text-on-primary font-bold">
				{data.selectedTag}
			</span>
			<a href={resolve('/')} class="text-outline hover:text-primary transition-colors text-[18px] font-bold ml-1 cursor-pointer" title="フィルタ解除">
				×
			</a>
		</div>
	{/if}

	<DiaryFeed entries={data.diaryEntries} />
</div>
