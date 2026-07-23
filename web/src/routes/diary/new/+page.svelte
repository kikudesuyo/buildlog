<script lang="ts">
	import { goto } from '$app/navigation';
	import { createDiary } from '$lib/api/client';

	let title = $state('');
	let content = $state('');
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
			await createDiary(title, content);
			goto('/');
		} catch (err) {
			errorMessage = 'つぶやきの保存に失敗しました。';
			isSubmitting = false;
		}
	}
</script>

<svelte:head>
	<title>新しいつぶやき — Essence</title>
</svelte:head>

<div class="max-w-[600px] mx-auto px-gutter py-8 flex flex-col gap-6">
	<!-- エディタヘッダー -->
	<header class="flex items-center justify-between border-b border-outline-variant/10 pb-3">
		<button
			type="button"
			onclick={() => goto('/')}
			class="font-label-sm text-label-sm flex cursor-pointer items-center gap-1 text-outline hover:text-primary transition-colors"
		>
			<span class="material-symbols-outlined !text-[16px]">arrow_back</span>
			戻る
		</button>

		<div class="flex items-center gap-3">
			{#if errorMessage}
				<span class="text-error font-body-sm text-body-sm">{errorMessage}</span>
			{/if}
			<button
				type="button"
				onclick={handleSave}
				disabled={isSubmitting}
				class="font-label-sm text-label-sm cursor-pointer rounded-lg bg-primary px-4 py-1.5 text-on-primary hover:bg-primary/95 transition-all disabled:opacity-50"
			>
				{isSubmitting ? '保存中...' : 'つぶやく'}
			</button>
		</div>
	</header>

	<!-- ライティングエリア -->
	<main class="flex flex-col gap-4">
		<input
			type="text"
			bind:value={title}
			placeholder="タイトルを入力..."
			class="w-full bg-transparent px-0 py-1 text-on-surface focus:outline-none text-2xl font-bold border-none"
			disabled={isSubmitting}
		/>

		<textarea
			use:autogrow
			bind:value={content}
			placeholder="いまどうしてる？"
			class="w-full bg-transparent px-0 py-1 text-on-surface focus:outline-none text-body-md leading-relaxed border-none resize-none min-h-[200px]"
			disabled={isSubmitting}
		></textarea>
	</main>
</div>
