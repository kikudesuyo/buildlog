<script lang="ts">
	import { resolve } from '$app/paths';
	import { updateProfile } from '$lib/api/client';
	import UnsavedChangesGuard from '$lib/components/UnsavedChangesGuard.svelte';

	const localAvatarUrl = '/profile.jpg';

	let { data } = $props();
	let savedProfile = $state(data.profileData);

	let name = $state(data.profileData.name);
	let title = $state(data.profileData.title);
	let quote = $state(data.profileData.quote);
	
	// bio は配列なので、改行で繋げてテキストエリアで編集
	let bioText = $state(data.profileData.bio.join('\n\n'));
	
	// highlights は配列のコピー
	let highlights = $state(JSON.parse(JSON.stringify(data.profileData.highlights)));
	
	let award = $state(data.profileData.award || '');
	
	// expertise は配列なので、カンマ区切りで編集
	let expertiseText = $state(data.profileData.expertise.join(', '));
	
	let contactEmail = $state(data.profileData.contactEmail);
	let finalQuote = $state(data.profileData.finalQuote);

	let isSubmitting = $state(false);
	let errorMessage = $state('');
	let successMessage = $state('');

	let isDirty = $derived(
		name !== savedProfile.name ||
		title !== savedProfile.title ||
		quote !== savedProfile.quote ||
		bioText !== savedProfile.bio.join('\n\n') ||
		JSON.stringify(highlights) !== JSON.stringify(savedProfile.highlights) ||
		award !== (savedProfile.award || '') ||
		expertiseText !== savedProfile.expertise.join(', ') ||
		contactEmail !== savedProfile.contactEmail ||
		finalQuote !== savedProfile.finalQuote
	);

	function addHighlight() {
		highlights = [...highlights, { title: '', period: '', description: '' }];
	}

	function removeHighlight(index: number) {
		highlights = highlights.filter((_: unknown, i: number) => i !== index);
	}

	async function handleSave() {
		if (!name.trim() || !contactEmail.trim()) {
			errorMessage = '名前とメールアドレスは必須です。';
			return;
		}

		isSubmitting = true;
		errorMessage = '';
		successMessage = '';

		const bio = bioText.split(/\n+/).map(p => p.trim()).filter(p => p.length > 0);
		const expertise = expertiseText.split(/[,，\n]+/).map(s => s.trim()).filter(s => s.length > 0);

		try {
			const updatedProfile = await updateProfile({
				name,
				title,
				quote,
				bio,
				highlights,
				award,
				expertise,
				contactEmail,
				finalQuote
			});
			savedProfile = updatedProfile;
			successMessage = 'プロフィールを更新しました。';
			setTimeout(() => {
				successMessage = '';
			}, 3000);
		} catch {
			errorMessage = 'プロフィールの更新に失敗しました。';
		} finally {
			isSubmitting = false;
		}
	}
</script>

<UnsavedChangesGuard {isDirty} {isSubmitting} />

<svelte:head>
	<title>プロフィール編集 — Buildlog</title>
</svelte:head>

