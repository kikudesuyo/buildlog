<script lang="ts">
import { page } from '$app/state';
import { resolve } from '$app/paths';

	let isSearchOpen = $state(false);
	let searchQuery = $state('');

	// 現在のパスが管理者画面（/admin で始まる）かどうかを判定
	let isAdmin = $derived(page.url.pathname.startsWith('/admin'));

	const navItems = [
		{ href: '/', label: 'Diary' },
		{ href: '/tech', label: 'Tech' },
		{ href: '/profile', label: 'Profile' },
		{ href: '/apps', label: 'Apps' }
	] as const;

	// 管理画面にいるときはリンク先を /admin 等に差し替える
	let computedNavItems = $derived(
		navItems.map(item => {
			if (isAdmin) {
				if (item.href === '/') return { ...item, href: '/admin' };
				if (item.href === '/tech') return { ...item, href: '/admin/tech' };
			}
			return item;
		})
	);

	function isActive(path: string): boolean {
		const current = page.url.pathname;
		if (path === '/' || path === '/admin') {
			return current === '/' || current === '/admin';
		}
		if (path === '/tech' || path === '/admin/tech') {
			return current.startsWith('/tech') || current.startsWith('/admin/tech');
		}
		return current.startsWith(path);
	}
</script>

<!-- TopNavBar -->
<nav class="fixed top-0 w-full z-50 bg-surface/80 backdrop-blur-md border-b border-outline-variant/30 transition-all duration-300">
	<div class="flex justify-between items-center max-w-container-max mx-auto h-16 px-gutter">
		<a href={isAdmin ? '/admin' : '/'} class="text-headline-md font-headline-md text-primary cursor-pointer transition-opacity active:opacity-70">
			Essence
		</a>
		<div class="flex items-center gap-stack-lg">
			{#each computedNavItems as item (item.href)}
				<a
					href={item.href}
					class="font-label-md text-label-md transition-colors duration-200 pb-0.5 {isActive(item.href)
						? 'text-primary font-bold border-b-2 border-primary'
						: 'text-on-surface-variant hover:text-primary'}"
				>
					{item.label}
				</a>
			{/each}
		</div>
		<div class="flex items-center gap-stack-md">
			<a
				href={isAdmin ? '/' : '/admin'}
				class="material-symbols-outlined text-on-surface-variant cursor-pointer hover:bg-surface-container-low rounded-lg p-2 transition-all duration-300 flex items-center justify-center"
				title={isAdmin ? "一般表示に切り替え" : "管理画面に切り替え"}
			>
				{isAdmin ? 'visibility' : 'admin_panel_settings'}
			</a>
			<button
				type="button"
				aria-label="Search"
				onclick={() => (isSearchOpen = !isSearchOpen)}
				class="material-symbols-outlined text-on-surface-variant cursor-pointer hover:bg-surface-container-low rounded-lg p-2 transition-all duration-300"
			>
				search
			</button>
			<button
				type="button"
				aria-label="Settings"
				class="material-symbols-outlined text-on-surface-variant cursor-pointer hover:bg-surface-container-low rounded-lg p-2 transition-all duration-300"
			>
				settings
			</button>
		</div>
	</div>
</nav>

<!-- Search Modal (Minimalist) -->
{#if isSearchOpen}
	<div
		role="button"
		tabindex="0"
		class="fixed inset-0 z-50 bg-primary/20 backdrop-blur-xs flex items-start justify-center pt-28 px-4"
		onclick={() => (isSearchOpen = false)}
		onkeydown={(e) => e.key === 'Escape' && (isSearchOpen = false)}
	>
		<div
			role="dialog"
			aria-modal="true"
			tabindex="-1"
			class="bg-surface-container-lowest border border-outline-variant/30 rounded-xl p-6 w-full max-w-[540px] shadow-xl"
			onclick={(e) => e.stopPropagation()}
			onkeydown={(e) => e.key === 'Escape' && (isSearchOpen = false)}
		>
			<div class="flex items-center gap-3 border-b border-outline-variant/20 pb-3">
				<span class="material-symbols-outlined text-on-surface-variant">search</span>
				<input
					type="text"
					bind:value={searchQuery}
					placeholder="記事やキーワードを検索..."
					class="w-full bg-transparent text-primary outline-none font-body-md placeholder:text-outline"
				/>
				<button
					type="button"
					onclick={() => (isSearchOpen = false)}
					class="text-label-sm text-outline hover:text-primary"
				>
					ESC
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
