<script lang="ts">
import { onMount } from 'svelte';
import { page } from '$app/state';
import { resolve } from '$app/paths';

	let isDarkMode = $state(false);

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

	function isActive(path: string): boolean {
		const current = page.url.pathname;
		if (path === '/') {
			return current === '/';
		}
		return current.startsWith(path);
	}
</script>

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
