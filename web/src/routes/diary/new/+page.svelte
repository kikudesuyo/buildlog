<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { createDiary } from '$lib/api/client';
	import UnsavedChangesGuard from '$lib/components/UnsavedChangesGuard.svelte';

	let title = $state('');
	const MAX_TITLE_LENGTH = 100;
	let contentElement = $state<HTMLTextAreaElement | null>(null);
	let content = $state('');
	let isSubmitting = $state(false);
	let errorMessage = $state('');
	let isDirty = $derived(title.trim().length > 0 || content.trim().length > 0);

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

	async function handleSave(status: 'draft' | 'published') {
		if (!title.trim() || !content.trim()) {
			errorMessage = 'タイトルと本文を入力してください。';
			return;
		}

		isSubmitting = true;
		errorMessage = '';
		try {
			await createDiary(title, content, status);
			goto(resolve('/admin'));
		} catch {
			errorMessage = 'つぶやきの保存に失敗しました。';
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
	<title>新しい下書き — Buildlog</title>
</svelte:head>

<!-- ヘッダー（全幅） -->
	<header class="fixed top-0 left-0 z-50 flex h-16 w-full items-center justify-between border-b border-outline-variant/20 bg-surface-container-lowest px-gutter">
	<div class="flex items-center gap-3">
		<a href={resolve('/admin')} class="text-headline-md font-headline-md text-primary font-bold tracking-tight">
			Buildlog
		</a>
		<span class="h-4 w-px bg-outline-variant/30"></span>
		<span class="text-outline font-label-md text-label-md">Drafts</span>
	</div>

	<div class="hidden items-center gap-6 md:flex">
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
			{isSubmitting ? '投稿中...' : '投稿する'}
		</button>
		<img
			src="/profile.jpg"
			alt="Profile"
			class="h-8 w-8 rounded-full object-cover border border-outline-variant/20"
		/>
	</div>
</header>

{#if errorMessage}
	<div class="mx-auto mt-20 max-w-container-max px-gutter text-body-sm text-error md:hidden" role="alert">{errorMessage}</div>
{/if}
<nav class="fixed inset-x-0 bottom-0 z-50 flex gap-2 border-t border-outline-variant/20 bg-white/95 p-3 pb-[max(0.75rem,env(safe-area-inset-bottom))] backdrop-blur md:hidden" aria-label="編集アクション">
	<button type="button" onclick={() => handleSave('draft')} disabled={isSubmitting} class="min-h-11 flex-1 rounded-lg border border-outline-variant/40 px-3 py-2 font-label-md text-outline disabled:opacity-50">下書き保存</button>
	<button type="button" onclick={() => handleSave('published')} disabled={isSubmitting} class="min-h-11 flex-1 rounded-lg bg-primary px-3 py-2 font-label-md text-on-primary disabled:opacity-50">{isSubmitting ? '投稿中...' : '投稿する'}</button>
</nav>

<!-- エディタ本体 -->
<div class="editorial-container relative mx-auto px-gutter pb-36 pt-24 md:pb-20">
	
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
