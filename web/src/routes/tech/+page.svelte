<script lang="ts">
	let { data } = $props();

	let selectedCategory = $state<string | null>(null);

	const categories = ['All', 'Architecture', 'Development', 'Data Science', 'Newsletter'];

	let filteredArticles = $derived(
		selectedCategory && selectedCategory !== 'All'
			? data.techArticles.filter((article) => article.category === selectedCategory)
			: data.techArticles
	);
</script>

<svelte:head>
	<title>Essence — Tech Feed</title>
</svelte:head>

<div class="editorial-container mx-auto px-gutter">
	<!-- Header Section -->
	<section class="mb-stack-lg">
		<h1 class="font-display-lg text-display-lg mb-stack-sm text-primary">技術と美学</h1>
		<p class="font-body-lg text-body-lg max-w-[600px] text-on-surface-variant">
			思考の断片を、構造化された知性へ。最新のテクノロジーと設計思想を綴る技術録。
		</p>
	</section>

	<!-- Filter Chips -->
	<div class="mb-12 flex flex-wrap gap-2">
		{#each categories as cat (cat)}
			<button
				type="button"
				onclick={() => (selectedCategory = cat === 'All' ? null : cat)}
				class="font-label-sm text-label-sm cursor-pointer rounded-full px-3 py-1 transition-all {selectedCategory ===
					cat ||
				(cat === 'All' && !selectedCategory)
					? 'bg-primary font-semibold text-on-primary'
					: 'bg-surface-container-high text-on-surface-variant hover:bg-surface-container-highest'}"
			>
				{cat}
			</button>
		{/each}
	</div>

	<!-- Featured Article -->
	{#if !selectedCategory || selectedCategory === 'Architecture'}
		<article class="group mb-20 cursor-pointer">
			<div class="mb-stack-sm flex flex-wrap items-center gap-stack-sm">
				<span
					class="font-label-sm text-label-sm rounded-full bg-secondary-container px-3 py-1 text-on-secondary-container"
				>
					{data.featuredArticle.category}
				</span>
				<span class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant">
					<span class="material-symbols-outlined text-[14px]">schedule</span>
					{data.featuredArticle.readTime}
				</span>
				<span class="font-label-sm text-label-sm ml-auto text-on-surface-variant">Featured</span>
			</div>
			<h2
				class="font-display-lg mb-stack-md text-[40px] leading-tight text-primary transition-colors group-hover:text-primary/80"
			>
				{data.featuredArticle.title}
			</h2>
			<p class="font-body-lg text-body-lg mb-6 text-on-surface-variant">
				{data.featuredArticle.excerpt}
			</p>
			<div class="relative h-1 w-full overflow-hidden rounded-full bg-surface-container-high">
				<div class="absolute top-0 left-0 h-full w-1/4 bg-primary/20"></div>
			</div>
		</article>
	{/if}

	<!-- Tech Feed List -->
	<div class="space-y-16">
		{#each filteredArticles as article (article.id)}
			{#if article.isNewsletter}
				<!-- Newsletter / Card Variant -->
				<article
					class="group cursor-pointer rounded-xl border border-outline-variant/30 bg-white/40 p-8 backdrop-blur-sm transition-all duration-500 hover:border-primary/20 md:p-10"
				>
					<div class="mb-6 flex items-center justify-between">
						<div class="flex gap-3">
							<span
								class="font-label-sm text-label-sm rounded-full border border-primary px-3 py-1 text-primary"
							>
								{article.category}
							</span>
							<span class="font-label-sm text-label-sm self-center text-on-surface-variant"
								>Weekly Pick</span
							>
						</div>
						<span class="font-label-sm text-label-sm text-on-surface-variant">{article.date}</span>
					</div>
					<h3 class="font-headline-lg text-headline-lg mb-4 text-primary">
						{article.title}
					</h3>
					<p class="font-body-md text-body-md mb-8 leading-relaxed text-on-surface-variant">
						{article.excerpt}
					</p>
					<div class="flex items-center justify-between border-t border-outline-variant/10 pt-6">
						<div class="flex gap-6">
							<span
								class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant"
							>
								<span class="material-symbols-outlined text-[16px]">timer</span>
								{article.readTime}
							</span>
							{#if article.views}
								<span
									class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant"
								>
									<span class="material-symbols-outlined text-[16px]">visibility</span>
									{article.views}
								</span>
							{/if}
						</div>
						<button
							type="button"
							class="font-label-md text-label-md flex cursor-pointer items-center gap-2 text-primary transition-transform group-hover:translate-x-1"
						>
							記事を読む <span class="material-symbols-outlined !text-[18px]">arrow_forward</span>
						</button>
					</div>
				</article>
			{:else}
				<!-- Standard Tech Article Item -->
				<article class="group flex cursor-pointer flex-col gap-3">
					<div class="flex items-center justify-between">
						<div class="flex items-center gap-stack-sm">
							<span
								class="font-label-sm text-label-sm rounded bg-primary-fixed px-2 py-0.5 text-primary"
							>
								{article.category}
							</span>
							<span
								class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant"
							>
								<span class="material-symbols-outlined text-[14px]">timer</span>
								{article.readTime}
							</span>
						</div>
						<span class="font-label-sm text-label-sm text-on-surface-variant">{article.date}</span>
					</div>
					<h3
						class="font-headline-lg text-headline-lg text-primary decoration-outline-variant decoration-1 underline-offset-4 transition-all group-hover:underline"
					>
						{article.title}
					</h3>
					<p class="font-body-md text-body-md line-clamp-2 max-w-[640px] text-on-surface-variant">
						{article.excerpt}
					</p>
					{#if article.views}
						<div class="mt-2 flex items-center gap-4">
							<span
								class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant"
							>
								<span class="material-symbols-outlined text-[14px]">trending_up</span>
								{article.views}
							</span>
						</div>
					{/if}
				</article>
			{/if}
		{/each}
	</div>
</div>
