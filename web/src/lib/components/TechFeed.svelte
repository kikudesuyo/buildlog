<script lang="ts">
	import { onMount } from 'svelte';
	import type { TechArticle } from '$lib/api/types';
	import { invalidateAll, replaceState } from '$app/navigation';
	import { page } from '$app/stores';
	import { resolve } from '$app/paths';
	import { fetchTechFeed, type ApiFetch, type TechSort, type TechSortOrder } from '$lib/api/client';
	import Button from './Button.svelte';

	type Props = {
		techArticles?: TechArticle[];
		hasMore?: boolean;
		loadError?: boolean;
		isAdmin?: boolean;
	};

	let { techArticles = [], hasMore = false, loadError = false, isAdmin = false }: Props = $props();
	const initialState = { techArticles, hasMore };
	let loadedTechArticles = $state(initialState.techArticles);
	let hasMoreArticles = $state(initialState.hasMore);
	let isLoadingMore = $state(false);
	let loadMoreError = $state(false);
	let sort = $state<TechSort>($page.url.searchParams.get('sort') === 'likes' ? 'likes' : 'newest');
	let sortOrder = $state<TechSortOrder>($page.url.searchParams.get('order') === 'asc' ? 'asc' : 'desc');
	let isDesktop = $state(false);
	const desktopInitialLimit = 8;
	const mobileLoadMoreLimit = 3;
	const desktopLoadMoreLimit = 6;
	const storageKey = 'tech-feed-count';
	let isRestoring = $state(false);
	let articles = $derived(loadedTechArticles);

	function formatDate(dateStr: string) {
		if (!dateStr) return '';
		const date = new Date(dateStr);
		return `${date.getFullYear()}.${String(date.getMonth() + 1).padStart(2, '0')}.${String(date.getDate()).padStart(2, '0')}`;
	}

	function articleHref(article: TechArticle) {
		return article.external?.url ?? '/tech';
	}

	async function loadMore() {
		if (isLoadingMore || !hasMoreArticles) return;
		isLoadingMore = true;
		loadMoreError = false;
		try {
			const offset = loadedTechArticles.length;
			const limit = isDesktop ? desktopLoadMoreLimit : mobileLoadMoreLimit;
			const next = await fetchTechFeed(window.fetch.bind(window) as ApiFetch, false, offset, limit, sortOrder, sort);
			loadedTechArticles = [...loadedTechArticles, ...next.techArticles];
			hasMoreArticles = next.hasMore;
			if (!isAdmin) sessionStorage.setItem(storageKey, String(articles.length));
		} catch {
			loadMoreError = true;
		} finally {
			isLoadingMore = false;
		}
	}

	async function loadDesktopInitial() {
		if (!isDesktop || !hasMoreArticles) return;
		isLoadingMore = true;
		loadMoreError = false;
		try {
			const next = await fetchTechFeed(window.fetch.bind(window) as ApiFetch, false, 0, desktopInitialLimit, sortOrder, sort);
			loadedTechArticles = next.techArticles;
			hasMoreArticles = next.hasMore;
		} catch {
			loadMoreError = true;
		} finally {
			isLoadingMore = false;
		}
	}

	async function handleSortChange(value: string) {
		const nextSort: TechSort = value === 'likes' ? 'likes' : 'newest';
		const nextOrder: TechSortOrder = value === 'oldest' ? 'asc' : 'desc';
		if (sort === nextSort && sortOrder === nextOrder) return;

		sort = nextSort;
		sortOrder = nextOrder;
		const nextUrl = new URL($page.url);
		if (nextSort === 'newest') nextUrl.searchParams.delete('sort');
		else nextUrl.searchParams.set('sort', nextSort);
		if (nextOrder === 'desc') nextUrl.searchParams.delete('order');
		else nextUrl.searchParams.set('order', nextOrder);
		// eslint-disable-next-line svelte/no-navigation-without-resolve
		await replaceState(`${resolve('/tech')}${nextUrl.search}`, {});
		isLoadingMore = true;
		loadMoreError = false;
		try {
			const limit = isDesktop ? desktopInitialLimit : mobileLoadMoreLimit;
			const next = await fetchTechFeed(window.fetch.bind(window) as ApiFetch, false, 0, limit, nextOrder, nextSort);
			loadedTechArticles = next.techArticles;
			hasMoreArticles = next.hasMore;
			if (!isAdmin) sessionStorage.setItem(storageKey, String(articles.length));
		} catch {
			loadMoreError = true;
		} finally {
			isLoadingMore = false;
		}
	}

	onMount(() => {
		void (async () => {
			isDesktop = window.matchMedia('(min-width: 768px)').matches;
			await loadDesktopInitial();
			const savedCount = Number(sessionStorage.getItem(storageKey));
			if (!isAdmin && savedCount > articles.length) {
				isRestoring = true;
				while (articles.length < savedCount && hasMoreArticles && !loadMoreError) await loadMore();
				isRestoring = false;
			}
		})();
		return () => undefined;
	});
</script>

