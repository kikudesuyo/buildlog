<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { createTech } from '$lib/api/client';
	import UnsavedChangesGuard from '$lib/components/UnsavedChangesGuard.svelte';
	import { defaultTechCategory, techCategories } from '$lib/tech/categories';

	let title = $state('');
	let content = $state('');
	let category = $state(defaultTechCategory);
	let views = $state('');

	let isSubmitting = $state(false);
	let errorMessage = $state('');
	let isDirty = $derived(
		title.trim().length > 0 ||
		content.trim().length > 0 ||
		category !== defaultTechCategory ||
		views.trim().length > 0
	);

	// ダミーのUIステート (image.pngの再現用)
	let tags = $state(['技術', 'プログラミング']);
	let isCommentsAllowed = $state(true);

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

	async function handleSave(status: 'draft' | 'published') {
		if (!title.trim() || !content.trim() || !category) {
			errorMessage = '必須項目（タイトル、本文、カテゴリ）を入力してください。';
			return;
		}

		isSubmitting = true;
		errorMessage = '';
		try {
			await createTech({
				title,
				content,
				category,
				views,
				status
			});
			goto(resolve('/admin/tech'));
		} catch {
			errorMessage = '技術記事の保存に失敗しました。';
			isSubmitting = false;
		}
	}

	function removeTag(tagToRemove: string) {
		tags = tags.filter(t => t !== tagToRemove);
	}

	function addTag() {
		const newTag = prompt('タグ名を入力してください：');
		if (newTag && newTag.trim()) {
			tags = [...tags, newTag.trim()];
		}
	}
</script>

<UnsavedChangesGuard {isDirty} {isSubmitting} />

<svelte:head>
	<title>新規技術記事作成 — Buildlog</title>
</svelte:head>

