<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { syncQiitaArticles } from '$lib/api/client';
	import TechFeed from '$lib/components/TechFeed.svelte';
	let { data } = $props();
	let isSyncing = $state(false);
	let syncMessage = $state('');
	let syncError = $state('');

	async function handleSync() {
		if (isSyncing) return;
		isSyncing = true;
		syncMessage = '';
		syncError = '';
		try {
			const count = await syncQiitaArticles();
			syncMessage = `Qiita記事を${count}件同期しました。`;
			await invalidateAll();
		} catch {
			syncError = 'Qiita記事の同期に失敗しました。';
		} finally {
			isSyncing = false;
		}
	}
</script>

<svelte:head><title>Buildlog — Admin Tech Feed</title></svelte:head>

<div class="editorial-container mx-auto mb-8 flex flex-col gap-3 px-gutter">
		<button type="button" onclick={handleSync} disabled={isSyncing} class="font-label-md text-label-md min-h-11 self-start rounded-lg bg-primary px-5 py-2.5 text-on-primary transition-colors hover:bg-primary/90 disabled:cursor-wait disabled:opacity-60">
			{isSyncing ? 'Qiitaから同期中…' : 'Qiitaから取り込む'}
		</button>
		<p class="font-body-sm text-body-sm text-on-surface-variant" role="status" aria-live="polite">{syncMessage}</p>
		{#if syncError}<p class="font-body-sm text-body-sm text-error" role="alert">{syncError}</p>{/if}
</div>

<TechFeed {...data} isAdmin />
