<script lang="ts">
import { onMount } from 'svelte';
import { page } from '$app/state';
import { resolve } from '$app/paths';

	let isDarkMode = $state(false);
	let isOpen = $state(false);
	let menuButton: HTMLButtonElement | null = null;
	let previouslyFocused: HTMLElement | null = null;

	onMount(() => {
		const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
		const syncTheme = () => {
			const savedTheme = localStorage.getItem('theme');
			isDarkMode = savedTheme ? savedTheme === 'dark' : mediaQuery.matches;
		};
		syncTheme();
		mediaQuery.addEventListener('change', syncTheme);
		return () => {
			mediaQuery.removeEventListener('change', syncTheme);
		};
	});

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

	const navItems = [
		{ href: '/', label: 'Diary' },
		{ href: '/tech', label: 'Tech' },
		{ href: '/profile', label: 'Profile' },
		{ href: '/apps', label: 'Apps' }
	] as const;
	const desktopNavItems = [...navItems, { href: '/contact', label: 'Contact' }] as const;
	const mobileNavItems = [...navItems, { href: '/contact', label: 'Contact' }] as const;

	function isActive(path: string): boolean {
		const current = page.url.pathname;
		if (path === '/') {
			return current === '/';
		}
		return current.startsWith(path);
	}

	function openMenu() {
		previouslyFocused = document.activeElement instanceof HTMLElement ? document.activeElement : menuButton;
		isOpen = true;
	}

	function closeMenu() {
		isOpen = false;
		requestAnimationFrame(() => previouslyFocused?.focus());
	}

	function handleKeydown(event: KeyboardEvent) {
		if (event.key === 'Escape' && isOpen) closeMenu();
	}

	$effect(() => {
		document.body.style.overflow = isOpen ? 'hidden' : '';
		return () => {
			document.body.style.overflow = '';
		};
	});
</script>

<svelte:window onkeydown={handleKeydown} />

{#if isOpen}
	<button type="button" aria-label="メニューを閉じる" class="mobile-menu-panel fixed inset-0 z-40 bg-primary/20" onclick={closeMenu}></button>
	<aside aria-label="メインメニュー" class="mobile-menu-panel fixed inset-y-0 left-0 z-50 flex w-64 flex-col border-r border-outline-variant/20 bg-surface-container-lowest">
		<div class="flex h-16 items-center justify-between border-b border-outline-variant/20 px-4">
			<span class="font-headline-md text-headline-md font-bold tracking-tight text-primary">Menu</span>
			<button type="button" aria-label="メニューを閉じる" class="material-symbols-outlined min-h-11 min-w-11 rounded-lg p-2 text-on-surface-variant hover:bg-surface-container" onclick={closeMenu}>close</button>
		</div>
		<nav aria-label="メインメニュー" class="flex flex-col gap-1 p-4">
			{#each mobileNavItems as item (item.href)}
				<a href={resolve(item.href)} onclick={closeMenu} class="font-label-md text-label-md rounded-lg px-4 py-3 {isActive(item.href) ? 'bg-primary text-on-primary font-bold' : 'text-on-surface-variant hover:bg-surface-container hover:text-primary'}">{item.label}</a>
			{/each}
		</nav>
	</aside>
{/if}

<!-- TopNavBar -->
<nav class="site-nav fixed top-0 z-50 w-full border-b border-outline-variant/30 bg-surface/80 backdrop-blur-md transition-all duration-300">
	<div class="site-nav-inner mx-auto flex h-16 max-w-container-max items-center justify-between px-gutter">
		<div class="flex items-center gap-1">
		<button
			bind:this={menuButton}
			type="button"
			aria-label="メニューを開く"
			aria-expanded={isOpen}
			onclick={openMenu}
			class="mobile-menu-button material-symbols-outlined min-h-11 min-w-11 rounded-lg p-2 text-on-surface-variant hover:bg-surface-container"
		>
			menu
		</button>
		<a href={resolve('/')} class="site-logo text-headline-md font-headline-md text-primary cursor-pointer transition-opacity active:opacity-70">
			Buildlog
		</a>
		</div>
		<div class="site-nav-links flex items-center gap-stack-lg" aria-label="メインナビゲーション">
			{#each desktopNavItems as item (item.href)}
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
				aria-pressed={isDarkMode}
				aria-label={isDarkMode ? '現在はダークモード。ライトモードに切り替える' : '現在はライトモード。ダークモードに切り替える'}
				onclick={toggleTheme}
				class="material-symbols-outlined text-on-surface-variant cursor-pointer hover:bg-surface-container-low rounded-lg p-2 transition-all duration-300"
				title={isDarkMode ? 'ライトモードに切り替え' : 'ダークモードに切り替え'}
			>
				{isDarkMode ? 'light_mode' : 'dark_mode'}
			</button>
		</div>
	</div>
</nav>

<style>
	.site-nav {
		padding-inline: env(safe-area-inset-left) env(safe-area-inset-right);
		padding-block-start: env(safe-area-inset-top);
	}

	.site-nav button {
		min-block-size: 2.75rem;
		min-inline-size: 2.75rem;
	}

	.mobile-menu-button,
	.mobile-menu-panel {
		display: none;
	}

	@media (max-width: 767px) {
		.mobile-menu-button {
			display: inline-flex;
		}

		.mobile-menu-panel {
			display: flex;
		}

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
