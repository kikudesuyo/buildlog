<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { syncQiitaArticles } from '$lib/api/client';
	import TechFeed from '$lib/components/TechFeed.svelte';
	let { data } = $props();
	let isSyncing = $state(false);
	let syncMessage = $state('');
	let syncError = $state('');
	let feedVersion = $state(0);

	async function handleQiitaSync() {
		isSyncing = true;
		syncMessage = '';
		syncError = '';
		try {
			const synced = await syncQiitaArticles();
			syncMessage = `${synced}件のQiita記事を同期しました。`;
			await invalidateAll();
			feedVersion += 1;
		} catch {
			syncError = 'Qiita記事の同期に失敗しました。';
		} finally {
			isSyncing = false;
		}
	}
</script>

<svelte:head><title>Buildlog — Admin Tech Feed</title></svelte:head>
<div class="editorial-container mx-auto flex flex-col gap-6 px-gutter">
	<section class="flex justify-end">
		<button type="button" onclick={handleQiitaSync} disabled={isSyncing} class="font-label-md text-label-md flex min-h-11 shrink-0 items-center gap-2 rounded-lg bg-primary px-5 py-2.5 text-on-primary transition-colors hover:bg-primary/95 disabled:cursor-wait disabled:opacity-60">
			<span class="material-symbols-outlined text-[18px]">{isSyncing ? 'sync' : 'sync_alt'}</span>
			{isSyncing ? '同期中…' : 'Qiitaから取り込む'}
		</button>
	</section>
	{#if syncMessage}<p class="font-body-sm text-body-sm rounded-lg border border-primary/20 bg-primary/5 px-4 py-3 text-primary" role="status">{syncMessage}</p>{/if}
	{#if syncError}<p class="font-body-sm text-body-sm rounded-lg border border-error/30 bg-error-container/30 px-4 py-3 text-error" role="alert">{syncError}</p>{/if}
	{#key feedVersion}
		<TechFeed
			featuredArticle={data.featuredArticle}
			techArticles={data.techArticles}
			isAdmin
		/>
	{/key}
</div>
