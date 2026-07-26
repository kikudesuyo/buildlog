<script lang="ts">
	import type { FeaturedTechArticle, TechArticle } from '$lib/api/types';
	import { resolve } from '$app/paths';

	type Props = {
		featuredArticle: FeaturedTechArticle;
		techArticles: TechArticle[];
		isAdmin?: boolean;
		onEdit?: (id: number) => void;
		onDelete?: (id: number) => void | Promise<boolean | void>;
	};

	let { featuredArticle, techArticles, isAdmin = false, onEdit, onDelete }: Props = $props();
	let deletedIds = $state<number[]>([]);
	let selectedCategory = $state<string | null>(null);
	const categories = ['All', 'Architecture', 'Development', 'Data Science', 'Newsletter'];
	let articles = $derived(
		(featuredArticle.title ? [featuredArticle, ...techArticles] : techArticles).filter(
			(article) => !deletedIds.includes(article.id)
		)
	);
	let featured = $derived(articles.length > 0 ? articles[0] : null);
	let filteredArticles = $derived(
		(selectedCategory ? articles.slice(1).filter((article) => article.category === selectedCategory) : articles.slice(1))
	);

	function formatDate(dateStr: string) {
		if (!dateStr) return '';
		const date = new Date(dateStr);
		return `${date.getFullYear()}.${String(date.getMonth() + 1).padStart(2, '0')}.${String(date.getDate()).padStart(2, '0')}`;
	}

	async function deleteArticle(id: number) {
		if ((await onDelete?.(id)) !== false) deletedIds = [...deletedIds, id];
	}
</script>

<div class="editorial-container mx-auto flex flex-col gap-8 px-gutter">
	<section class="flex items-start justify-between">
		<div>
			<h1 class="font-display-lg text-display-lg mb-stack-sm text-primary">技術と美学</h1>
			<p class="font-body-lg text-body-lg max-w-[600px] text-on-surface-variant">思考の断片を、構造化された知性へ。最新のテクノロジーと設計思想を綴る技術録。</p>
		</div>
		{#if isAdmin}
			<a href={resolve(isAdmin ? '/admin/tech/new' : '/tech/new')} class="font-label-md text-label-md flex shrink-0 cursor-pointer items-center gap-1.5 rounded-lg bg-primary px-6 py-2.5 text-on-primary transition-all hover:bg-primary/95 active:scale-95"><span class="material-symbols-outlined text-[18px]">add</span>記事を書く</a>
		{/if}
	</section>

	<div class="flex flex-wrap gap-2">
		{#each categories as category (category)}
			<button type="button" onclick={() => (selectedCategory = category === 'All' ? null : category)} class="font-label-sm text-label-sm cursor-pointer rounded-full px-3 py-1 transition-all {selectedCategory === category || (category === 'All' && !selectedCategory) ? 'bg-primary font-semibold text-on-primary' : 'bg-surface-container-high text-on-surface-variant hover:bg-surface-container-highest'}">{category}</button>
		{/each}
	</div>

	{#if (!selectedCategory || selectedCategory === 'Architecture') && featured}
		<article class="group relative -mx-4 rounded-xl border border-transparent p-4">
			<div class="mb-stack-sm flex flex-wrap items-center gap-stack-sm">
				<span class="font-label-sm text-label-sm rounded-full bg-secondary-container px-3 py-1 text-on-secondary-container">{featured.category}</span>
				<span class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant"><span class="material-symbols-outlined text-[14px]">schedule</span>{featured.readTime}</span>
				<span class="font-label-sm text-label-sm text-on-surface-variant">Featured</span>
				{#if isAdmin}<div class="ml-auto flex gap-2"><button type="button" onclick={() => onEdit?.(featured.id)} class="p-1 text-outline opacity-60 hover:text-primary hover:opacity-100" title="編集"><span class="material-symbols-outlined text-[18px]">edit</span></button><button type="button" onclick={() => deleteArticle(featured.id)} class="p-1 text-outline opacity-60 hover:text-error hover:opacity-100" title="削除"><span class="material-symbols-outlined text-[18px]">delete</span></button></div>{/if}
			</div>
			<h2 class="font-display-lg mb-stack-md text-[40px] leading-tight text-primary transition-colors group-hover:text-primary/80">{featured.title}</h2>
			<p class="font-body-lg text-body-lg mb-6 text-on-surface-variant">{featured.excerpt}</p>
			<div class="relative h-1 w-full overflow-hidden rounded-full bg-surface-container-high"><div class="absolute top-0 left-0 h-full w-1/4 bg-primary/20"></div></div>
		</article>
	{/if}

	<div class="space-y-12">
		{#each filteredArticles as article (article.id)}
			<article class:rounded-xl={article.isNewsletter} class="group relative flex flex-col gap-3 border border-transparent p-4 {article.isNewsletter ? 'border-outline-variant/30 bg-white/40 p-8 backdrop-blur-sm md:p-10' : ''}">
				<div class="flex items-center justify-between">
					<div class="flex items-center gap-stack-sm"><span class="font-label-sm text-label-sm rounded bg-primary-fixed px-2 py-0.5 text-primary">{article.category}</span><span class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant"><span class="material-symbols-outlined text-[14px]">timer</span>{article.readTime}</span></div>
					<div class="flex items-center gap-4"><span class="font-label-sm text-label-sm text-on-surface-variant">{formatDate(article.createdAt)}</span>{#if isAdmin}<div class="flex gap-2"><button type="button" onclick={() => onEdit?.(article.id)} class="p-1 text-outline opacity-60 hover:text-primary hover:opacity-100" title="編集"><span class="material-symbols-outlined text-[18px]">edit</span></button><button type="button" onclick={() => deleteArticle(article.id)} class="p-1 text-outline opacity-60 hover:text-error hover:opacity-100" title="削除"><span class="material-symbols-outlined text-[18px]">delete</span></button></div>{/if}</div>
				</div>
				<h3 class="font-headline-lg text-headline-lg text-primary decoration-outline-variant decoration-1 underline-offset-4 transition-all group-hover:underline">{article.title}</h3>
				<p class="font-body-md text-body-md line-clamp-2 max-w-[640px] text-on-surface-variant">{article.excerpt}</p>
				{#if article.views}<div class="mt-2 flex items-center gap-4"><span class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant"><span class="material-symbols-outlined text-[14px]">trending_up</span>{article.views}</span></div>{/if}
			</article>
		{/each}
	</div>
</div>
