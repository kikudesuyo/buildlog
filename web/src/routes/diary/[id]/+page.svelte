<script lang="ts">
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import LikeButton from '$lib/components/LikeButton.svelte';
	import CommentSection from '$lib/components/CommentSection.svelte';

	let { data } = $props();
	const returnPath = $derived(page.url.searchParams.get('from') === 'calendar' ? '/profile#post-history' : '/');
	const returnLabel = $derived(page.url.searchParams.get('from') === 'calendar' ? '投稿履歴とカレンダーへ戻る' : '日々のつぶやきへ戻る');

	function formatDate(dateStr: string) {
		if (!dateStr) return '';
		const date = new Date(dateStr);
		return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`;
	}
</script>

<svelte:head>
	<title>{data.diary.title} — Buildlog</title>
	<meta name="description" content={data.diary.content ? data.diary.content.substring(0, 120) + '...' : ''} />
</svelte:head>

<div class="editorial-container mx-auto max-w-[800px] px-gutter pb-20">
	<a
		href={resolve(returnPath)}
		class="font-label-md text-label-md mb-8 inline-flex items-center gap-2 text-outline transition-colors duration-200 hover:text-primary"
	>
		<span class="material-symbols-outlined text-[18px]">arrow_back</span>
		{returnLabel}
	</a>

	<article class="flex flex-col gap-6">
		<header class="flex flex-col gap-4 border-b border-outline-variant/20 pb-8">
			<span class="font-label-sm text-label-sm text-on-surface-variant">{formatDate(data.diary.createdAt)}</span>
			<h1 class="font-display-lg text-[clamp(1.75rem,7vw,2.5rem)] leading-tight font-bold tracking-tight text-primary md:text-[40px]">
				{data.diary.title}
			</h1>
		</header>

		<p class="font-body-md text-body-md whitespace-pre-wrap break-words pt-4 text-[1rem] leading-7 text-on-surface md:text-body-md md:leading-relaxed">
			{data.diary.content}
		</p>

		<div class="flex items-center gap-4 border-t border-outline-variant/10 pt-6">
			<LikeButton postId={data.diary.id} initialLikesCount={data.diary.likesCount} initialHasLiked={data.diary.hasLiked} />
		</div>
		<CommentSection postId={data.diary.id} />
	</article>
</div>
