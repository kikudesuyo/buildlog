<script lang="ts">
	import { onMount } from 'svelte';
	import type { FeaturedTechArticle, TechArticle } from '$lib/api/types';
	import { resolve } from '$app/paths';
	import { techCategories } from '$lib/tech/categories';
	import LikeButton from './LikeButton.svelte';
	import { invalidateAll } from '$app/navigation';
	import { fetchTechFeed, type ApiFetch } from '$lib/api/client';

	type Props = {
		featuredArticle?: FeaturedTechArticle | null;
		techArticles?: TechArticle[];
		hasMore?: boolean;
		loadError?: boolean;
		isAdmin?: boolean;
		onEdit?: (id: number) => void;
		onDelete?: (id: number) => void | Promise<boolean | void>;
	};

	let { featuredArticle = null, techArticles = [], hasMore = false, loadError = false, isAdmin = false, onEdit, onDelete }: Props = $props();
	let loadedTechArticles = $state(techArticles);
	let loadedFeaturedArticle = $state(featuredArticle);
	let hasMoreArticles = $state(hasMore);
	let isLoadingMore = $state(false);
	let loadMoreError = $state(false);
	let deletedIds = $state<number[]>([]);
	let selectedCategory = $state<string | null>(null);
	const categories = ['All', ...techCategories];
	let categoryScroller = $state<HTMLDivElement | null>(null);
	let canScrollLeft = $state(false);
	let canScrollRight = $state(false);
	let articles = $derived(
		(loadedFeaturedArticle?.title ? [loadedFeaturedArticle, ...loadedTechArticles] : loadedTechArticles).filter(
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

	async function loadMore() {
		if (isLoadingMore || !hasMoreArticles) return;
		isLoadingMore = true;
		loadMoreError = false;
		try {
			const offset = (loadedFeaturedArticle?.title ? 1 : 0) + loadedTechArticles.length;
			const next = await fetchTechFeed(window.fetch.bind(window) as ApiFetch, false, offset, 3);
			loadedTechArticles = [...loadedTechArticles, ...next.techArticles];
			hasMoreArticles = next.hasMore;
		} catch {
			loadMoreError = true;
		} finally {
			isLoadingMore = false;
		}
	}

	function updateCategoryOverflow() {
		const scroller = categoryScroller;
		if (!scroller) return;
		canScrollLeft = scroller.scrollLeft > 0;
		canScrollRight = scroller.scrollLeft + scroller.clientWidth < scroller.scrollWidth - 1;
	}

	function selectCategory(category: string) {
		selectedCategory = category === 'All' ? null : category;
		requestAnimationFrame(() => {
			categoryScroller?.querySelector<HTMLElement>(`[data-category="${category}"]`)?.scrollIntoView({
				behavior: 'smooth',
				block: 'nearest',
				inline: 'center'
			});
			updateCategoryOverflow();
		});
	}

	onMount(() => {
		updateCategoryOverflow();
		window.addEventListener('resize', updateCategoryOverflow);
		return () => window.removeEventListener('resize', updateCategoryOverflow);
	});
</script>

<div class="editorial-container mx-auto flex flex-col gap-8 px-gutter">
	<section class="flex items-start justify-between">
		<div>
			{#if isAdmin}
				<p class="font-label-sm text-label-sm mb-stack-sm tracking-[0.2em] text-outline uppercase">Content Manager / Tech</p>
			{/if}
			<h1 class="font-display-lg text-display-lg mb-stack-sm text-primary">{isAdmin ? '技術記事管理' : '技術と美学'}</h1>
			<p class="font-body-lg text-body-lg max-w-[600px] text-on-surface-variant">思考の断片を、構造化された知性へ。最新のテクノロジーと設計思想を綴る技術録。</p>
		</div>
		{#if isAdmin}
			<a href={resolve(isAdmin ? '/admin/tech/new' : '/tech/new')} class="font-label-md text-label-md flex shrink-0 cursor-pointer items-center gap-1.5 rounded-lg bg-primary px-6 py-2.5 text-on-primary transition-all hover:bg-primary/95 active:scale-95"><span class="material-symbols-outlined text-[18px]">add</span>記事を書く</a>
		{/if}
	</section>

	<div class="relative" aria-label="技術記事カテゴリ">
		<div
			bind:this={categoryScroller}
			class="category-scroller flex snap-x snap-mandatory gap-2 overflow-x-auto overscroll-contain px-1 py-1"
			tabindex="0"
			onscroll={updateCategoryOverflow}
		>
			{#each categories as category (category)}
				<button
					type="button"
					data-category={category}
					aria-pressed={selectedCategory === category || (category === 'All' && !selectedCategory)}
					onclick={() => selectCategory(category)}
					class="font-label-sm text-label-sm min-h-11 shrink-0 snap-start cursor-pointer rounded-full px-4 py-2 transition-all {selectedCategory === category || (category === 'All' && !selectedCategory) ? 'bg-primary font-semibold text-on-primary' : 'bg-surface-container-high text-on-surface-variant hover:bg-surface-container-highest'}"
				>
					{category}
				</button>
			{/each}
		</div>
		{#if canScrollLeft}<span class="pointer-events-none absolute inset-y-0 left-0 flex items-center bg-gradient-to-r from-surface to-transparent pr-3 pl-1 text-primary" aria-hidden="true">‹</span>{/if}
		{#if canScrollRight}<span class="pointer-events-none absolute inset-y-0 right-0 flex items-center bg-gradient-to-l from-surface to-transparent pl-3 pr-1 text-primary" aria-hidden="true">›</span>{/if}
		<p class="sr-only" aria-live="polite">{selectedCategory ?? 'All'}カテゴリを選択中</p>
	</div>

	{#if loadError}
		<section class="rounded-xl border border-error/30 bg-error-container/30 p-8 text-center" role="alert">
			<h2 class="font-headline-md text-headline-md text-on-surface">記事を読み込めませんでした</h2>
			<p class="font-body-md text-body-md mt-2 text-on-surface-variant">通信を確認して、もう一度お試しください。</p>
			<button type="button" onclick={() => invalidateAll()} class="font-label-md text-label-md mt-5 min-h-11 rounded-lg bg-primary px-5 py-2 text-on-primary">再試行</button>
		</section>
	{:else if articles.length === 0}
		<section class="rounded-xl border border-outline-variant/30 bg-surface-container-low p-8 text-center">
			<h2 class="font-headline-md text-headline-md text-primary">技術記事はまだありません</h2>
			<p class="font-body-md text-body-md mt-2 text-on-surface-variant">新しい記事が公開されるまでお待ちください。</p>
		</section>
	{:else if selectedCategory && filteredArticles.length === 0}
		<section class="rounded-xl border border-outline-variant/30 bg-surface-container-low p-8 text-center">
			<h2 class="font-headline-md text-headline-md text-primary">該当する記事がありません</h2>
			<p class="font-body-md text-body-md mt-2 text-on-surface-variant">「{selectedCategory}」の記事はありません。</p>
			<button type="button" onclick={() => (selectedCategory = null)} class="font-label-md text-label-md mt-5 min-h-11 rounded-lg border border-outline-variant/50 px-5 py-2 text-primary">Allへ戻る</button>
		</section>
	{/if}

	{#if (!selectedCategory || selectedCategory === featured?.category) && featured}
		<article aria-label="Featured article" class="group relative rounded-xl border border-outline-variant/40 bg-surface-container-low p-4 shadow-xs transition-all duration-300 hover:shadow-md hover:border-primary/20 md:p-6">
			<div class="mb-stack-sm flex flex-wrap items-center gap-stack-sm">
				<span class="font-label-sm text-label-sm rounded-full bg-secondary-container px-3 py-1 text-on-secondary-container">{featured.category}</span>
				<span class="font-label-sm text-label-sm text-on-surface-variant">Featured</span>
				{#if featured.status === 'draft'}
					<span class="font-label-sm text-label-sm px-2 py-0.5 rounded bg-outline-variant/40 text-on-surface-variant">下書き</span>
				{/if}
				{#if isAdmin}<div class="ml-auto flex gap-2"><button type="button" onclick={() => onEdit?.(featured.id)} class="p-1 text-outline opacity-60 hover:text-primary hover:opacity-100" title="編集"><span class="material-symbols-outlined text-[18px]">edit</span></button><button type="button" onclick={() => deleteArticle(featured.id)} class="p-1 text-outline opacity-60 hover:text-error hover:opacity-100" title="削除"><span class="material-symbols-outlined text-[18px]">delete</span></button></div>{/if}
			</div>
			<h2 class="font-display-lg mb-stack-md text-[22px] leading-tight text-primary transition-colors group-hover:text-primary/80 md:text-[28px]">
				<a href={resolve(`/tech/${featured.id}`)} class="hover:underline">{featured.title}</a>
			</h2>
			<p class="font-body-md text-body-md mb-4 text-on-surface-variant line-clamp-2 md:mb-6 md:line-clamp-3">{featured.content}</p>
			<a href={resolve(`/tech/${featured.id}`)} class="font-label-md text-label-md text-primary hover:underline">続きを読む</a>
			<div class="flex items-center justify-between mb-4">
				<LikeButton postId={featured.id} initialLikesCount={featured.likesCount} initialHasLiked={featured.hasLiked} />
				<a href={resolve(`/tech/${featured.id}`)} class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant hover:text-primary hover:underline" aria-label={`コメント${featured.commentsCount}件を表示`}>
					<span class="material-symbols-outlined text-[16px]" aria-hidden="true">comment</span>{featured.commentsCount}
				</a>
			</div>
			<div class="relative h-1 w-full overflow-hidden rounded-full bg-surface-container-high"><div class="absolute top-0 left-0 h-full w-1/4 bg-primary/20"></div></div>
		</article>
	{/if}

	<div class="space-y-6 md:space-y-12">
		{#each filteredArticles as article (article.id)}
			<article class="group relative flex flex-col gap-3 rounded-xl border border-outline-variant/20 bg-surface-container-lowest p-4 shadow-2xs transition-all duration-300 hover:shadow-md hover:border-primary/20 md:p-6">
				<div class="flex items-center justify-between">
					<div class="flex items-center gap-stack-sm">
						<span class="font-label-sm text-label-sm rounded bg-primary-fixed px-2 py-0.5 text-primary">{article.category}</span>
						{#if article.status === 'draft'}
							<span class="font-label-sm text-label-sm px-2 py-0.5 rounded bg-outline-variant/40 text-on-surface-variant">下書き</span>
						{/if}
					</div>
					<div class="flex items-center gap-4"><span class="font-label-sm text-label-sm text-on-surface-variant">{formatDate(article.createdAt)}</span>{#if isAdmin}<div class="flex gap-2"><button type="button" onclick={() => onEdit?.(article.id)} class="p-1 text-outline opacity-60 hover:text-primary hover:opacity-100" title="編集"><span class="material-symbols-outlined text-[18px]">edit</span></button><button type="button" onclick={() => deleteArticle(article.id)} class="p-1 text-outline opacity-60 hover:text-error hover:opacity-100" title="削除"><span class="material-symbols-outlined text-[18px]">delete</span></button></div>{/if}</div>
				</div>
				<h3 class="font-headline-lg text-[20px] leading-tight text-primary decoration-outline-variant decoration-1 underline-offset-4 transition-all group-hover:underline md:text-headline-lg">
					<a href={resolve(`/tech/${article.id}`)}>{article.title}</a>
				</h3>
				<p class="font-body-md text-body-md line-clamp-2 max-w-[640px] text-on-surface-variant">{article.content}</p>
				<a href={resolve(`/tech/${article.id}`)} class="font-label-md text-label-md text-primary hover:underline">続きを読む</a>
				<div class="mt-3 flex items-center gap-4">
					<LikeButton postId={article.id} initialLikesCount={article.likesCount} initialHasLiked={article.hasLiked} />
					<a href={resolve(`/tech/${article.id}`)} class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant hover:text-primary hover:underline" aria-label={`コメント${article.commentsCount}件を表示`}>
						<span class="material-symbols-outlined text-[16px]" aria-hidden="true">comment</span>{article.commentsCount}
					</a>
					{#if article.views}<span class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant"><span class="material-symbols-outlined text-[14px]">trending_up</span>{article.views}</span>{/if}
				</div>
			</article>
		{/each}
	</div>

	{#if !isAdmin && hasMoreArticles}
		<div class="flex flex-col items-center gap-3">
			{#if loadMoreError}<p class="font-body-sm text-body-sm text-error" role="alert">記事を追加で読み込めませんでした。</p>{/if}
			<button type="button" onclick={loadMore} disabled={isLoadingMore} class="font-label-md text-label-md min-h-11 rounded-lg border border-outline-variant/60 px-6 py-2 text-primary transition-colors hover:bg-surface-container-high disabled:cursor-wait disabled:opacity-60">
				{isLoadingMore ? '読み込み中…' : 'もっと見る'}
			</button>
		</div>
	{/if}

	<style>
		.category-scroller {
			scrollbar-width: none;
		}

		.category-scroller::-webkit-scrollbar {
			display: none;
		}
	</style>
</div>
