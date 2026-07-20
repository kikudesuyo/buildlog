<script lang="ts">
	let { data } = $props();

	let visibleCount = $state(3);
	let displayEntries = $derived(data.diaryEntries.slice(0, visibleCount));

	function loadMore() {
		visibleCount += 2;
	}
</script>

<svelte:head>
	<title>Essence — Diary</title>
</svelte:head>

<div class="editorial-container mx-auto px-gutter">
	<!-- Intro / Header -->
	<header class="mb-stack-lg">
		<h1 class="font-display-lg text-display-lg mb-stack-sm text-primary">日々のつぶやき</h1>
	</header>

	<!-- Blog Post List -->
	<div class="flex flex-col gap-section-gap">
		{#each displayEntries as entry (entry.id)}
			<article class="group cursor-pointer">
				{#if entry.date}
					<div class="mb-stack-sm flex items-center">
						<span class="font-label-sm text-label-sm text-outline">{entry.date}</span>
					</div>
				{/if}

				<h2
					class="font-headline-lg text-headline-lg mb-stack-md text-primary transition-colors group-hover:text-primary-container"
				>
					{entry.title}
				</h2>
				<p class="font-body-md text-body-md mb-stack-md leading-relaxed text-on-surface-variant">
					{entry.excerpt}
				</p>
			</article>
		{/each}
	</div>

	<!-- Pagination or Load More -->
	{#if visibleCount < data.diaryEntries.length}
		<div class="mt-section-gap flex justify-center">
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
