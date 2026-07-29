<script lang="ts">
	import './layout.css';
	import Navbar from '$lib/components/Navbar.svelte';
	import AdminNavbar from './admin/AdminNavbar.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import { page } from '$app/stores';

	let { children } = $props();

	// エディタ系のルートかどうかを判定
	let isEditor = $derived(
		$page.url.pathname.includes('/new') ||
		$page.url.pathname.includes('/edit')
	);
	let isAdmin = $derived($page.url.pathname.startsWith('/admin'));
</script>

<svelte:head>
	<link rel="icon" href="/favicon.svg" type="image/svg+xml" />
	<link rel="apple-touch-icon" href="/favicon.svg" />
	<link rel="manifest" href="/manifest.webmanifest" />
	<meta name="theme-color" content="#12201c" />
	<title>Buildlog — テクノロジーと美学をつなぐ開発録</title>
	<meta name="description" content="日々の開発プロセス、思考の断片、そして設計思想をミニマルかつ洗練されたデザインで綴る個人開発ログ。" />
	<meta property="og:title" content="Buildlog" />
	<meta property="og:description" content="日々の開発プロセス、思考の断片、そして設計思想をミニマルかつ洗練されたデザインで綴る個人開発ログ。" />
	<meta property="og:type" content="website" />
	<meta property="og:url" content="https://buildlog.dev{$page.url.pathname}" />
	<meta property="og:image" content="https://buildlog.dev/ogp.png" />
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:title" content="Buildlog" />
	<meta name="twitter:description" content="日々の開発プロセス、思考の断片、そして設計思想をミニマルかつ洗練されたデザインで綴る個人開発ログ。" />
</svelte:head>

<div class="min-h-screen flex flex-col justify-between bg-surface text-on-surface">
	{#if !isEditor}
		{#if isAdmin}
			<AdminNavbar />
		{:else}
			<Navbar />
		{/if}
	{/if}
	<main class="flex-grow {isEditor ? 'pt-0 pb-0' : isAdmin ? 'pt-24 pb-section-gap md:pl-64 md:pt-12' : 'pt-28 pb-section-gap'}">
		{@render children()}
	</main>
	{#if !isEditor}
		<Footer />
	{/if}
</div>