<!-- ヘッダー（全幅） -->
<header class="fixed top-0 left-0 w-full h-16 bg-white border-b border-outline-variant/20 px-gutter flex items-center justify-between z-50">
	<div class="flex items-center gap-3">
		<a href={resolve('/admin/tech')} class="text-headline-md font-headline-md text-primary font-bold tracking-tight">
			Buildlog
		</a>
		<span class="h-4 w-px bg-outline-variant/30"></span>
		<span class="text-outline font-label-md text-label-md">Drafts</span>
	</div>

	<div class="flex items-center gap-6">
		{#if errorMessage}
			<span class="text-error font-body-sm text-body-sm">{errorMessage}</span>
		{/if}
		<button
			type="button"
			onclick={() => handleSave('draft')}
			disabled={isSubmitting}
			class="text-outline font-label-md text-label-md hover:text-primary transition-colors cursor-pointer"
		>
			下書き保存
		</button>
		<button
			type="button"
			onclick={() => handleSave('published')}
			disabled={isSubmitting}
			class="bg-primary text-on-primary font-label-md text-label-md px-5 py-2 rounded-lg font-medium hover:bg-primary/95 transition-colors cursor-pointer disabled:opacity-50"
		>
			{isSubmitting ? '公開中...' : '公開する'}
		</button>
		<img
			src="https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=100&q=80"
			alt="Profile"
			class="h-8 w-8 rounded-full object-cover border border-outline-variant/20"
		/>
	</div>
</header>

<!-- エディタ本体 -->
<div class="editorial-container mx-auto px-gutter pt-24 pb-20 relative">
	
	<!-- 左フローティングツールバー (絶対配置) -->
	<aside class="absolute -left-12 top-24 hidden md:flex flex-col items-center gap-1.5 bg-white border border-outline-variant/20 rounded-xl p-1.5 shadow-xs w-11">
		<button type="button" class="w-8 h-8 flex items-center justify-center font-bold text-outline hover:text-primary transition-colors rounded hover:bg-surface-container" title="太字">B</button>
		<button type="button" class="w-8 h-8 flex items-center justify-center italic text-outline hover:text-primary transition-colors rounded hover:bg-surface-container" title="斜体">I</button>
		<button type="button" class="w-8 h-8 flex items-center justify-center text-outline hover:text-primary transition-colors rounded hover:bg-surface-container" title="リスト">
			<span class="material-symbols-outlined text-[18px]">format_list_bulleted</span>
		</button>
		<button type="button" class="w-8 h-8 flex items-center justify-center text-outline hover:text-primary transition-colors rounded hover:bg-surface-container" title="リンク">
			<span class="material-symbols-outlined text-[18px]">link</span>
		</button>
	</aside>

	<main class="flex flex-col gap-6">
		<!-- タイトル -->
		<div class="border-b border-outline-variant/10 pb-4 mb-4">
			<input
				type="text"
				bind:value={title}
				placeholder="タイトルを入力..."
				class="w-full bg-transparent px-0 py-1 text-on-surface focus:outline-none text-[36px] font-bold tracking-tight border-none placeholder:text-outline-variant/50"
				disabled={isSubmitting}
			/>
		</div>

		<!-- 本文 -->
		<div class="flex flex-col gap-1.5">
			<label for="tech-content" class="font-label-md text-label-md font-bold text-on-surface">本文 *</label>
			<textarea
				id="tech-content"
				use:autogrow
				bind:value={content}
				placeholder="本文を書き始めましょう..."
				class="w-full bg-transparent px-0 py-1 text-on-surface focus:outline-none text-body-lg leading-relaxed border-none resize-none min-h-[300px] placeholder:text-outline-variant/50"
				disabled={isSubmitting}
			></textarea>
		</div>

		<!-- 下部設定セクション -->
		<footer class="border-t border-outline-variant/10 pt-8 mt-12 grid grid-cols-1 md:grid-cols-2 gap-8">
			<!-- 左カラム: タグ・カテゴリ設定 -->
			<div class="flex flex-col gap-5">
				<!-- カテゴリ選択 -->
				<div class="flex flex-col gap-1.5">
					<label for="tech-category" class="font-label-md text-label-md font-bold text-on-surface">カテゴリ *</label>
					<select
						id="tech-category"
						bind:value={category}
						disabled={isSubmitting}
						class="w-full rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md cursor-pointer"
					>
						{#each techCategories as techCategory (techCategory)}
							<option value={techCategory}>{techCategory}</option>
						{/each}
					</select>
				</div>

				<!-- タグ設定 -->
				<div class="flex flex-col gap-3 mt-2">
					<h3 class="font-label-md text-label-md font-bold text-on-surface">タグ設定</h3>
					<div class="flex flex-wrap gap-2 items-center">
						{#each tags as tag (tag)}
							<span class="inline-flex items-center gap-1 bg-surface-container-high px-3 py-1 rounded text-body-sm text-on-surface-variant">
								{tag}
								<button type="button" onclick={() => removeTag(tag)} class="hover:text-error transition-colors cursor-pointer font-bold text-[10px]">×</button>
							</span>
						{/each}
						<button
							type="button"
							onclick={addTag}
							class="border border-dashed border-outline-variant/60 hover:border-primary px-3 py-1 rounded text-body-sm text-outline hover:text-primary transition-all cursor-pointer"
						>
							+ タグを追加
						</button>
					</div>
				</div>
			</div>

			<!-- 右カラム: 記事設定 & 公開設定 -->
			<div class="flex flex-col gap-5">
				<h3 class="font-label-md text-label-md font-bold text-on-surface">記事設定</h3>
				
				<div class="grid grid-cols-1 gap-4">
					<div class="flex flex-col gap-1.5">
						<label for="tech-views" class="font-label-xs text-[11px] text-outline">閲覧数（任意）</label>
						<input
							id="tech-views"
							type="text"
							bind:value={views}
							placeholder="例: 1,024 views"
							disabled={isSubmitting}
							class="rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md"
						/>
					</div>
				</div>

				<div class="flex flex-col gap-4 mt-2 border-t border-outline-variant/10 pt-4">
					<div class="flex items-center justify-between">
						<span class="text-body-md text-on-surface-variant">コメントを許可する</span>
						<button
							type="button"
							aria-label="コメント許可トグル"
							onclick={() => (isCommentsAllowed = !isCommentsAllowed)}
							class="w-10 h-6 rounded-full p-0.5 transition-colors relative flex items-center cursor-pointer {isCommentsAllowed ? 'bg-primary' : 'bg-outline-variant/40'}"
						>
							<div class="w-5 h-5 bg-white rounded-full shadow transition-transform duration-200 {isCommentsAllowed ? 'translate-x-4' : 'translate-x-0'}"></div>
						</button>
					</div>
				</div>
			</div>
		</footer>
	</main>
</div>
