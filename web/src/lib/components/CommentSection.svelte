<script lang="ts">
	import { onMount } from 'svelte';
	import { createComment, fetchComments } from '$lib/api/client';
	import type { CommentEntry } from '$lib/api/types';

	let { postId }: { postId: number } = $props();
	let commentList = $state<CommentEntry[]>([]);
	let newCommentContent = $state('');
	let replyTargetId = $state<number | null>(null);
	let replyContent = $state('');
	let isSubmitting = $state(false);
	let errorMessage = $state('');

	onMount(async () => {
		try {
			commentList = await fetchComments(fetch, postId);
		} catch {
			errorMessage = 'コメントを読み込めませんでした。';
		}
	});

	function formatDate(dateString: string) {
		const date = new Date(dateString);
		return `${date.getFullYear()}.${String(date.getMonth() + 1).padStart(2, '0')}.${String(date.getDate()).padStart(2, '0')} ${String(date.getHours()).padStart(2, '0')}:${String(date.getMinutes()).padStart(2, '0')}`;
	}

	function appendReply(parentId: number, reply: CommentEntry) {
		commentList = commentList.map((comment) =>
			comment.id === parentId ? { ...comment, replies: [...comment.replies, reply] } : comment
		);
	}

	async function submitComment(parentId: number | null) {
		const content = (parentId === null ? newCommentContent : replyContent).trim();
		if (!content || isSubmitting) return;

		isSubmitting = true;
		errorMessage = '';
		try {
			const comment = await createComment(postId, { parentId, content });
			if (parentId === null) {
				commentList = [...commentList, comment];
				newCommentContent = '';
			} else {
				appendReply(parentId, comment);
				replyTargetId = null;
				replyContent = '';
			}
		} catch {
			errorMessage = parentId === null ? 'コメントを投稿できませんでした。' : '返信を投稿できませんでした。';
		} finally {
			isSubmitting = false;
		}
	}
</script>

<section class="mt-16 border-t border-outline-variant/20 pt-10" aria-labelledby="comments-heading">
	<h2 id="comments-heading" class="font-headline-lg text-headline-lg mb-6 flex items-center gap-2 font-bold text-primary">
		<span class="material-symbols-outlined text-[22px]" aria-hidden="true">forum</span>
		コメント
	</h2>

	{#if errorMessage}
		<p class="mb-4 rounded-lg border border-error/30 bg-error-container/30 p-3 text-body-sm text-on-error-container" role="alert">
			{errorMessage}
		</p>
	{/if}

	<div class="mb-8 flex flex-col gap-4">
		{#if commentList.length === 0}
			<p class="rounded-xl border border-dashed border-outline-variant/30 py-8 text-center text-body-md text-outline">
				最初のコメントを投稿しましょう。
			</p>
		{:else}
			{#each commentList as comment (comment.id)}
				<article class="rounded-xl border border-outline-variant/15 bg-surface-container-lowest/30 p-4 shadow-2xs">
					<div class="mb-3 flex items-center justify-between gap-3">
						<time class="text-label-sm text-outline" datetime={comment.createdAt}>{formatDate(comment.createdAt)}</time>
						<button
							type="button"
							class="min-h-11 rounded-lg px-3 text-label-sm text-primary transition-colors hover:bg-primary/10 focus-visible:outline focus-visible:outline-2 focus-visible:outline-primary"
							onclick={() => {
								replyTargetId = replyTargetId === comment.id ? null : comment.id;
								replyContent = '';
							}}
						>
							<span class="material-symbols-outlined mr-1 align-middle text-[16px]" aria-hidden="true">reply</span>
							返信する
						</button>
					</div>

					<p class="whitespace-pre-wrap text-body-md leading-relaxed text-on-surface-variant">{comment.content}</p>

					{#if replyTargetId === comment.id}
						<form
							class="mt-4 flex flex-col gap-3 rounded-xl border border-outline-variant/20 bg-surface-container-low p-4"
							onsubmit={(event) => {
								event.preventDefault();
								void submitComment(comment.id);
							}}
						>
							<label for={`reply-${comment.id}`} class="text-label-md font-bold text-on-surface">返信内容</label>
							<textarea
								id={`reply-${comment.id}`}
								bind:value={replyContent}
								class="min-h-24 resize-y rounded-lg border border-outline-variant bg-surface-container-lowest px-3 py-2 text-body-md text-on-surface focus-visible:outline focus-visible:outline-2 focus-visible:outline-primary"
								placeholder="返信内容を入力してください"
								disabled={isSubmitting}
							></textarea>
							<div class="flex justify-end">
								<button
									type="submit"
									class="min-h-11 rounded-lg bg-primary px-4 text-label-md text-on-primary transition-colors hover:bg-primary/90 focus-visible:outline focus-visible:outline-2 focus-visible:outline-primary disabled:opacity-50"
									disabled={isSubmitting || !replyContent.trim()}
								>
									{isSubmitting ? '送信中...' : '返信を送信'}
								</button>
							</div>
						</form>
					{/if}

					{#if comment.replies.length > 0}
						<div class="mt-4 flex flex-col gap-3 border-l-2 border-outline-variant/30 pl-4" aria-label="返信">
							{#each comment.replies as reply (reply.id)}
								<div class="rounded-lg bg-surface-container-low p-3">
									<time class="mb-2 block text-label-sm text-outline" datetime={reply.createdAt}>{formatDate(reply.createdAt)}</time>
									<p class="whitespace-pre-wrap text-body-md leading-relaxed text-on-surface-variant">{reply.content}</p>
								</div>
							{/each}
						</div>
					{/if}
				</article>
			{/each}
		{/if}
	</div>

	<form
		class="flex flex-col gap-4 rounded-xl border border-outline-variant/20 bg-surface-container-low p-5"
		onsubmit={(event) => {
			event.preventDefault();
			void submitComment(null);
		}}
	>
		<h3 class="font-headline-sm text-headline-sm flex items-center gap-2 font-bold text-primary">
			<span class="material-symbols-outlined text-[20px]" aria-hidden="true">add_comment</span>
			コメントを書く
		</h3>
		<label for="comment-content" class="text-label-md font-bold text-on-surface">コメント内容</label>
		<textarea
			id="comment-content"
			bind:value={newCommentContent}
			class="min-h-28 resize-y rounded-lg border border-outline-variant bg-surface-container-lowest px-3 py-2 text-body-md text-on-surface focus-visible:outline focus-visible:outline-2 focus-visible:outline-primary"
			placeholder="コメント内容を入力してください"
			disabled={isSubmitting}
		></textarea>
		<div class="flex justify-end">
			<button
				type="submit"
				class="min-h-11 rounded-lg bg-primary px-5 text-label-md text-on-primary transition-colors hover:bg-primary/90 focus-visible:outline focus-visible:outline-2 focus-visible:outline-primary disabled:opacity-50"
				disabled={isSubmitting || !newCommentContent.trim()}
			>
				{isSubmitting ? '送信中...' : 'コメントを送信'}
			</button>
		</div>
	</form>
</section>
