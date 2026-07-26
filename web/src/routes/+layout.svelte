<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/Navbar.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import { page } from '$app/stores';

	let { children } = $props();

	// エディタ系のルートかどうかを判定
	let isEditor = $derived(
		$page.url.pathname.includes('/new') ||
		$page.url.pathname.includes('/edit')
	);
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<title>Essence — Quietude Editorial</title>
</svelte:head>

<div class="min-h-screen flex flex-col justify-between bg-surface text-on-surface">
	{#if !isEditor}
		<Navbar />
	{/if}
	<main class="flex-grow {isEditor ? 'pt-0 pb-0' : 'pt-28 pb-section-gap'}">
		{@render children()}
	</main>
	{#if !isEditor}
		<Footer />
	{/if}
</div>
