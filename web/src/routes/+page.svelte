<script lang="ts">
	import { goto } from '$app/navigation';
	import { deleteDiary } from '$lib/api/client';
	import type { DiaryEntry } from '$lib/api/types';
	import { page } from '$app/stores';

	let { data } = $props();

	// 管理者かどうか判定
	let isAdmin = $derived($page.url.pathname.startsWith('/admin'));

	let diariesList = $state<DiaryEntry[]>([]);

	// Svelte 5 の reactive な data 追跡警告対策
	$effect(() => {
		diariesList = [...data.diaryEntries];
	});

	let visibleCount = $state(3);
	let displayEntries = $derived(diariesList.slice(0, visibleCount));

	function loadMore() {
		visibleCount += 2;
	}

	function formatDate(dateStr: string) {
		if (!dateStr) return '';
		const date = new Date(dateStr);
		return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`;
	}

	async function handleDelete(id: number) {
		if (!confirm('この記事を削除してもよろしいですか？')) return;
		try {
			await deleteDiary(id);
			diariesList = diariesList.filter((d) => d.id !== id);
		} catch (err) {
			alert('削除に失敗しました。');
		}
	}
</script>

<svelte:head>
	<title>Essence — Diary</title>
</svelte:head>

<div class="editorial-container mx-auto px-gutter relative flex flex-col gap-8">
	<!-- Intro / Header -->
	<header class="flex items-center justify-between">
		<h1 class="font-display-lg text-display-lg text-primary">日々のつぶやき</h1>
		{#if isAdmin}
			<button
				type="button"
				onclick={() => goto('/diary/new')}
				class="font-label-md text-label-md cursor-pointer rounded-lg bg-primary px-6 py-2.5 text-on-primary transition-all hover:bg-primary/95 active:scale-95 flex items-center gap-1.5"
			>
				<span class="material-symbols-outlined text-[18px]">add</span>
				つぶやく
			</button>
		{/if}
	</header>

	<!-- つぶやき一覧 -->
	<div class="flex flex-col gap-6">
		{#each displayEntries as entry (entry.id)}
			<article class="group relative rounded-xl border border-transparent p-4 -mx-4">
				<div class="mb-stack-sm flex items-center justify-between">
					<span class="font-label-sm text-label-sm text-outline">{formatDate(entry.createdAt)}</span>
					
					<!-- アクションボタン -->
					{#if isAdmin}
						<div class="flex gap-2">
							<button
								type="button"
								onclick={() => goto(`/diary/${entry.id}/edit`)}
								class="p-1 cursor-pointer text-outline opacity-60 hover:opacity-100 hover:text-primary transition-all duration-200"
								title="編集"
							>
								<span class="material-symbols-outlined text-[18px]">edit</span>
							</button>
							<button
								type="button"
								onclick={() => handleDelete(entry.id)}
								class="p-1 cursor-pointer text-outline opacity-60 hover:opacity-100 hover:text-error transition-all duration-200"
								title="削除"
							>
								<span class="material-symbols-outlined text-[18px]">delete</span>
							</button>
						</div>
					{/if}
				</div>

				<h2
					class="font-headline-lg text-headline-lg mb-stack-md text-primary transition-colors group-hover:text-primary-container"
				>
					{entry.title}
				</h2>
				<p class="font-body-md text-body-md leading-relaxed text-on-surface-variant whitespace-pre-wrap">
					{entry.content}
				</p>
			</article>
		{/each}
	</div>

	<!-- Pagination or Load More -->
	{#if visibleCount < diariesList.length}
		<div class="flex justify-center">
			<button
				type="button"
				onclick={loadMore}
				class="font-label-md text-label-md cursor-pointer rounded-lg border border-primary px-8 py-3 text-primary transition-all hover:bg-primary hover:text-on-primary active:scale-95"
			>
				過去の記録を見る
			</button>
		</div>
	{/if}
</div>
