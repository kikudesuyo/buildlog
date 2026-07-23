<script lang="ts">
	import { goto } from '$app/navigation';
	import { updateDiary } from '$lib/api/client';

	let { data } = $props();

	let title = $state(data.diary.title);
	let content = $state(data.diary.content);
	let isSubmitting = $state(false);
	let errorMessage = $state('');

	// オートリサイズ用のアクション
	function autogrow(node: HTMLTextAreaElement) {
		function adjust() {
			node.style.height = 'auto';
			node.style.height = `${node.scrollHeight}px`;
		}
		adjust();
		node.addEventListener('input', adjust);
		return {
			destroy() {
				node.removeEventListener('input', adjust);
			}
		};
	}

	async function handleSave() {
		if (!title.trim() || !content.trim()) {
			errorMessage = 'タイトルと本文を入力してください。';
			return;
		}

		isSubmitting = true;
		errorMessage = '';
		try {
			await updateDiary(data.diary.id, title, content);
			goto('/');
		} catch (err) {
			errorMessage = 'つぶやきの更新に失敗しました。';
			isSubmitting = false;
		}
	}
</script>

<svelte:head>
	<title>つぶやきを編集 — Essence</title>
</svelte:head>

<div class="max-w-[720px] mx-auto px-gutter py-12 flex flex-col gap-8">
	<!-- エディタヘッダー -->
	<header class="flex items-center justify-between border-b border-outline-variant/10 pb-4">
		<button
			type="button"
			onclick={() => goto('/')}
			class="font-label-md text-label-md flex cursor-pointer items-center gap-1 text-outline hover:text-primary transition-colors"
		>
			<span class="material-symbols-outlined !text-[18px]">arrow_back</span>
			戻る
		</button>

		<div class="flex items-center gap-4">
			{#if errorMessage}
				<span class="text-error font-body-sm text-body-sm">{errorMessage}</span>
			{/if}
			<button
				type="button"
				onclick={handleSave}
				disabled={isSubmitting}
				class="font-label-md text-label-md cursor-pointer rounded-lg bg-primary px-6 py-2 text-on-primary hover:bg-primary/95 transition-all disabled:opacity-50"
			>
				{isSubmitting ? '保存中...' : '更新する'}
			</button>
		</div>
	</header>

	<!-- ライティングエリア -->
	<main class="flex flex-col gap-6">
		<input
			type="text"
			bind:value={title}
			placeholder="無題"
			class="w-full bg-transparent px-0 py-2 text-on-surface focus:outline-none text-[36px] font-display-lg leading-tight border-none"
			disabled={isSubmitting}
		/>

		<textarea
			use:autogrow
			bind:value={content}
			placeholder="ここにつぶやきを入力してください..."
			class="w-full bg-transparent px-0 py-2 text-on-surface focus:outline-none text-body-lg leading-relaxed border-none resize-none min-h-[300px]"
			disabled={isSubmitting}
		></textarea>
	</main>
</div>
