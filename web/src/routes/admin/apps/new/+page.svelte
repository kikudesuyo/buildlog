<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { createApp } from '$lib/api/client';
	import UnsavedChangesGuard from '$lib/components/UnsavedChangesGuard.svelte';

	let slug = $state('');
	let name = $state('');
	let category = $state('');
	let tagsText = $state('');
	let description = $state('');
	let icon = $state('explore');
	let iconUrl = $state('');
	let demoUrl = $state('');
	let codeUrl = $state('');

	let isSubmitting = $state(false);
	let errorMessage = $state('');
	let isDirty = $derived(
		slug.trim().length > 0 ||
		name.trim().length > 0 ||
		category.trim().length > 0 ||
		tagsText.trim().length > 0 ||
		description.trim().length > 0 ||
		iconUrl.trim().length > 0 ||
		demoUrl.trim().length > 0 ||
		codeUrl.trim().length > 0
	);

	async function handleSave() {
		if (!slug.trim() || !name.trim() || !category.trim() || !description.trim()) {
			errorMessage = '必須項目を入力してください。';
			return;
		}

		isSubmitting = true;
		errorMessage = '';
		
		const tags = tagsText.split(/[,，\n]+/).map(t => t.trim()).filter(t => t.length > 0);

		try {
			await createApp({
				slug: slug.trim(),
				name: name.trim(),
				category: category.trim(),
				tags,
				description: description.trim(),
				icon: icon.trim(),
				iconUrl: iconUrl.trim(),
				demoUrl: demoUrl.trim(),
				codeUrl: codeUrl.trim()
			});
			goto(resolve('/admin/apps'));
		} catch {
			errorMessage = 'アプリの保存に失敗しました。';
			isSubmitting = false;
		}
	}
</script>

<UnsavedChangesGuard {isDirty} {isSubmitting} />

<svelte:head>
	<title>新規アプリ登録 — Buildlog Admin</title>
</svelte:head>

<!-- ヘッダー -->
<header class="fixed top-0 left-0 w-full h-16 bg-white border-b border-outline-variant/20 px-gutter flex items-center justify-between z-50">
	<div class="flex items-center gap-3">
		<a href={resolve('/admin/apps')} class="text-headline-md font-headline-md text-primary font-bold tracking-tight">
			Buildlog
		</a>
		<span class="h-4 w-px bg-outline-variant/30"></span>
		<span class="text-outline font-label-md text-label-md">Apps</span>
	</div>

	<div class="flex items-center gap-6">
		{#if errorMessage}
			<span class="text-error font-body-sm text-body-sm">{errorMessage}</span>
		{/if}
		<button
			type="button"
			onclick={handleSave}
			disabled={isSubmitting}
			class="bg-primary text-on-primary font-label-md text-label-md px-5 py-2 rounded-lg font-medium hover:bg-primary/95 transition-colors cursor-pointer disabled:opacity-50"
		>
			{isSubmitting ? '登録中...' : '登録する'}
		</button>
	</div>
</header>

<div class="editorial-container mx-auto px-gutter pt-24 pb-20 max-w-[600px]">
	<h1 class="font-display-lg text-display-lg mb-8 text-primary font-bold tracking-tight">新規アプリ登録</h1>

	<form class="flex flex-col gap-6" onsubmit={(e) => e.preventDefault()}>
		<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
			<div class="flex flex-col gap-1.5">
				<label for="slug" class="font-label-md text-label-md font-bold text-on-surface">識別子 / Slug *</label>
				<input id="slug" type="text" placeholder="例: whichway" bind:value={slug} class="rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md" disabled={isSubmitting} />
			</div>

			<div class="flex flex-col gap-1.5">
				<label for="name" class="font-label-md text-label-md font-bold text-on-surface">アプリ名 / Name *</label>
				<input id="name" type="text" placeholder="例: Whichway" bind:value={name} class="rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md" disabled={isSubmitting} />
			</div>

			<div class="flex flex-col gap-1.5 md:col-span-2">
				<label for="category" class="font-label-md text-label-md font-bold text-on-surface">カテゴリ / Category *</label>
				<input id="category" type="text" placeholder="例: Tool / Decision Support" bind:value={category} class="rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md" disabled={isSubmitting} />
			</div>

			<div class="flex flex-col gap-1.5 md:col-span-2">
				<label for="tags" class="font-label-md text-label-md font-bold text-on-surface">タグ / Tags (カンマ区切り)</label>
				<input id="tags" type="text" placeholder="例: TypeScript, Go" bind:value={tagsText} class="rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md" disabled={isSubmitting} />
			</div>

			<div class="flex flex-col gap-1.5 md:col-span-2">
				<label for="description" class="font-label-md text-label-md font-bold text-on-surface">説明 / Description *</label>
				<textarea id="description" placeholder="アプリの説明を入力..." bind:value={description} class="rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md min-h-[120px] resize-y" disabled={isSubmitting}></textarea>
			</div>

			<div class="flex flex-col gap-1.5">
				<label for="icon" class="font-label-md text-label-md font-bold text-on-surface">マテリアルアイコン名 / Icon *</label>
				<input id="icon" type="text" placeholder="例: explore" bind:value={icon} class="rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md" disabled={isSubmitting} />
			</div>

			<div class="flex flex-col gap-1.5">
				<label for="iconUrl" class="font-label-md text-label-md font-bold text-on-surface">アイコン画像 URL / Icon URL</label>
				<input id="iconUrl" type="text" placeholder="例: /whichway-icon.svg" bind:value={iconUrl} class="rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md" disabled={isSubmitting} />
			</div>

			<div class="flex flex-col gap-1.5 md:col-span-2">
				<label for="demoUrl" class="font-label-md text-label-md font-bold text-on-surface">デモ URL / Demo URL</label>
				<input id="demoUrl" type="url" placeholder="例: https://..." bind:value={demoUrl} class="rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md" disabled={isSubmitting} />
			</div>

			<div class="flex flex-col gap-1.5 md:col-span-2">
				<label for="codeUrl" class="font-label-md text-label-md font-bold text-on-surface">ソースコード URL / Code URL</label>
				<input id="codeUrl" type="url" placeholder="例: https://github.com/..." bind:value={codeUrl} class="rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-on-surface focus:outline-none text-body-md" disabled={isSubmitting} />
			</div>
		</div>
	</form>
</div>
