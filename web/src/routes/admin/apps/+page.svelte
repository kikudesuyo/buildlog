<script lang="ts">
	import { resolve } from '$app/paths';
	import { deleteApp } from '$lib/api/client';
	import IconButton from '$lib/components/IconButton.svelte';

	let { data } = $props();
	let appProjects = $state(data.appProjects);

	async function handleDelete(id: string) {
		if (!confirm('このアプリを削除してもよろしいですか？')) return;
		try {
			await deleteApp(Number(id));
			appProjects = appProjects.filter(app => app.id !== id);
		} catch {
			alert('削除に失敗しました。');
		}
	}
</script>

<svelte:head>
	<title>Apps 一覧 — Buildlog Admin</title>
</svelte:head>

<div class="editorial-container mx-auto px-gutter pt-8">
	<!-- ヘッダー -->
	<header class="flex justify-between items-center mb-8">
		<div>
			<h1 class="font-display-lg text-display-lg text-primary font-bold tracking-tight">Apps 管理</h1>
			<p class="font-body-md text-body-md text-on-surface-variant mt-1">紹介するプロダクトや実験を管理します。</p>
		</div>
		<a
			href={resolve('/admin/apps/new')}
			class="bg-primary text-on-primary font-label-md text-label-md px-5 py-2.5 rounded-lg font-medium hover:bg-primary/95 transition-colors shadow-sm flex items-center gap-2"
		>
			<span class="material-symbols-outlined text-[18px]">add</span>
			アプリを追加
		</a>
	</header>

	<!-- アプリ一覧 -->
	<div class="flex flex-col gap-4">
		{#if appProjects.length === 0}
			<div class="text-center py-12 border border-dashed border-outline-variant/30 rounded-2xl bg-surface-container-low text-outline">
				登録されているアプリはありません。
			</div>
		{:else}
			<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
				{#each appProjects as app (app.id)}
					<div class="bg-surface-container-lowest border border-outline-variant/20 rounded-2xl p-6 flex items-start gap-4 shadow-xs relative group transition-all duration-300 hover:shadow-md hover:border-primary/30">
						<div class="h-12 w-12 shrink-0 overflow-hidden rounded-xl border border-outline-variant/30 bg-surface-container-low flex items-center justify-center text-primary">
							{#if app.iconUrl}
								<img src={app.iconUrl} alt={app.name} class="h-full w-full object-contain p-1" />
							{:else}
								<span class="material-symbols-outlined text-2xl">{app.icon}</span>
							{/if}
						</div>
						<div class="flex-grow flex flex-col gap-1.5 pr-20">
							<h3 class="font-headline-sm text-headline-sm text-primary font-bold">{app.name}</h3>
							<p class="font-label-sm text-label-sm text-outline tracking-wider uppercase">{app.category}</p>
							<p class="font-body-md text-body-md text-on-surface-variant line-clamp-2 mt-1">{app.description}</p>
						</div>

						<!-- アクションボタン -->
						<div class="absolute top-4 right-4 flex gap-1">
							<a
								href={resolve(`/admin/apps/${app.id}/edit`)}
								class="material-symbols-outlined rounded-lg p-2 text-on-surface-variant hover:bg-surface-container-high transition-colors"
								title="編集"
							>
								edit
							</a>
							<IconButton icon="delete" variant="danger" type="button" onclick={() => handleDelete(app.id)} title="削除" aria-label="削除" />
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</div>
