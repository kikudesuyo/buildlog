<script lang="ts">
	import { onMount } from 'svelte';
	import { likePost, unlikePost } from '$lib/api/client';

	interface Props {
		postId: number;
		initialLikesCount: number;
		initialHasLiked?: boolean;
	}

	let { postId, initialLikesCount, initialHasLiked = false }: Props = $props();

	let likesCount = $state(initialLikesCount);
	let hasLiked = $state(initialHasLiked);
	let isAnimating = $state(false);

	const STORAGE_KEY = 'buildlog_liked_posts';

	onMount(() => {
		try {
			const stored = localStorage.getItem(STORAGE_KEY);
			const likedIds = stored ? JSON.parse(stored) as number[] : [];
			if (likedIds.includes(postId)) {
				hasLiked = true;
			}
		} catch (e) {
			// ignore
		}
	});

	async function toggleLike() {
		if (isAnimating) return;
		isAnimating = true;

		const previousHasLiked = hasLiked;
		const previousCount = likesCount;

		hasLiked = !hasLiked;
		likesCount = hasLiked ? likesCount + 1 : Math.max(0, likesCount - 1);

		try {
			const stored = localStorage.getItem(STORAGE_KEY);
			let likedIds = stored ? JSON.parse(stored) as number[] : [];
			if (hasLiked) {
				if (!likedIds.includes(postId)) likedIds.push(postId);
			} else {
				likedIds = likedIds.filter(id => id !== postId);
			}
			localStorage.setItem(STORAGE_KEY, JSON.stringify(likedIds));
		} catch (e) {
			// ignore
		}

		try {
			if (hasLiked) {
				const res = await likePost(postId);
				likesCount = res.likes_count;
				hasLiked = res.has_liked;
			} else {
				const res = await unlikePost(postId);
				likesCount = res.likes_count;
				hasLiked = res.has_liked;
			}
		} catch (e) {
			hasLiked = previousHasLiked;
			likesCount = previousCount;
		} finally {
			setTimeout(() => {
				isAnimating = false;
			}, 300);
		}
	}
</script>

<button
	type="button"
	onclick={toggleLike}
	class="inline-flex items-center gap-1.5 rounded-full px-3 py-1.5 text-label-md font-label-md transition-all duration-200 border cursor-pointer select-none
		{hasLiked 
			? 'bg-rose-50 border-rose-200/50 text-rose-600 hover:bg-rose-100/70' 
			: 'bg-surface-container-low border-outline-variant/10 text-on-surface-variant hover:bg-surface-container hover:text-primary'}"
	aria-label="いいねボタン"
>
	<span 
		class="material-symbols-outlined transition-transform duration-300 text-[18px] 
			{hasLiked ? 'fill-1 text-rose-600 scale-110' : 'text-outline'} 
			{isAnimating ? 'animate-heart-pop' : ''}"
	>
		favorite
	</span>
	<span class="font-medium tabular-nums">{likesCount}</span>
</button>

<style>
	:global(.material-symbols-outlined.fill-1) {
		font-variation-settings: 'FILL' 1, 'wght' 400, 'GRAD' 0, 'opsz' 24;
	}

	@keyframes heart-pop {
		0% {
			transform: scale(1);
		}
		50% {
			transform: scale(1.4);
		}
		100% {
			transform: scale(1.1);
		}
	}

	.animate-heart-pop {
		animation: heart-pop 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275) forwards;
	}
</style>
