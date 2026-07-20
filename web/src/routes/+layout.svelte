<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/Navbar.svelte';
	import Footer from '$lib/components/Footer.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<title>Essence — Quietude Editorial</title>
</svelte:head>

<div class="min-h-screen flex flex-col justify-between bg-surface text-on-surface">
	<Navbar />
	<main class="flex-grow pt-28 pb-section-gap">
		{@render children()}
	</main>
	<Footer />
</div>
