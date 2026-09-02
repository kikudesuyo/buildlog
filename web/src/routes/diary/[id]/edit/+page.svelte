<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { updateDiary } from '$lib/api/client';
	import UnsavedChangesGuard from '$lib/components/UnsavedChangesGuard.svelte';

	let { data } = $props();
	const initialData = data;

	let title = $state(initialData.diary.title);
	const MAX_TITLE_LENGTH = 100;
	let contentElement = $state<HTMLTextAreaElement | null>(null);
	let content = $state(initialData.diary.content);
	let status = $state(initialData.diary.status || 'draft');
	let isSubmitting = $state(false);
	let errorMessage = $state('');
	let isAutoSaving = $state(false);
	let autoSaveError = $state(false);
	let isDirty = $derived(
		title !== data.diary.title ||
		content !== data.diary.content ||
		status !== (data.diary.status || 'draft')
	);

	async function autoSaveDraft(titleSnapshot: string, contentSnapshot: string, statusSnapshot: 'draft' | 'published') {
		if (!titleSnapshot.trim() || !contentSnapshot.trim() || isSubmitting || isAutoSaving) return;
		isAutoSaving = true;
		autoSaveError = false;
		try {
			await updateDiary(data.diary.id, titleSnapshot, contentSnapshot, statusSnapshot);
		} catch {
			autoSaveError = true;
		} finally {
			isAutoSaving = false;
		}
	}

	$effect(() => {
		const currentTitle = title;
		const currentContent = content;
		const currentStatus = status;
		if (!currentTitle.trim() || !currentContent.trim() || isSubmitting) return;
		const timer = setTimeout(() => void autoSaveDraft(currentTitle, currentContent, currentStatus), 5000);
		return () => clearTimeout(timer);
	});

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
<!-- エディタ本体 -->
<div class="editorial-container relative mx-auto px-gutter pb-20 pt-8">
	
	<aside class="absolute -left-12 top-24 hidden w-11 flex-col items-center gap-1.5 rounded-xl border border-outline-variant/20 bg-surface-container-lowest p-1.5 shadow-xs md:flex">
		<button type="button" class="flex h-8 w-8 items-center justify-center rounded font-bold text-outline transition-colors hover:bg-surface-container hover:text-primary" onclick={() => applyMarkdown('**')} title="太字">B</button>
		<button type="button" class="flex h-8 w-8 items-center justify-center rounded text-outline italic transition-colors hover:bg-surface-container hover:text-primary" onclick={() => applyMarkdown('*')} title="斜体">I</button>
		<button type="button" class="flex h-8 w-8 items-center justify-center rounded text-outline transition-colors hover:bg-surface-container hover:text-primary" onclick={() => applyMarkdown('- ', '')} title="リスト"><span class="material-symbols-outlined text-[18px]">format_list_bulleted</span></button>
		<button type="button" class="flex h-8 w-8 items-center justify-center rounded text-outline transition-colors hover:bg-surface-container hover:text-primary" onclick={insertLink} title="リンク"><span class="material-symbols-outlined text-[18px]">link</span></button>
	</aside>

	<main class="flex flex-col gap-6">
		<div class="hidden items-center justify-end gap-3 border-b border-outline-variant/20 pb-4 md:flex">
			{#if isAutoSaving}<span class="font-body-sm text-body-sm text-outline">自動保存中…</span>{:else if autoSaveError}<span class="font-body-sm text-body-sm text-error">自動保存に失敗しました</span>{:else if !isDirty}<span class="font-body-sm text-body-sm text-outline">保存済み</span>{/if}
			<button type="button" onclick={() => handleSave('draft')} disabled={isSubmitting} class="font-label-md text-label-md text-outline hover:text-primary">下書き保存</button>
			<button type="button" onclick={() => handleSave('published')} disabled={isSubmitting} class="font-label-md text-label-md rounded-lg bg-primary px-5 py-2 text-on-primary disabled:opacity-50">{isSubmitting ? '更新中...' : '更新する'}</button>
		</div>
		<!-- タイトル -->
		<div class="border-b border-outline-variant/10 pb-4 mb-4">
			<textarea
				bind:value={title}
				maxlength={MAX_TITLE_LENGTH}
				rows="2"
				aria-describedby="title-count"
				placeholder="タイトルを入力..."
				class="min-h-20 w-full resize-y overflow-y-auto bg-transparent px-0 py-1 text-[clamp(1.5rem,6vw,2rem)] font-bold leading-tight tracking-tight text-on-surface focus:outline-none md:text-[28px]"
				disabled={isSubmitting}
			></textarea>
			<p id="title-count" class="text-right text-body-sm text-outline">{title.length}/{MAX_TITLE_LENGTH}</p>
		</div>

		<!-- 本文 -->
		<div id="editor-body">
		<textarea
			bind:value={content}
			placeholder="物語を書き始めましょう..."
			class="min-h-[300px] w-full resize-y overflow-y-auto border-none bg-transparent px-0 py-1 text-body-md leading-relaxed text-on-surface focus:outline-none placeholder:text-outline-variant/50"
			disabled={isSubmitting}
		></textarea>
		</div>

	</main>
</div>
