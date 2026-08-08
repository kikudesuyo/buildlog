<script lang="ts">
	import { onMount } from 'svelte';
	import type { DiaryEntry } from '$lib/api/types';
	import { fetchDiaryEntries, type ApiFetch } from '$lib/api/client';
	import { resolve } from '$app/paths';
	import { tick } from 'svelte';
	import LikeButton from './LikeButton.svelte';

	type Props = {
		entries: DiaryEntry[];
		isAdmin?: boolean;
		onEdit?: (id: number) => void;
		onDelete?: (id: number) => void | Promise<boolean | void>;
	};

	let { entries: initialEntries, isAdmin = false, onEdit, onDelete }: Props = $props();
	let entries = $state(initialEntries);
	const pageSize = 3;
	let isLoading = $state(false);
	let loadError = $state(false);
	let hasMore = $state(!isAdmin && initialEntries.length === pageSize);
	const storageKey = `diary-feed-count:${isAdmin ? 'admin' : 'public'}`;
	let restoreTarget = 0;
	let displayEntries = $derived(entries);
	let actionRefs = $state<Record<number, HTMLButtonElement | undefined>>({});
	let headingRef: HTMLHeadingElement;
	let notification = $state<{ message: string; kind: 'success' | 'error' } | null>(null);

	function formatDate(dateStr: string) {
		if (!dateStr) return '';
		const date = new Date(dateStr);
		return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`;
	}

	async function deleteEntry(entry: DiaryEntry) {
		if (!confirm(`「${entry.title}」を削除しますか？`)) return;

		const entryIndex = displayEntries.findIndex((displayEntry) => displayEntry.id === entry.id);
		const nextFocusId = displayEntries[entryIndex + 1]?.id ?? displayEntries[entryIndex - 1]?.id;
		if ((await onDelete?.(entry.id)) === false) {
			notification = { message: `「${entry.title}」の削除に失敗しました。`, kind: 'error' };
			return;
		}

		entries = entries.filter((currentEntry) => currentEntry.id !== entry.id);
		notification = { message: `「${entry.title}」を削除しました。`, kind: 'success' };
		await tick();
		(actionRefs[nextFocusId ?? -1] ?? headingRef)?.focus();
	}

	async function loadMore() {
		if (isLoading || !hasMore) return;
		isLoading = true;
		loadError = false;
		try {
			const nextEntries = await fetchDiaryEntries(fetch as ApiFetch, isAdmin, entries.length, pageSize);
			entries = [...entries, ...nextEntries];
			hasMore = nextEntries.length === pageSize;
			sessionStorage.setItem(storageKey, String(entries.length));
		} catch {
			loadError = true;
		} finally {
			isLoading = false;
		}
	}

	onMount(async () => {
		const savedCount = Number(sessionStorage.getItem(storageKey));
		if (savedCount > entries.length) {
			restoreTarget = savedCount;
			while (entries.length < restoreTarget && hasMore && !loadError) await loadMore();
		}
	});
</script>

<div class="editorial-container mx-auto px-gutter relative flex flex-col gap-8">
	<header class="flex items-center justify-between">
		<div>
			{#if isAdmin}
				<p class="font-label-sm text-label-sm mb-stack-sm tracking-[0.2em] text-outline uppercase">Content Manager / Diary</p>
			{/if}
			<h1 bind:this={headingRef} tabindex="-1" class="font-display-lg text-display-lg text-primary">{isAdmin ? 'つぶやき管理' : '日々のつぶやき'}</h1>
	</div>
		{#if isAdmin}
			<a
				href={resolve(isAdmin ? '/admin/diary/new' : '/diary/new')}
				class="font-label-md text-label-md flex cursor-pointer items-center gap-1.5 rounded-lg bg-primary px-6 py-2.5 text-on-primary transition-all hover:bg-primary/95 active:scale-95"
			>
				<span class="material-symbols-outlined text-[18px]">add</span>
				つぶやく
			</a>
		{/if}
	</header>

	<div class="flex flex-col gap-6">
		{#each displayEntries as entry (entry.id)}
			<article class="group relative rounded-xl border border-outline-variant/30 bg-surface-container-lowest p-6 shadow-2xs transition-all duration-300 hover:shadow-md hover:border-primary/20">
				<div class="mb-stack-sm flex items-center justify-between">
					<div class="flex items-center gap-3">
						<span class="font-label-sm text-label-sm text-outline">{formatDate(entry.createdAt)}</span>
						{#if entry.status === 'draft'}
							<span class="font-label-sm text-label-sm px-2 py-0.5 rounded bg-outline-variant/40 text-on-surface-variant">下書き</span>
						{/if}
					</div>
					{#if isAdmin}
						<div class="flex gap-2">
							<button type="button" onclick={() => onEdit?.(entry.id)} class="min-h-11 min-w-11 rounded-lg p-1 text-outline opacity-60 transition-all duration-200 hover:text-primary hover:opacity-100" title={`編集: ${entry.title}`} aria-label={`編集: ${entry.title}`}>
								<span class="material-symbols-outlined text-[18px]">edit</span>
							</button>
							<button bind:this={actionRefs[entry.id]} type="button" onclick={() => deleteEntry(entry)} class="min-h-11 min-w-11 rounded-lg p-1 text-outline opacity-60 transition-all duration-200 hover:text-error hover:opacity-100" title={`削除: ${entry.title}`} aria-label={`削除: ${entry.title}`}>
								<span class="material-symbols-outlined text-[18px]">delete</span>
							</button>
						</div>
					{/if}
				</div>
				<a href={isAdmin ? undefined : resolve(`/diary/${entry.id}`)} class="block">
					<h2 class="font-headline-lg text-headline-lg mb-stack-md text-primary transition-colors group-hover:text-primary-container">{entry.title}</h2>
					<p class="font-body-md text-body-md leading-relaxed whitespace-pre-wrap text-on-surface-variant mb-4 line-clamp-3">{entry.content}</p>
					{#if !isAdmin}
						<span class="font-label-md text-label-md text-primary group-hover:underline">続きを読む</span>
					{/if}
				</a>
				<div class="flex items-center justify-between">
					<div class="flex items-center gap-4">
						<LikeButton postId={entry.id} initialLikesCount={entry.likesCount} initialHasLiked={entry.hasLiked} />
						<a href={resolve(`/diary/${entry.id}#comments-heading`)} class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant hover:text-primary hover:underline" aria-label={`コメント${entry.commentsCount}件を表示`}>
							<span class="material-symbols-outlined text-[16px]" aria-hidden="true">comment</span>
							{entry.commentsCount}
						</a>
					</div>
				</div>
			</article>
		{/each}
	</div>

	{#if notification}
		<p class="sr-only" role={notification.kind === 'error' ? 'alert' : 'status'} aria-live="polite">{notification.message}</p>
	{/if}

	{#if loadError}
		<div class="flex flex-col items-center gap-3" role="alert">
			<p class="font-body-sm text-body-sm text-error">過去の記録を読み込めませんでした。</p>
			<button type="button" onclick={loadMore} class="font-label-md text-label-md cursor-pointer rounded-lg border border-error px-6 py-2 text-error transition-colors hover:bg-error hover:text-on-error">再試行</button>
		</div>
	{:else if isLoading}
		<div class="flex items-center justify-center gap-2" role="status" aria-live="polite">
			<span class="material-symbols-outlined animate-spin text-primary" aria-hidden="true">progress_activity</span>
			<span class="font-body-sm text-body-sm text-on-surface-variant">読み込み中…</span>
		</div>
	{:else if hasMore}
		<div class="flex justify-center">
			<button type="button" onclick={loadMore} class="font-label-md text-label-md cursor-pointer rounded-lg border border-primary px-8 py-3 text-primary transition-all hover:bg-primary hover:text-on-primary active:scale-95">
				過去の記録を見る
			</button>
		</div>
	{:else if !isAdmin}
		<p class="font-body-sm text-body-sm text-center text-outline" role="status">すべての記録を表示しました。</p>
	{/if}
</div>
