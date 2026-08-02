<script lang="ts">
import { onMount } from 'svelte';
import { page } from '$app/state';
import { resolve } from '$app/paths';

	let isSearchOpen = $state(false);
	let isMenuOpen = $state(false);
	let searchQuery = $state('');
	let isDarkMode = $state(false);
	let menuButton: HTMLButtonElement | null = null;
	let closeButton = $state<HTMLButtonElement | null>(null);
	let searchInput = $state<HTMLInputElement | null>(null);
	let previouslyFocused: HTMLElement | null = null;
	let searchHistoryEntry = false;

	onMount(() => {
		const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
		const syncTheme = () => {
			const savedTheme = localStorage.getItem('theme');
			isDarkMode = savedTheme ? savedTheme === 'dark' : mediaQuery.matches;
		};
		syncTheme();
		mediaQuery.addEventListener('change', syncTheme);
		const handlePopState = () => {
			if (isSearchOpen) closeSearch(true);
		};
		window.addEventListener('popstate', handlePopState);
		return () => {
			mediaQuery.removeEventListener('change', syncTheme);
			window.removeEventListener('popstate', handlePopState);
		};
	});

	function openSearch() {
		previouslyFocused = document.activeElement instanceof HTMLElement ? document.activeElement : null;
		isSearchOpen = true;
		window.history.pushState({ searchModal: true }, '');
		searchHistoryEntry = true;
	}

	function closeSearch(fromHistory = false) {
		if (!isSearchOpen) return;
		isSearchOpen = false;
		if (!fromHistory && searchHistoryEntry) {
			searchHistoryEntry = false;
			window.history.back();
		} else {
			searchHistoryEntry = false;
		}
		requestAnimationFrame(() => previouslyFocused?.focus());
	}

	function toggleTheme() {
		isDarkMode = !isDarkMode;
		if (isDarkMode) {
			document.documentElement.classList.add('dark');
			document.documentElement.classList.remove('light');
			document.documentElement.dataset.theme = 'dark';
			localStorage.setItem('theme', 'dark');
		} else {
			document.documentElement.classList.add('light');
			document.documentElement.classList.remove('dark');
			document.documentElement.dataset.theme = 'light';
			localStorage.setItem('theme', 'light');
		}
	}

	function toggleMenu() {
		if (isMenuOpen) {
			closeMenu();
			return;
		}

		previouslyFocused = document.activeElement instanceof HTMLElement ? document.activeElement : menuButton;
		isMenuOpen = true;
	}

	function closeMenu() {
		isMenuOpen = false;
		requestAnimationFrame(() => previouslyFocused?.focus());
	}

	function handleWindowKeydown(event: KeyboardEvent) {
		if (event.key === 'Escape' && isMenuOpen) {
			closeMenu();
		}
		if (event.key === 'Escape' && isSearchOpen) {
			closeSearch();
		}
	}

	$effect(() => {
		if (isSearchOpen) requestAnimationFrame(() => searchInput?.focus());
	});

	$effect(() => {
		if (isMenuOpen) {
			document.body.style.overflow = 'hidden';
			requestAnimationFrame(() => closeButton?.focus());
		} else {
			document.body.style.overflow = '';
		}

		return () => {
			document.body.style.overflow = '';
		};
	});

	const navItems = [
		{ href: '/', label: 'Diary' },
		{ href: '/tech', label: 'Tech' },
		{ href: '/profile', label: 'Profile' },
		{ href: '/apps', label: 'Apps' }
	] as const;

	function isActive(path: string): boolean {
		const current = page.url.pathname;
		if (path === '/') {
			return current === '/';
		}
		return current.startsWith(path);
	}
</script>

<svelte:window onkeydown={handleWindowKeydown} />

<!-- TopNavBar -->
<nav class="site-nav fixed top-0 z-50 w-full border-b border-outline-variant/30 bg-surface/80 backdrop-blur-md transition-all duration-300">
	<div class="site-nav-inner mx-auto flex h-16 max-w-container-max items-center justify-between px-gutter">
		<a href={resolve('/')} class="site-logo text-headline-md font-headline-md text-primary cursor-pointer transition-opacity active:opacity-70">
			Buildlog
		</a>
		<div class="site-nav-links flex items-center gap-stack-lg" aria-label="メインナビゲーション">
			{#each navItems as item (item.href)}
				<a
					href={resolve(item.href)}
					class="font-label-md text-label-md transition-colors duration-200 pb-0.5 {isActive(item.href)
						? 'text-primary font-bold border-b-2 border-primary'
						: 'text-on-surface-variant hover:text-primary'}"
				>
					{item.label}
				</a>
			{/each}
		</div>
		<div class="site-nav-actions flex items-center gap-stack-md">
			<button
				type="button"
				aria-label="検索を開く"
				onclick={() => (isSearchOpen ? closeSearch() : openSearch())}
				class="material-symbols-outlined text-on-surface-variant cursor-pointer hover:bg-surface-container-low rounded-lg p-2 transition-all duration-300"
			>
				search
			</button>
			<button
				type="button"
				aria-pressed={isDarkMode}
				aria-label={isDarkMode ? '現在はダークモード。ライトモードに切り替える' : '現在はライトモード。ダークモードに切り替える'}
				onclick={toggleTheme}
				class="material-symbols-outlined text-on-surface-variant cursor-pointer hover:bg-surface-container-low rounded-lg p-2 transition-all duration-300"
				title={isDarkMode ? 'ライトモードに切り替え' : 'ダークモードに切り替え'}
			>
				{isDarkMode ? 'light_mode' : 'dark_mode'}
			</button>
			<button
				bind:this={menuButton}
				type="button"
				aria-label={isMenuOpen ? 'メニューを閉じる' : 'メニューを開く'}
				aria-controls="mobile-navigation"
				aria-expanded={isMenuOpen}
				onclick={toggleMenu}
				class="material-symbols-outlined min-h-11 min-w-11 cursor-pointer rounded-lg p-2 text-on-surface-variant transition-all duration-300 hover:bg-surface-container-low md:hidden"
			>
				{isMenuOpen ? 'close' : 'menu'}
			</button>
		</div>
	</div>
