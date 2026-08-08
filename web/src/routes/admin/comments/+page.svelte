<script lang="ts">
	import { deleteComment } from '$lib/api/client';
	import type { CommentEntry } from '$lib/api/types';

	let { data } = $props<{ data: { comments: CommentEntry[] } }>();
	let comments = $state<CommentEntry[]>(data.comments);
	let errorMessage = $state('');

	function formatDate(value: string) {
		return new Date(value).toLocaleString('ja-JP');
	}

	async function removeComment(comment: CommentEntry) {
		if (!confirm('このコメントを削除しますか？')) return;
		errorMessage = '';
		try {
			await deleteComment(comment.id);
			comments = comments.filter((item) => item.id !== comment.id);
		} catch {
			errorMessage = 'コメントを削除できませんでした。';
		}
	}
</script>

<svelte:head><title>Buildlog — コメント管理</title></svelte:head>

<section class="editorial-container mx-auto px-gutter" aria-labelledby="comments-heading">
	<header class="mb-8">
		<h1 id="comments-heading" class="font-display-lg text-display-lg text-primary">コメント管理</h1>
		<p class="font-body-md text-body-md mt-2 text-on-surface-variant">投稿されたコメントを確認・削除できます。</p>
	</header>

	{#if errorMessage}<p class="mb-4 rounded-lg border border-error/30 bg-error-container/30 p-3 text-body-sm text-on-error-container" role="alert">{errorMessage}</p>{/if}

	{#if comments.length === 0}
		<p class="rounded-xl border border-dashed border-outline-variant/30 py-12 text-center text-body-md text-outline">コメントはありません。</p>
	{:else}
		<div class="space-y-4">
			{#each comments as comment (comment.id)}
				<article class="rounded-xl border border-outline-variant/20 bg-surface-container-lowest p-5">
					<div class="mb-3 flex flex-wrap items-center justify-between gap-3">
						<div class="font-label-sm text-label-sm text-outline">コメント #{comment.id} · 投稿 #{comment.postId} · {formatDate(comment.createdAt)}</div>
						<button type="button" onclick={() => removeComment(comment)} class="min-h-11 rounded-lg border border-error/40 px-4 py-2 text-label-md text-error hover:bg-error-container/30">削除</button>
					</div>
					<p class="whitespace-pre-wrap text-body-md leading-relaxed text-on-surface-variant">{comment.content}</p>
				</article>
			{/each}
		</div>
	{/if}
</section>
