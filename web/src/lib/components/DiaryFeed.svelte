<script lang="ts">
	import type { DiaryEntry } from '$lib/api/types';
	import { resolve } from '$app/paths';

	type Props = {
		entries: DiaryEntry[];
		isAdmin?: boolean;
		onEdit?: (id: number) => void;
		onDelete?: (id: number) => void | Promise<boolean | void>;
	};

	let { entries: initialEntries, isAdmin = false, onEdit, onDelete }: Props = $props();
	let entries = $state(initialEntries);
	let visibleCount = $state(3);
	let displayEntries = $derived(entries.slice(0, visibleCount));

	function formatDate(dateStr: string) {
		if (!dateStr) return '';
		const date = new Date(dateStr);
		return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`;
	}

	async function deleteEntry(id: number) {
		if ((await onDelete?.(id)) !== false) entries = entries.filter((entry) => entry.id !== id);
	}
</script>

<div class="editorial-container mx-auto px-gutter relative flex flex-col gap-8">
	<header class="flex items-center justify-between">
		<div>
			{#if isAdmin}
				<p class="font-label-sm text-label-sm mb-stack-sm tracking-[0.2em] text-outline uppercase">Content Manager / Diary</p>
			{/if}
			<h1 class="font-display-lg text-display-lg text-primary">{isAdmin ? 'つぶやき管理' : '日々のつぶやき'}</h1>
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
			<article class="group relative -mx-4 rounded-xl border border-transparent p-4">
				<div class="mb-stack-sm flex items-center justify-between">
					<span class="font-label-sm text-label-sm text-outline">{formatDate(entry.createdAt)}</span>
					{#if isAdmin}
						<div class="flex gap-2">
							<button type="button" onclick={() => onEdit?.(entry.id)} class="p-1 text-outline opacity-60 transition-all duration-200 hover:text-primary hover:opacity-100" title="編集">
								<span class="material-symbols-outlined text-[18px]">edit</span>
							</button>
							<button type="button" onclick={() => deleteEntry(entry.id)} class="p-1 text-outline opacity-60 transition-all duration-200 hover:text-error hover:opacity-100" title="削除">
								<span class="material-symbols-outlined text-[18px]">delete</span>
							</button>
						</div>
					{/if}
				</div>
				<h2 class="font-headline-lg text-headline-lg mb-stack-md text-primary transition-colors group-hover:text-primary-container">{entry.title}</h2>
				<p class="font-body-md text-body-md leading-relaxed whitespace-pre-wrap text-on-surface-variant">{entry.content}</p>
				{#if entry.tags && entry.tags.length > 0}
					<div class="mt-3 flex flex-wrap gap-1.5">
						{#each entry.tags as tag (tag)}
							<a
								href={resolve('/?tag=' + encodeURIComponent(tag))}
								class="font-label-sm text-[11px] px-2 py-0.5 rounded-full bg-surface-container-high text-on-surface-variant hover:bg-primary hover:text-on-primary transition-all cursor-pointer"
							>
								#{tag}
							</a>
						{/each}
					</div>
				{/if}
			</article>
		{/each}
	</div>

	{#if visibleCount < entries.length}
		<div class="flex justify-center">
			<button type="button" onclick={() => (visibleCount += 2)} class="font-label-md text-label-md cursor-pointer rounded-lg border border-primary px-8 py-3 text-primary transition-all hover:bg-primary hover:text-on-primary active:scale-95">
				過去の記録を見る
			</button>
		</div>
	{/if}
</div>
