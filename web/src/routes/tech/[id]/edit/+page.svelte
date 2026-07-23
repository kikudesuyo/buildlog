<script lang="ts">
	import { goto } from '$app/navigation';
	import { updateTech } from '$lib/api/client';

	let { data } = $props();

	let title = $state(data.tech.title);
	let excerpt = $state(data.tech.excerpt);
	let category = $state(data.tech.category);
	let readTime = $state(data.tech.readTime);
	let views = $state(data.tech.views || '');
	let isNewsletter = $state(data.tech.isNewsletter || false);

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
		if (!title.trim() || !excerpt.trim() || !category || !readTime.trim()) {
			errorMessage = '必須項目を入力してください。';
			return;
		}

		isSubmitting = true;
		errorMessage = '';
		try {
			await updateTech(data.tech.id, {
				title,
				excerpt,
				category,
				readTime,
				views,
				isNewsletter
			});
			goto('/tech');
		} catch (err) {
			errorMessage = '技術記事の更新に失敗しました。';
			isSubmitting = false;
		}
	}
</script>

<svelte:head>
	<title>技術記事を編集 — Essence</title>
</svelte:head>

<div class="max-w-[720px] mx-auto px-gutter py-12 flex flex-col gap-8">
	<!-- エディタヘッダー -->
	<header class="flex items-center justify-between border-b border-outline-variant/10 pb-4">
		<button
			type="button"
			onclick={() => goto('/tech')}
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
			placeholder="記事タイトル"
			class="w-full bg-transparent px-0 py-2 text-on-surface focus:outline-none text-[36px] font-display-lg leading-tight border-none"
			disabled={isSubmitting}
		/>

		<!-- 設定グリッド (メタ情報) -->
		<section class="grid grid-cols-2 gap-4 rounded-xl bg-surface-container/50 border border-outline-variant/10 p-4">
			<div class="flex flex-col gap-1.5">
				<label for="tech-category" class="font-label-sm text-label-sm text-outline">カテゴリ *</label>
				<select
					id="tech-category"
					bind:value={category}
					disabled={isSubmitting}
					class="rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md"
				>
					<option value="Architecture">Architecture</option>
					<option value="Development">Development</option>
					<option value="Data Science">Data Science</option>
					<option value="Newsletter">Newsletter</option>
				</select>
			</div>

			<div class="flex flex-col gap-1.5">
				<label for="tech-readtime" class="font-label-sm text-label-sm text-outline">読了目安時間 *</label>
				<input
					id="tech-readtime"
					type="text"
					bind:value={readTime}
					placeholder="例: 5 min read"
					disabled={isSubmitting}
					class="rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md"
				/>
			</div>

			<div class="flex flex-col gap-1.5">
				<label for="tech-views" class="font-label-sm text-label-sm text-outline">閲覧数（任意）</label>
				<input
					id="tech-views"
					type="text"
					bind:value={views}
					disabled={isSubmitting}
					class="rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md"
				/>
			</div>

			<div class="flex items-center gap-2 mt-6 pl-1">
				<input
					id="tech-newsletter"
					type="checkbox"
					bind:checked={isNewsletter}
					disabled={isSubmitting}
					class="h-4 w-4 rounded border-outline-variant text-primary focus:ring-primary focus:ring-2 focus:ring-offset-0"
				/>
				<label for="tech-newsletter" class="font-label-md text-label-md text-on-surface cursor-pointer select-none">
					Weekly Pick (Card表示)
				</label>
			</div>
		</section>

		<div class="flex flex-col gap-1.5">
			<label for="tech-excerpt" class="font-label-sm text-label-sm text-outline">概要 / 本文 *</label>
			<textarea
				id="tech-excerpt"
				use:autogrow
				bind:value={excerpt}
				class="w-full bg-transparent px-0 py-2 text-on-surface focus:outline-none text-body-lg leading-relaxed border-none resize-none min-h-[250px]"
				disabled={isSubmitting}
			></textarea>
		</div>
	</main>
</div>
