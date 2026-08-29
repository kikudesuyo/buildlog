<script lang="ts">
	import { invalidateAll } from '$app/navigation';

	let { data } = $props();
	let failedImages = $state<Record<string, boolean>>({});
	let loadingImages = $state<Record<string, boolean>>({});
	let retrying = $state(false);

	function onImageError(id: string) {
		failedImages[id] = true;
		loadingImages[id] = false;
	}

	function onImageLoad(id: string) {
		loadingImages[id] = false;
	}

	async function retry() {
		retrying = true;
		await invalidateAll();
		retrying = false;
	}
</script>

<svelte:head>
	<title>Buildlog — Apps Showcase</title>
</svelte:head>

<div class="mx-auto w-full max-w-6xl px-gutter">
	<!-- Page Header -->
	<header class="mb-section-gap">
		<h1 class="font-display-lg text-display-lg mb-stack-sm text-primary">アプリ一覧</h1>
	</header>

	<!-- Project List -->
	{#if data.appProjects.length === 0}
		<section class="rounded-xl border border-outline-variant/30 bg-surface-container-low p-8 text-center" aria-live="polite">
			<span class="material-symbols-outlined mb-3 text-4xl text-outline">inventory_2</span>
			<h2 class="font-headline-md text-headline-md text-primary">公開中のアプリはありません</h2>
			<p class="mt-2 text-body-md text-on-surface-variant">時間をおいて再読み込みしてください。</p>
			<button type="button" onclick={retry} disabled={retrying} class="mt-6 min-h-11 rounded-lg bg-primary px-4 py-2 font-label-md text-on-primary disabled:opacity-60">
				{retrying ? '再読み込み中…' : '再読み込み'}
			</button>
		</section>
	{:else}
	<div class="grid grid-cols-1 gap-8 md:grid-cols-2 xl:grid-cols-3" aria-busy={retrying}>
		{#each data.appProjects as project (project.id)}
			<article class="group flex min-w-0 flex-col items-start gap-5 rounded-2xl border border-outline-variant/20 bg-surface-container-lowest p-6 shadow-sm transition-shadow hover:shadow-md">
				<a
					href={project.codeUrl || project.demoUrl}
					target="_blank"
					rel="noopener noreferrer"
				class="flex h-24 w-24 shrink-0 items-center justify-center overflow-hidden rounded-2xl border border-outline-variant/30 bg-surface-container-low p-1.5 transition-all duration-300 group-hover:border-primary/50 group-hover:shadow-md cursor-pointer"
					aria-label={`${project.name}の${project.codeUrl ? 'コード' : 'デモ'}を開く（外部サイト）`}
				>
					{#if project.iconUrl && !failedImages[project.id]}
						<img
							src={project.iconUrl}
							alt={project.name}
							class="h-full w-full object-contain transition-transform duration-300 group-hover:scale-105 {loadingImages[project.id] ? 'animate-pulse opacity-50' : ''}"
							loading="lazy"
							width="96"
							height="96"
							onload={() => onImageLoad(project.id)}
							onerror={() => onImageError(project.id)}
						/>
					{:else}
						<span
							class="material-symbols-outlined text-4xl text-outline opacity-50 transition-transform duration-300 group-hover:scale-110"
						>
							{project.icon}
						</span>
					{/if}
				</a>
				<div class="flex w-full min-w-0 flex-1 flex-col gap-stack-sm">
					<div class="flex flex-wrap items-start justify-between gap-4">
						<div>
							<h2 class="font-headline-md text-headline-md text-primary">
								<a
									href={project.codeUrl || project.demoUrl}
									target="_blank"
									rel="noopener noreferrer"
									class="hover:underline transition-colors hover:text-primary-container"
								>
									{project.name}
								</a>
							</h2>
					<p class="font-label-sm text-label-sm mt-1 tracking-widest text-on-surface-variant uppercase">
								{project.category}
							</p>
						</div>
						<div class="flex flex-wrap gap-2">
							{#each project.tags as tag (tag)}
								<span
									class="font-label-sm text-label-sm rounded border border-outline-variant/20 bg-secondary-container px-2 py-0.5 text-on-secondary-container"
								>
									{tag}
								</span>
							{/each}
						</div>
					</div>
					<p class="font-body-md text-body-md line-clamp-3 max-w-[560px] text-on-surface-variant">
						{project.description}
					</p>
					<div class="mt-2 flex flex-wrap gap-2">
						{#if project.demoUrl}
							<a
								href={project.demoUrl}
								target="_blank"
								rel="noopener noreferrer"
								class="text-label-md font-label-md flex min-h-11 items-center gap-1.5 rounded-lg bg-primary px-3 py-2 text-on-primary transition-colors hover:bg-primary-container"
								aria-label={`${project.name}のデモを開く（外部サイト）`}
							>
								<span class="material-symbols-outlined text-[18px]" aria-hidden="true">open_in_new</span>
								View Demo
							</a>
						{/if}
						{#if project.codeUrl}
							<a
								href={project.codeUrl}
								target="_blank"
								rel="noopener noreferrer"
								class="text-label-md font-label-md flex min-h-11 items-center gap-1.5 rounded-lg border border-outline-variant/30 px-3 py-2 text-on-surface-variant transition-colors hover:border-primary hover:text-primary"
								aria-label={`${project.name}のソースコードを開く（外部サイト）`}
							>
								<svg class="h-4 w-4 fill-current" viewBox="0 0 24 24" aria-hidden="true">
									<path
										fill-rule="evenodd"
										clip-rule="evenodd"
										d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.53 1.032 1.53 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z"
									/>
								</svg>
								View Code
							</a>
						{/if}
					</div>
				</div>
			</article>

		{/each}
	</div>
	{/if}

</div>
