<script lang="ts">
	import { resolve } from '$app/paths';
	let { data } = $props();

	function formatDate(dateStr: string) {
		if (!dateStr) return '';
		const date = new Date(dateStr);
		return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`;
	}
</script>

<svelte:head>
	<title>{data.tech.title} — Buildlog</title>
</svelte:head>

<div class="editorial-container mx-auto px-gutter pt-24 pb-20 max-w-[800px]">
	<!-- 戻るリンク -->
	<a
		href={resolve('/tech')}
		class="font-label-md text-label-md mb-8 inline-flex items-center gap-2 text-outline hover:text-primary transition-colors duration-200"
	>
		<span class="material-symbols-outlined text-[18px]">arrow_back</span>
		技術録へ戻る
	</a>

	<article class="flex flex-col gap-6">
		<header class="flex flex-col gap-4 border-b border-outline-variant/20 pb-8">
			<!-- メタ情報 -->
			<div class="flex flex-wrap items-center gap-stack-sm">
				<span class="font-label-sm text-label-sm rounded-full bg-secondary-container px-3 py-1 text-on-secondary-container">
					{data.tech.category}
				</span>
				<span class="font-label-sm text-label-sm text-on-surface-variant">
					{formatDate(data.tech.createdAt)}
				</span>
				{#if data.tech.views}
					<span class="h-1.5 w-1.5 rounded-full bg-outline-variant/30"></span>
					<span class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant">
						<span class="material-symbols-outlined text-[14px]">trending_up</span>
						{data.tech.views}
					</span>
				{/if}
			</div>

			<!-- タグリスト -->
			{#if data.tech.tags && data.tech.tags.length > 0}
				<div class="flex flex-wrap gap-1.5 mt-1">
					{#each data.tech.tags as tag (tag)}
						<a
							href={resolve('/tech?tag=' + encodeURIComponent(tag))}
							class="font-label-sm text-[11px] px-2.5 py-0.5 rounded-full bg-surface-container-high text-on-surface-variant hover:bg-primary hover:text-on-primary transition-all cursor-pointer"
						>
							#{tag}
						</a>
					{/each}
				</div>
			{/if}

			<!-- タイトル -->
			<h1 class="font-display-lg text-[40px] leading-tight text-primary font-bold tracking-tight">
				{data.tech.title}
			</h1>
		</header>

		<!-- 本文 (Content) -->
		<section class="font-body-md text-body-md leading-relaxed whitespace-pre-wrap text-on-surface pt-4">
			{data.tech.content}
		</section>
	</article>
</div>