<!-- ヘッダー -->
<header class="fixed top-16 left-0 w-full h-16 bg-surface border-b border-outline-variant/20 px-gutter flex items-center justify-between z-40 md:top-0 md:left-64 md:w-[calc(100%-16rem)] md:z-50">
	<div class="flex items-center gap-3">
		<a href={resolve('/admin')} class="text-headline-md font-headline-md text-primary font-bold tracking-tight">
			Buildlog
		</a>
		<span class="h-4 w-px bg-outline-variant/30"></span>
		<span class="text-outline font-label-md text-label-md">Admin</span>
	</div>

	<div class="flex items-center gap-6">
		{#if errorMessage}
			<span class="text-error font-body-sm text-body-sm">{errorMessage}</span>
		{/if}
		{#if successMessage}
			<span class="text-primary font-body-sm text-body-sm font-semibold">{successMessage}</span>
		{/if}
		<button
			type="button"
			onclick={handleSave}
			disabled={isSubmitting}
			class="bg-primary text-on-primary font-label-md text-label-md px-5 py-2 rounded-lg font-medium hover:bg-primary/95 transition-colors cursor-pointer disabled:opacity-50"
		>
			{isSubmitting ? '保存中...' : '保存する'}
		</button>
	</div>
</header>

	<div class="editorial-container mx-auto max-w-[800px] px-gutter pb-16 pt-28 md:pb-20 md:pt-24">
	<h1 class="font-display-lg mb-6 text-display-lg text-primary font-bold tracking-tight md:mb-8">プロフィール編集</h1>

	<form class="flex flex-col gap-7 rounded-2xl border border-outline-variant/20 bg-surface-container-lowest p-4 shadow-xs md:gap-8 md:p-8" onsubmit={(e) => e.preventDefault()}>
		<!-- 基本情報セクション -->
		<section class="flex flex-col gap-6">
			<h2 class="font-headline-md text-headline-md text-primary font-semibold border-b border-outline-variant/20 pb-2">基本情報</h2>
			
			<div class="grid min-w-0 grid-cols-1 gap-5 md:grid-cols-2 md:gap-6">
				<div class="flex flex-col gap-1.5">
					<label for="name" class="font-label-md text-label-md font-bold text-on-surface">氏名 / Name *</label>
					<input id="name" type="text" bind:value={name} class="variable-input min-w-0 rounded-none border-0 border-b border-outline-variant/50 bg-transparent px-0 py-2 text-body-md text-on-surface focus:border-primary focus:outline-none focus:ring-0" disabled={isSubmitting} />
				</div>

				<div class="flex flex-col gap-1.5">
					<label for="title" class="font-label-md text-label-md font-bold text-on-surface">肩書き / Title</label>
					<input id="title" type="text" bind:value={title} class="variable-input min-w-0 rounded-none border-0 border-b border-outline-variant/50 bg-transparent px-0 py-2 text-body-md text-on-surface focus:border-primary focus:outline-none focus:ring-0" disabled={isSubmitting} />
				</div>

				<div class="flex md:col-span-2">
					<img src={localAvatarUrl} alt="プロフィール画像" class="h-16 w-16 rounded-full border border-outline-variant/30 object-cover" />
				</div>
			</div>
		</section>

		<!-- 経歴・哲学セクション -->
		<section class="flex flex-col gap-6">
			<h2 class="font-headline-md text-headline-md text-primary font-semibold border-b border-outline-variant/20 pb-2">経歴と哲学</h2>

			<div class="flex flex-col gap-1.5">
				<label for="quote" class="font-label-md text-label-md font-bold text-on-surface">座右の銘 / Quote</label>
				<input id="quote" type="text" bind:value={quote} class="variable-input min-w-0 rounded-none border-0 border-b border-outline-variant/50 bg-transparent px-0 py-2 text-body-md text-on-surface focus:border-primary focus:outline-none focus:ring-0" disabled={isSubmitting} />
			</div>

			<div class="flex flex-col gap-1.5">
				<label for="bio" class="font-label-md text-label-md font-bold text-on-surface">自己紹介文 / Biography (改行で段落区切り)</label>
				<textarea id="bio" bind:value={bioText} class="variable-input min-h-[120px] min-w-0 resize-y rounded-none border-0 border-b border-outline-variant/50 bg-transparent px-0 py-2 text-body-md text-on-surface focus:border-primary focus:outline-none focus:ring-0 md:min-h-[160px]" disabled={isSubmitting}></textarea>
			</div>

			<div class="flex flex-col gap-1.5">
				<label for="finalQuote" class="font-label-md text-label-md font-bold text-on-surface">締めの言葉 / Final Quote</label>
				<input id="finalQuote" type="text" bind:value={finalQuote} class="variable-input min-w-0 rounded-none border-0 border-b border-outline-variant/50 bg-transparent px-0 py-2 text-body-md text-on-surface focus:border-primary focus:outline-none focus:ring-0" disabled={isSubmitting} />
			</div>
		</section>

		<!-- 主要な実績セクション -->
		<section class="flex flex-col gap-6">
			<div class="flex justify-between items-center border-b border-outline-variant/20 pb-2">
				<h2 class="font-headline-md text-headline-md text-primary font-semibold">主要な実績 / Highlights</h2>
				<button type="button" onclick={addHighlight} class="border border-dashed border-outline-variant/60 hover:border-primary px-3 py-1 rounded text-body-sm text-outline hover:text-primary transition-all cursor-pointer">
					+ 実績を追加
				</button>
			</div>

			<div class="flex flex-col gap-6">
				{#each highlights as highlight, index (index)}
					<div class="relative flex min-w-0 flex-col gap-4 rounded-xl border border-outline-variant/20 bg-surface-container-low p-4 group md:p-5">
						<button type="button" onclick={() => removeHighlight(index)} class="absolute top-4 right-4 text-outline hover:text-error transition-colors cursor-pointer font-bold text-sm" title="削除">
							削除
						</button>
						<div class="mr-10 flex min-w-0 flex-col gap-4">
							<div class="flex flex-col gap-1">
								<label for="hl-period-{index}" class="font-label-sm text-label-sm text-outline">期間 / Period</label>
								<input id="hl-period-{index}" type="text" bind:value={highlight.period} class="variable-input min-w-0 rounded-none border-0 border-b border-outline-variant/50 bg-transparent px-0 py-1 text-body-md text-on-surface focus:border-primary focus:outline-none focus:ring-0" disabled={isSubmitting} />
							</div>
							<div class="flex flex-col gap-1">
								<label for="hl-title-{index}" class="font-label-sm text-label-sm text-outline">実績名 / Title</label>
								<input id="hl-title-{index}" type="text" bind:value={highlight.title} class="variable-input min-w-0 rounded-none border-0 border-b border-outline-variant/50 bg-transparent px-0 py-1 text-body-md text-on-surface focus:border-primary focus:outline-none focus:ring-0" disabled={isSubmitting} />
							</div>
							<div class="flex flex-col gap-1">
								<label for="hl-desc-{index}" class="font-label-sm text-label-sm text-outline">説明 / Description</label>
								<textarea id="hl-desc-{index}" bind:value={highlight.description} class="variable-input h-16 min-w-0 resize-none rounded-none border-0 border-b border-outline-variant/50 bg-transparent px-0 py-1 text-body-md text-on-surface focus:border-primary focus:outline-none focus:ring-0" disabled={isSubmitting}></textarea>
							</div>
						</div>
					</div>
				{/each}
			</div>
		</section>

		<!-- 専門スキル・その他セクション -->
		<section class="flex flex-col gap-6">
			<h2 class="font-headline-md text-headline-md text-primary font-semibold border-b border-outline-variant/20 pb-2">専門スキル & 連絡先</h2>

			<div class="grid min-w-0 grid-cols-1 gap-5 md:grid-cols-2 md:gap-6">
				<div class="flex flex-col gap-1.5 md:col-span-2">
					<label for="expertise" class="font-label-md text-label-md font-bold text-on-surface">専門領域 / Expertise (カンマ区切り)</label>
					<input id="expertise" type="text" bind:value={expertiseText} class="variable-input min-w-0 rounded-none border-0 border-b border-outline-variant/50 bg-transparent px-0 py-2 text-body-md text-on-surface focus:border-primary focus:outline-none focus:ring-0" disabled={isSubmitting} />
				</div>

				<div class="flex flex-col gap-1.5">
					<label for="award" class="font-label-md text-label-md font-bold text-on-surface">受賞歴 / Award</label>
					<input id="award" type="text" bind:value={award} class="variable-input min-w-0 rounded-none border-0 border-b border-outline-variant/50 bg-transparent px-0 py-2 text-body-md text-on-surface focus:border-primary focus:outline-none focus:ring-0" disabled={isSubmitting} />
				</div>

				<div class="flex flex-col gap-1.5">
					<label for="contactEmail" class="font-label-md text-label-md font-bold text-on-surface">連絡先メール / Email *</label>
					<input id="contactEmail" type="email" bind:value={contactEmail} class="variable-input min-w-0 rounded-none border-0 border-b border-outline-variant/50 bg-transparent px-0 py-2 text-body-md text-on-surface focus:border-primary focus:outline-none focus:ring-0" disabled={isSubmitting} />
				</div>
			</div>
		</section>
	</form>
</div>

<style>
	.variable-input {
		align-self: flex-start;
		width: fit-content;
		min-width: min(12rem, 100%);
		max-width: 100%;
		field-sizing: content;
	}
</style>
