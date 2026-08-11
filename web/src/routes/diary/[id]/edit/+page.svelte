<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { updateDiary } from '$lib/api/client';
	import UnsavedChangesGuard from '$lib/components/UnsavedChangesGuard.svelte';

	let { data } = $props();

	let title = $state(data.diary.title);
	const MAX_TITLE_LENGTH = 100;
	let contentElement = $state<HTMLTextAreaElement | null>(null);
	let content = $state(data.diary.content);
	let status = $state(data.diary.status || 'draft');
	let isSubmitting = $state(false);
	let errorMessage = $state('');
	let isDirty = $derived(
		title !== data.diary.title ||
		content !== data.diary.content ||
		status !== (data.diary.status || 'draft')
	);

	// オートリサイズ用のアクション
	function autogrow(node: HTMLTextAreaElement) {
		function adjust() {
			const scrollY = window.scrollY;
			node.style.height = '0px';
			node.style.height = `${node.scrollHeight}px`;
			window.scrollTo(0, scrollY);
		}
		adjust();
		node.addEventListener('input', adjust);
		return {
			destroy() {
				node.removeEventListener('input', adjust);
			}
		};
	}

	async function handleSave(statusVal: 'draft' | 'published') {
		if (!title.trim() || !content.trim()) {
			errorMessage = 'タイトルと本文を入力してください。';
			return;
		}

		isSubmitting = true;
		errorMessage = '';
		try {
			await updateDiary(data.diary.id, title, content, statusVal);
			goto(resolve('/admin'));
		} catch {
			errorMessage = 'つぶやきの更新に失敗しました。';
			isSubmitting = false;
		}
	}

	function applyMarkdown(prefix: string, suffix = prefix) {
		if (!contentElement) return;
		const start = contentElement.selectionStart;
		const end = contentElement.selectionEnd;
		content = `${content.slice(0, start)}${prefix}${content.slice(start, end)}${suffix}${content.slice(end)}`;
		requestAnimationFrame(() => {
			contentElement?.focus();
			contentElement?.setSelectionRange(start + prefix.length, end + prefix.length);
		});
	}

	function insertLink() {
		const url = window.prompt('リンクURLを入力してください：');
		if (url?.trim()) applyMarkdown('[', `](${url.trim()})`);
	}

</script>

<UnsavedChangesGuard {isDirty} {isSubmitting} />

<svelte:head>
	<title>つぶやきを編集 — Buildlog</title>
</svelte:head>

{#if errorMessage}
	<div class="mx-auto max-w-container-max px-gutter text-body-sm text-error md:hidden" role="alert">{errorMessage}</div>
{/if}
<nav class="fixed inset-x-0 bottom-0 z-50 flex gap-2 border-t border-outline-variant/20 bg-white/95 p-3 pb-[max(0.75rem,env(safe-area-inset-bottom))] backdrop-blur md:hidden" aria-label="編集アクション">
	<button type="button" onclick={() => handleSave('draft')} disabled={isSubmitting} class="min-h-11 flex-1 rounded-lg border border-outline-variant/40 px-3 py-2 font-label-md text-outline disabled:opacity-50">下書き保存</button>
	<button type="button" onclick={() => handleSave('published')} disabled={isSubmitting} class="min-h-11 flex-1 rounded-lg bg-primary px-3 py-2 font-label-md text-on-primary disabled:opacity-50">{isSubmitting ? '更新中...' : '更新する'}</button>
</nav>

<!-- エディタ本体 -->
<div class="editorial-container relative mx-auto px-gutter pb-36 pt-8 md:pb-20">
	
	<!-- 左フローティングツールバー (絶対配置) -->
	<aside class="absolute -left-12 top-24 hidden md:flex flex-col items-center gap-1.5 bg-surface-container-lowest border border-outline-variant/20 rounded-xl p-1.5 shadow-xs w-11">
		<button type="button" class="w-8 h-8 flex items-center justify-center font-bold text-outline hover:text-primary transition-colors rounded hover:bg-surface-container" onclick={() => applyMarkdown('**')} title="太字">B</button>
		<button type="button" class="w-8 h-8 flex items-center justify-center italic text-outline hover:text-primary transition-colors rounded hover:bg-surface-container" onclick={() => applyMarkdown('*')} title="斜体">I</button>
		<button type="button" class="w-8 h-8 flex items-center justify-center text-outline hover:text-primary transition-colors rounded hover:bg-surface-container" onclick={() => applyMarkdown('- ' , '')} title="リスト">
			<span class="material-symbols-outlined text-[18px]">format_list_bulleted</span>
		</button>
		<button type="button" class="w-8 h-8 flex items-center justify-center text-outline hover:text-primary transition-colors rounded hover:bg-surface-container" onclick={insertLink} title="リンク">
			<span class="material-symbols-outlined text-[18px]">link</span>
		</button>
	</aside>

	<main class="flex flex-col gap-6">
		<div class="hidden items-center justify-end gap-3 border-b border-outline-variant/20 pb-4 md:flex">
			<button type="button" onclick={() => handleSave('draft')} disabled={isSubmitting} class="font-label-md text-label-md text-outline hover:text-primary">下書き保存</button>
			<button type="button" onclick={() => handleSave('published')} disabled={isSubmitting} class="font-label-md text-label-md rounded-lg bg-primary px-5 py-2 text-on-primary disabled:opacity-50">{isSubmitting ? '更新中...' : '更新する'}</button>
		</div>
		<nav class="sticky top-20 z-10 flex gap-2 rounded-lg border border-outline-variant/20 bg-surface-container-lowest/95 p-2 backdrop-blur md:hidden" aria-label="編集セクション">
			<a href="#editor-body" class="min-h-11 flex-1 rounded px-3 py-2 text-center font-label-md text-label-md text-primary">本文</a>
			<a href="#editor-settings" class="min-h-11 flex-1 rounded px-3 py-2 text-center font-label-md text-label-md text-primary">設定</a>
		</nav>
		<!-- タイトル -->
		<div class="border-b border-outline-variant/10 pb-4 mb-4">
			<textarea
				use:autogrow
				bind:value={title}
				maxlength={MAX_TITLE_LENGTH}
				rows="2"
				aria-describedby="title-count"
				placeholder="タイトルを入力..."
				class="min-h-20 w-full resize-none overflow-hidden bg-transparent px-0 py-1 text-[clamp(1.5rem,6vw,2rem)] font-bold leading-tight tracking-tight text-on-surface focus:outline-none md:text-[28px]"
				disabled={isSubmitting}
			></textarea>
			<p id="title-count" class="text-right text-body-sm text-outline">{title.length}/{MAX_TITLE_LENGTH}</p>
		</div>

		<!-- 本文 -->
		<div id="editor-body">
		<textarea
			use:autogrow
			bind:this={contentElement}
			bind:value={content}
			placeholder="物語を書き始めましょう..."
			class="min-h-[300px] w-full resize-none border-none bg-transparent px-0 py-1 text-body-md leading-relaxed text-on-surface focus:outline-none placeholder:text-outline-variant/50"
			disabled={isSubmitting}
		></textarea>
		</div>

	</main>
</div>
