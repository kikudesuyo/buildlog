<script lang="ts">
	import { goto } from '$app/navigation';
	import { deleteTech } from '$lib/api/client';
	import type { TechArticle } from '$lib/api/types';

	let { data } = $props();

	let selectedCategory = $state<string | null>(null);

	const categories = ['All', 'Architecture', 'Development', 'Data Science', 'Newsletter'];

	// 記事の一元管理
	let allArticles = $state<TechArticle[]>([]);

	// Svelte 5 reactive 同期
	$effect(() => {
		allArticles = [
			...(data.featuredArticle && data.featuredArticle.title ? [data.featuredArticle as TechArticle] : []),
			...data.techArticles
		];
	});

	// 一元管理した配列から derived で Featured と残りのリストを動的抽出
	let featuredArticle = $derived(allArticles.length > 0 ? allArticles[0] : null);
	let techArticlesList = $derived(allArticles.length > 1 ? allArticles.slice(1) : []);

	let filteredArticles = $derived(
		selectedCategory && selectedCategory !== 'All'
			? techArticlesList.filter((article) => article.category === selectedCategory)
			: techArticlesList
	);

	function formatDate(dateStr: string) {
		if (!dateStr) return '';
		const date = new Date(dateStr);
		return `${date.getFullYear()}.${String(date.getMonth() + 1).padStart(2, '0')}.${String(date.getDate()).padStart(2, '0')}`;
	}

	async function handleDelete(id: number) {
		if (!confirm('この記事を削除してもよろしいですか？')) return;
		try {
			await deleteTech(id);
			allArticles = allArticles.filter((a) => a.id !== id);
		} catch (err) {
			alert('削除に失敗しました。');
		}
	}
</script>

<svelte:head>
	<title>Essence — Tech Feed</title>
</svelte:head>

<div class="editorial-container mx-auto px-gutter flex flex-col gap-8">
	<!-- Header Section -->
	<section class="flex items-start justify-between">
		<div>
			<h1 class="font-display-lg text-display-lg mb-stack-sm text-primary">技術と美学</h1>
			<p class="font-body-lg text-body-lg max-w-[600px] text-on-surface-variant">
				思考の断片を、構造化された知性へ。最新のテクノロジーと設計思想を綴る技術録。
			</p>
		</div>
		<button
			type="button"
			onclick={() => goto('/tech/new')}
			class="font-label-md text-label-md cursor-pointer rounded-lg bg-primary px-6 py-2.5 text-on-primary transition-all hover:bg-primary/95 active:scale-95 flex items-center gap-1.5 shrink-0"
		>
			<span class="material-symbols-outlined text-[18px]">add</span>
			記事を書く
		</button>
	</section>

	<!-- Filter Chips -->
	<div class="flex flex-wrap gap-2">
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
	{#if (!selectedCategory || selectedCategory === 'Architecture') && featuredArticle}
		<article class="group relative rounded-xl border border-transparent p-4 -mx-4 transition-all duration-300 hover:bg-surface-container-low hover:border-outline-variant/10">
			<div class="mb-stack-sm flex flex-wrap items-center gap-stack-sm">
				<span
					class="font-label-sm text-label-sm rounded-full bg-secondary-container px-3 py-1 text-on-secondary-container"
				>
					{featuredArticle.category}
				</span>
				<span class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant">
					<span class="material-symbols-outlined text-[14px]">schedule</span>
					{featuredArticle.readTime}
				</span>
				<span class="font-label-sm text-label-sm text-on-surface-variant">Featured</span>

				<!-- アクションボタン -->
				<div class="flex gap-2 ml-auto opacity-55 group-hover:opacity-100 transition-opacity">
					<button
						type="button"
						onclick={() => goto(`/tech/${featuredArticle.id}/edit`)}
						class="p-1 hover:text-primary transition-colors cursor-pointer text-outline hover:opacity-100"
						title="編集"
					>
						<span class="material-symbols-outlined text-[18px]">edit</span>
					</button>
					<button
						type="button"
						onclick={() => handleDelete(featuredArticle.id)}
						class="p-1 hover:text-error transition-colors cursor-pointer text-outline hover:opacity-100"
						title="削除"
					>
						<span class="material-symbols-outlined text-[18px]">delete</span>
					</button>
				</div>
			</div>
			<h2
				class="font-display-lg mb-stack-md text-[40px] leading-tight text-primary transition-colors group-hover:text-primary/80"
			>
				{featuredArticle.title}
			</h2>
			<p class="font-body-lg text-body-lg mb-6 text-on-surface-variant">
				{featuredArticle.excerpt}
			</p>
			<div class="relative h-1 w-full overflow-hidden rounded-full bg-surface-container-high">
				<div class="absolute top-0 left-0 h-full w-1/4 bg-primary/20"></div>
			</div>
		</article>
	{/if}

	<!-- Tech Feed List -->
	<div class="space-y-12">
		{#each filteredArticles as article (article.id)}
			{#if article.isNewsletter}
				<!-- Newsletter / Card Variant -->
				<article
					class="group relative rounded-xl border border-outline-variant/30 bg-white/40 p-8 backdrop-blur-sm transition-all duration-500 hover:border-primary/20 md:p-10"
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
						<div class="flex items-center gap-4">
							<span class="font-label-sm text-label-sm text-on-surface-variant">{formatDate(article.createdAt)}</span>
							
							<!-- アクションボタン -->
							<div class="flex gap-2 opacity-55 group-hover:opacity-100 transition-opacity">
								<button
									type="button"
									onclick={() => goto(`/tech/${article.id}/edit`)}
									class="p-1 hover:text-primary transition-colors cursor-pointer text-outline hover:opacity-100"
									title="編集"
								>
									<span class="material-symbols-outlined text-[18px]">edit</span>
								</button>
								<button
									type="button"
									onclick={() => handleDelete(article.id)}
									class="p-1 hover:text-error transition-colors cursor-pointer text-outline hover:opacity-100"
									title="削除"
								>
									<span class="material-symbols-outlined text-[18px]">delete</span>
								</button>
							</div>
						</div>
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
					</div>
				</article>
			{:else}
				<!-- Standard Tech Article Item -->
				<article class="group relative flex flex-col gap-3 rounded-xl border border-transparent p-4 -mx-4 transition-all duration-300 hover:bg-surface-container-low hover:border-outline-variant/10">
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
						<div class="flex items-center gap-4">
							<span class="font-label-sm text-label-sm text-on-surface-variant">{formatDate(article.createdAt)}</span>
							
							<!-- アクションボタン -->
							<div class="flex gap-2 opacity-55 group-hover:opacity-100 transition-opacity">
								<button
									type="button"
									onclick={() => goto(`/tech/${article.id}/edit`)}
									class="p-1 hover:text-primary transition-colors cursor-pointer text-outline hover:opacity-100"
									title="編集"
								>
									<span class="material-symbols-outlined text-[18px]">edit</span>
								</button>
								<button
									type="button"
									onclick={() => handleDelete(article.id)}
									class="p-1 hover:text-error transition-colors cursor-pointer text-outline hover:opacity-100"
									title="削除"
								>
									<span class="material-symbols-outlined text-[18px]">delete</span>
								</button>
							</div>
						</div>
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