</nav>

{#if isMenuOpen}
	<div id="mobile-navigation" class="fixed inset-0 z-40 md:hidden" role="presentation">
		<button
			type="button"
			aria-label="メニューを閉じる"
			class="absolute inset-0 bg-primary/20 backdrop-blur-xs"
			onclick={closeMenu}
		></button>
		<nav
			class="absolute right-0 top-0 flex h-full w-[min(84vw,22rem)] flex-col border-l border-outline-variant/30 bg-surface p-6 pt-24 shadow-xl"
			aria-label="モバイルナビゲーション"
		>
			<button
				bind:this={closeButton}
				type="button"
				aria-label="メニューを閉じる"
				onclick={closeMenu}
				class="material-symbols-outlined absolute right-4 top-4 min-h-11 min-w-11 rounded-lg p-2 text-on-surface-variant hover:bg-surface-container-low"
			>
				close
			</button>
			<div class="flex flex-col gap-2">
				{#each navItems as item (item.href)}
					<a
						href={resolve(item.href)}
						onclick={closeMenu}
						aria-current={isActive(item.href) ? 'page' : undefined}
						class="flex min-h-11 items-center rounded-lg px-4 font-label-md text-label-md transition-colors hover:bg-surface-container-low {isActive(item.href)
							? 'bg-surface-container-low font-bold text-primary'
							: 'text-on-surface-variant'}"
					>
						{item.label}
					</a>
				{/each}
			</div>
		</nav>
	</div>
{/if}
<!-- Search Modal (Minimalist) -->
{#if isSearchOpen}
	<div
		class="fixed inset-0 z-50 flex items-start justify-center overflow-y-auto bg-primary/20 px-4 pt-[max(7rem,calc(env(safe-area-inset-top)+5rem))] pb-8 backdrop-blur-xs"
		onclick={(event) => event.target === event.currentTarget && closeSearch()}
	>
		<div
			role="dialog"
			aria-modal="true"
			aria-labelledby="search-modal-title"
			tabindex="-1"
			class="max-h-[calc(100dvh-8rem)] w-full max-w-[540px] overflow-y-auto rounded-xl border border-outline-variant/30 bg-surface-container-lowest p-6 shadow-xl"
			onclick={(e) => e.stopPropagation()}
		>
			<div class="flex items-center gap-3 border-b border-outline-variant/20 pb-3">
				<span class="material-symbols-outlined text-on-surface-variant">search</span>
				<h2 id="search-modal-title" class="sr-only">記事を検索</h2>
				<input
					bind:this={searchInput}
					type="text"
					bind:value={searchQuery}
					aria-label="記事やキーワード"
					placeholder="記事やキーワードを検索..."
					class="w-full bg-transparent text-primary outline-none font-body-md placeholder:text-outline"
				/>
				<button
					type="button"
					aria-label="検索を閉じる"
					onclick={() => closeSearch()}
					class="min-h-11 min-w-11 rounded-lg px-2 text-label-sm text-outline hover:bg-surface-container-low hover:text-primary"
				>
					閉じる
				</button>
			</div>
			<div class="mt-4 text-label-sm text-on-surface-variant">
				{#if searchQuery.trim()}
					<p class="py-2">「{searchQuery}」の検索結果（デモ機能）</p>
				{:else}
					<p class="py-2 text-outline">キーワードを入力してください</p>
				{/if}
			</div>
		</div>
</div>
{/if}

<style>
	.site-nav {
		padding-inline: env(safe-area-inset-left) env(safe-area-inset-right);
		padding-block-start: env(safe-area-inset-top);
	}

	.site-nav button {
		min-block-size: 2.75rem;
		min-inline-size: 2.75rem;
	}

	@media (max-width: 767px) {
		.site-nav-inner {
			gap: 0.5rem;
			padding-inline: 1rem;
		}

		.site-logo {
			font-size: 1.25rem;
			line-height: 1.5rem;
			white-space: nowrap;
		}

		.site-nav-links {
			display: none;
		}

		.site-nav-actions {
			margin-left: auto;
			gap: 0.25rem;
		}

		.site-nav-actions button {
			min-height: 2.75rem;
			min-width: 2.75rem;
			padding: 0.625rem;
		}
	}
</style>
