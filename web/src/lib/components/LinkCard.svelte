<script lang="ts">
	import { onMount } from 'svelte';

	interface Props {
		url: string;
	}

	let { url }: Props = $props();

	interface OgpData {
		title: string;
		description: string;
		image: string;
		siteName: string;
	}

	let loading = $state(true);
	let data = $state<OgpData | null>(null);

	onMount(async () => {
		try {
			const res = await fetch(`/api/ogp?url=${encodeURIComponent(url)}`);
			if (res.ok) {
				data = await res.json();
			}
		} catch (e) {
			// ignore
		} finally {
			loading = false;
		}
	});
</script>

<div class="my-6">
	{#if loading}
		<!-- Loading Skeleton -->
		<div class="flex flex-col sm:flex-row h-auto sm:h-32 w-full animate-pulse rounded-xl border border-outline-variant/10 bg-surface-container-low overflow-hidden">
			<div class="flex-grow p-4 flex flex-col justify-between gap-2">
				<div class="h-5 w-3/4 rounded bg-outline-variant/20"></div>
				<div class="h-4 w-5/6 rounded bg-outline-variant/10"></div>
				<div class="h-3 w-1/4 rounded bg-outline-variant/10"></div>
			</div>
			<div class="hidden sm:block w-32 h-full bg-outline-variant/10"></div>
		</div>
	{:else if data}
		<a 
			href={url} 
			target="_blank" 
			rel="noopener noreferrer"
			class="flex flex-col sm:flex-row w-full rounded-xl border border-outline-variant/15 bg-surface-container-low overflow-hidden text-inherit no-underline transition-all duration-300 hover:border-primary/20 hover:bg-surface-container hover:shadow-md hover:-translate-y-[2px]"
		>
			<div class="flex-grow p-4 flex flex-col justify-between gap-1.5 min-w-0">
				<div class="flex flex-col gap-1">
					<span class="text-headline-xs font-bold line-clamp-1 text-on-surface-variant group-hover:text-primary">
						{data.title}
					</span>
					{#if data.description}
						<p class="text-body-sm text-outline line-clamp-2 m-0 leading-relaxed">
							{data.description}
						</p>
					{/if}
				</div>
				<div class="flex items-center gap-1.5 text-label-sm text-outline/80 mt-1">
					<span class="material-symbols-outlined text-[14px]">public</span>
					<span>{data.siteName}</span>
				</div>
			</div>
			{#if data.image}
				<div class="relative w-full sm:w-44 h-32 sm:h-auto overflow-hidden shrink-0 border-t sm:border-t-0 sm:border-l border-outline-variant/10">
					<img 
						src={data.image} 
						alt={data.title}
						class="w-full h-full object-cover transition-transform duration-500 hover:scale-105"
						loading="lazy"
					/>
				</div>
			{/if}
		</a>
	{:else}
		<!-- Fallback simple link -->
		<a href={url} target="_blank" rel="noopener noreferrer" class="text-primary hover:underline">
			{url}
		</a>
	{/if}
</div>
