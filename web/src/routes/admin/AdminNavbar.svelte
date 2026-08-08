<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/state';
	import { resolve } from '$app/paths';

	const navItems = [
		{ href: '/admin', label: 'Diary', description: 'つぶやきを管理', icon: 'edit_note' },
		{ href: '/admin/tech', label: 'Tech', description: '技術記事を管理', icon: 'article' },
		{ href: '/admin/profile', label: 'Profile', description: '自己紹介を管理', icon: 'person' },
		{ href: '/admin/apps', label: 'Apps', description: 'プロダクトを管理', icon: 'apps' },
		{ href: '/admin/analytics', label: 'Analytics', description: 'アクセス統計・分析', icon: 'analytics' },
		{ href: '/admin/trash', label: 'Trash', description: 'ゴミ箱', icon: 'delete_outline' },
		{ href: '/admin/comments', label: 'Comments', description: 'コメント管理', icon: 'comment' }
	] as const;

	function resolveAdminPath(path: (typeof navItems)[number]['href']) {
		switch (path) {
			case '/admin/analytics':
				return resolve('/admin/analytics');
			case '/admin/tech':
				return resolve('/admin/tech');
			case '/admin/trash':
				return resolve('/admin/trash');
			case '/admin/comments':
				return resolve('/admin/comments');
			case '/admin/apps':
				return resolve('/admin/apps');
			case '/admin/profile':
				return resolve('/admin/profile');
			default:
				return resolve('/admin');
		}
	}

	let isOpen = $state(false);
	let isDarkMode = $state(false);

	onMount(() => {
		isDarkMode = document.documentElement.classList.contains('dark');
	});

	function toggleTheme() {
		isDarkMode = !isDarkMode;
		document.documentElement.classList.toggle('dark', isDarkMode);
		document.documentElement.classList.toggle('light', !isDarkMode);
		localStorage.setItem('theme', isDarkMode ? "dark" : "light");
	}
	let menuButton: HTMLButtonElement | null = null;
	let closeButton = $state<HTMLButtonElement | null>(null);
	let previouslyFocused: HTMLElement | null = null;

	function isActive(href: string) {
		if (href === '/admin') return page.url.pathname === '/admin';
		return page.url.pathname.startsWith(href);
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
		if (isOpen) requestAnimationFrame(() => closeButton?.focus());
		return () => {
			document.body.style.overflow = '';
		};
	});
</script>

<svelte:window onkeydown={handleKeydown} />

{#if isOpen}
	<button
		type="button"
		aria-label="管理メニューを閉じる"
		class="fixed inset-0 z-40 bg-primary/20 md:hidden"
		onclick={closeMenu}
	></button>
{/if}

<aside aria-label="管理メニュー" aria-modal={isOpen ? 'true' : undefined} class="fixed inset-y-0 left-0 z-50 w-64 flex-col border-r border-outline-variant/20 bg-surface-container-lowest {isOpen ? 'flex' : 'hidden'} md:flex">
	<button
		bind:this={closeButton}
		type="button"
		aria-label="管理メニューを閉じる"
		onclick={closeMenu}
		class="material-symbols-outlined absolute right-3 top-3 z-10 min-h-11 min-w-11 rounded-lg p-2 text-on-surface-variant hover:bg-surface-container md:hidden"
	>
		close
	</button>
	<div class="flex h-20 items-center border-b border-outline-variant/20 px-6">
		<a href={resolve('/admin')} class="flex items-center gap-3 text-primary transition-opacity hover:opacity-80">
			<span class="material-symbols-outlined text-2xl">dashboard_customize</span>
			<div>
				<div class="font-headline-md text-headline-md font-bold tracking-tight">Buildlog</div>
				<div class="font-label-sm text-label-sm tracking-widest text-outline uppercase">Admin Console</div>
			</div>
		</a>
	</div>

	<div class="flex flex-1 flex-col px-3 py-6">
		<p class="font-label-sm text-label-sm mb-3 px-3 tracking-[0.2em] text-outline uppercase">Workspace</p>
		<nav aria-label="管理メニュー" class="flex flex-col gap-1">
			{#each navItems as item (item.href)}
				<a
					href={resolveAdminPath(item.href)}
					onclick={closeMenu}
					class="flex items-center gap-3 rounded-xl px-3 py-3 transition-colors {isActive(item.href) ? 'bg-primary text-on-primary shadow-sm' : 'text-on-surface-variant hover:bg-surface-container hover:text-primary'}"
				>
					<span class="material-symbols-outlined text-[21px]">{item.icon}</span>
					<span class="flex flex-col">
						<span class="font-label-md text-label-md font-semibold">{item.label}</span>
						<span class="font-label-sm text-label-sm {isActive(item.href) ? 'text-on-primary/70' : 'text-outline'}">{item.description}</span>
					</span>
				</a>
			{/each}
		</nav>

		<div class="mt-auto rounded-xl border border-outline-variant/20 bg-surface-container-low p-4">
			<p class="font-label-sm text-label-sm mb-1 text-outline">現在のモード</p>
			<p class="font-label-md text-label-md flex items-center gap-2 font-semibold text-primary">
				<span class="h-2 w-2 rounded-full bg-primary"></span>
				Content Manager
			</p>
		</div>
	</div>

	<div class="border-t border-outline-variant/20 p-4">
		<a href={resolve('/')} onclick={closeMenu} class="font-label-md text-label-md flex items-center gap-2 rounded-lg px-3 py-2 text-on-surface-variant transition-colors hover:bg-surface-container hover:text-primary">
			<span class="material-symbols-outlined text-[18px]">visibility</span>
			公開サイトを見る
		</a>
	</div>
</aside>

<header class="fixed top-0 z-50 flex h-16 w-full items-center justify-between border-b border-outline-variant/20 bg-surface-container-lowest px-gutter md:hidden">
	<div class="flex items-center gap-3">
		<button
			bind:this={menuButton}
			type="button"
			aria-label="管理メニューを開く"
			aria-expanded={isOpen}
			onclick={openMenu}
			class="material-symbols-outlined min-h-11 min-w-11 rounded-lg p-2 text-on-surface-variant hover:bg-surface-container"
		>
			menu
		</button>
		<a href={resolve('/admin')} class="font-headline-md text-headline-md font-bold tracking-tight text-primary">Buildlog <span class="font-label-sm text-label-sm ml-2 text-outline">Admin</span></a>
	</div>
	<a href={resolve('/')} class="material-symbols-outlined rounded-lg p-2 text-on-surface-variant hover:bg-surface-container">visibility</a>
	<button type="button" aria-label="テーマを切り替える" onclick={toggleTheme} class="material-symbols-outlined min-h-11 min-w-11 rounded-lg p-2 text-on-surface-variant hover:bg-surface-container" title={isDarkMode ? "ライトモードに切り替え" : "ダークモードに切り替え"}>
		{isDarkMode ? 'light_mode' : 'dark_mode'}
	</button>
</header>