<div class="editorial-container mx-auto flex flex-col gap-8 px-gutter md:!max-w-6xl">
	<section class="flex flex-wrap items-start justify-between gap-4">
		<div>
			{#if isAdmin}
				<p class="font-label-sm text-label-sm mb-stack-sm tracking-[0.2em] text-outline uppercase">Content Manager / Tech</p>
			{/if}
			<h1 class="font-display-lg text-display-lg mb-stack-sm text-primary">{isAdmin ? '技術記事一覧' : '技術ブログ'}</h1>
		</div>
		<label class="font-label-sm text-label-sm flex items-center gap-2 text-on-surface-variant">
			<span>並び順</span>
			<select
				value={sort === 'likes' ? 'likes' : sortOrder === 'asc' ? 'oldest' : 'newest'}
				onchange={(event) => handleSortChange(event.currentTarget.value)}
				class="font-label-sm text-label-sm min-h-11 rounded-lg border border-outline-variant/40 bg-surface-container-lowest px-3 py-2 text-primary"
				aria-label="技術記事の並び順"
			>
				<option value="newest">新しい順</option>
				<option value="oldest">古い順</option>
				<option value="likes">いいね順</option>
			</select>
		</label>
	</section>

	{#if loadError}
		<section class="rounded-xl border border-error/30 bg-error-container/30 p-8 text-center" role="alert">
			<h2 class="font-headline-md text-headline-md text-on-surface">記事を読み込めませんでした</h2>
			<p class="font-body-md text-body-md mt-2 text-on-surface-variant">通信を確認して、もう一度お試しください。</p>
			<Button type="button" class="mt-5" onclick={() => invalidateAll()}>再試行</Button>
		</section>
	{:else if articles.length === 0}
		<section class="rounded-xl border border-outline-variant/30 bg-surface-container-low p-8 text-center">
			<h2 class="font-headline-md text-headline-md text-primary">技術記事はまだありません</h2>
			<p class="font-body-md text-body-md mt-2 text-on-surface-variant">新しい記事が公開されるまでお待ちください。</p>
		</section>
	{/if}

	<div class="grid gap-6 md:grid-cols-2 md:gap-8">
		{#each articles as article (article.key)}
			<article class="group relative flex flex-col gap-3 rounded-xl border border-outline-variant/20 bg-surface-container-lowest p-4 shadow-2xs transition-all duration-300 hover:shadow-md hover:border-primary/20 md:p-6">
				<div class="flex items-center justify-between">
					<div class="flex items-center gap-stack-sm">
						{#if article.external}<span class="font-label-sm text-label-sm rounded bg-secondary-container px-2 py-0.5 text-on-secondary-container">{article.external.provider}</span>{/if}
						{#if article.status === 'draft'}
							<span class="font-label-sm text-label-sm px-2 py-0.5 rounded bg-outline-variant/40 text-on-surface-variant">下書き</span>
						{/if}
					</div>
					<div class="flex items-center gap-4"><span class="font-label-sm text-label-sm text-on-surface-variant">{formatDate(article.createdAt)}</span></div>
				</div>
				<h3 class="font-headline-lg text-[20px] leading-tight text-primary decoration-outline-variant decoration-1 underline-offset-4 transition-all group-hover:underline max-[640px]:!text-[20px] md:!text-[20px]">
				<!-- eslint-disable-next-line svelte/no-navigation-without-resolve -->
				<a href={article.external ? articleHref(article) : resolve('/tech')} target={article.external ? '_blank' : undefined} rel={article.external ? 'noreferrer' : undefined}>{article.title}</a>
				</h3>
				<p class="font-body-md text-body-md line-clamp-2 max-w-[640px] text-on-surface-variant max-[640px]:!text-[14px]">{article.content}</p>
				{#if article.external?.thumbnailUrl}<img src={article.external.thumbnailUrl} alt="" class="aspect-[2/1] w-full rounded-lg object-cover" loading="lazy" />{/if}
			<div class="flex flex-wrap items-center justify-between gap-3">
				<!-- eslint-disable-next-line svelte/no-navigation-without-resolve -->
				<a href={article.external ? articleHref(article) : resolve('/tech')} target={article.external ? '_blank' : undefined} rel={article.external ? 'noreferrer' : undefined} class="font-label-md text-label-md text-primary hover:underline">続きを読む</a>
				{#if article.external}
					<span class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant" aria-label={`Qiitaのいいね数: ${article.likesCount}件`}>
						<span class="material-symbols-outlined fill-1 text-[16px] text-error" aria-hidden="true">favorite</span>
						Qiita いいね {article.likesCount}
					</span>
				{/if}
			</div>
			</article>
		{/each}
	</div>

	{#if !isAdmin && hasMoreArticles && !isRestoring}
		<div class="flex flex-col items-center gap-3">
			{#if loadMoreError}<p class="font-body-sm text-body-sm text-error" role="alert">記事を追加で読み込めませんでした。</p>{/if}
			<Button type="button" variant="outline" onclick={loadMore} disabled={isLoadingMore} class="border-outline-variant/60 px-6 hover:bg-surface-container-high hover:text-primary">
				{isLoadingMore ? '読み込み中…' : 'もっと見る'}
			</Button>
		</div>
	{/if}

</div>
