<script lang="ts">
	import { restoreEntry } from '$lib/api/client';

	let { data } = $props();
	const initialData = data;
	let trashEntries = $state(initialData.trashEntries);

	function formatDate(dateStr: string) {
		if (!dateStr) return '';
		const date = new Date(dateStr);
		return `${date.getFullYear()}.${String(date.getMonth() + 1).padStart(2, '0')}.${String(date.getDate()).padStart(2, '0')}`;
	}

	async function handleRestore(id: number) {
		try {
			await restoreEntry(id);
			trashEntries = trashEntries.filter(entry => entry.id !== id);
			alert('記事を復元しました。');
		} catch {
			alert('復元に失敗しました。');
		}
	}
</script>

<svelte:head>
	<title>アーカイブ — Buildlog Admin</title>
</svelte:head>

<div class="editorial-container mx-auto px-gutter pt-8">
	<!-- ヘッダー -->
	<header class="mb-8">
		<h1 class="font-display-lg text-display-lg text-primary font-bold tracking-tight">アーカイブ (削除済み一覧)</h1>
		<p class="font-body-md text-body-md text-on-surface-variant mt-1">削除されたつぶやきや技術記事を管理・復元できます。</p>
	</header>

	<!-- ゴミ箱一覧 -->
	<div class="flex flex-col gap-4">
		{#if trashEntries.length === 0}
			<div class="text-center py-12 border border-dashed border-outline-variant/30 rounded-2xl bg-surface-container-low text-outline">
				アーカイブは空です。
			</div>
		{:else}
			<div class="grid grid-cols-1 gap-4">
				{#each trashEntries as entry (entry.id)}
					<div class="bg-surface-container-lowest border border-outline-variant/20 rounded-2xl p-6 flex flex-col md:flex-row md:items-center justify-between gap-4 shadow-xs relative group transition-all duration-300 hover:shadow-md hover:border-primary/30">
						<div class="flex-grow flex flex-col gap-1.5 pr-8">
							<div class="flex items-center gap-2 mb-1">
								<span class="font-label-sm text-label-sm px-2 py-0.5 rounded uppercase font-semibold
									{entry.type === 'diary' ? 'bg-amber-100 text-amber-800' : 'bg-blue-100 text-blue-800'}">
									{entry.type === 'diary' ? 'つぶやき' : '技術記事'}
								</span>
								{#if entry.category}
									<span class="font-label-sm text-label-sm text-outline-variant">|</span>
									<span class="font-label-sm text-label-sm text-outline">{entry.category}</span>
								{/if}
								<span class="text-outline-variant">•</span>
								<span class="font-label-sm text-label-sm text-outline">削除日: {formatDate(entry.deletedAt)}</span>
							</div>

							<h3 class="font-headline-sm text-headline-sm text-primary font-bold">{entry.title}</h3>
							<p class="font-body-md text-body-md text-on-surface-variant line-clamp-2 mt-1">{entry.content}</p>
						</div>

						<!-- 復元ボタン -->
						<div class="flex items-center gap-2 self-end md:self-center">
							<button
								type="button"
								onclick={() => handleRestore(entry.id)}
								class="bg-surface-container-high hover:bg-primary hover:text-on-primary text-primary font-label-md text-label-md px-4 py-2 rounded-lg font-medium transition-all duration-200 flex items-center gap-1.5 cursor-pointer"
							>
								<span class="material-symbols-outlined text-[18px]">restore</span>
								復元する
							</button>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</div>
