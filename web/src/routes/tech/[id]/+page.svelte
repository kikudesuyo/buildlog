<script lang="ts">
	import { onMount, mount } from 'svelte';
	import { resolve } from '$app/paths';
	import { marked } from 'marked';
	import LikeButton from '$lib/components/LikeButton.svelte';
	import LinkCard from '$lib/components/LinkCard.svelte';
	let { data } = $props();
	let copyStatus = $state('');

	function formatDate(dateStr: string) {
		if (!dateStr) return '';
		const date = new Date(dateStr);
		return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`;
	}

	// 単独行のURLをLinkCardプレースホルダーに置き換える処理
	function parseContent(content: string): string {
		if (!content) return '';
		const lines = content.split('\n');
		const processedLines = lines.map(line => {
			const trimmed = line.trim();
			// 単独行のURL
			if (/^https?:\/\/[^\s]+$/.test(trimmed)) {
				return `<div class="link-card-placeholder" data-url="${trimmed}"></div>`;
			}
			return line;
		});
		return marked.parse(processedLines.join('\n'), { async: false }) as string;
	}

	let parsedHtml = $derived(parseContent(data.tech.content));

	async function copyCode(code: string) {
		try {
			await navigator.clipboard.writeText(code);
			copyStatus = 'コードをコピーしました';
		} catch {
			copyStatus = 'コピーできませんでした';
		}
		setTimeout(() => (copyStatus = ''), 2000);
	}

	onMount(() => {
		const placeholders = document.querySelectorAll('.link-card-placeholder');
		placeholders.forEach(el => {
			const url = el.getAttribute('data-url');
			if (url) {
				mount(LinkCard, {
					target: el,
					props: { url }
				});
			}
		});

		const buttons: HTMLButtonElement[] = [];
		document.querySelectorAll('pre').forEach((block) => {
			const button = document.createElement('button');
			button.type = 'button';
			button.textContent = 'コピー';
			button.className = 'code-copy-button';
			button.setAttribute('aria-label', 'コードをコピー');
			button.addEventListener('click', () =>
				copyCode(block.querySelector('code')?.textContent ?? block.textContent ?? '')
			);
			block.append(button);
			buttons.push(button);
		});

		return () => buttons.forEach((button) => button.remove());
	});
</script>

<svelte:head>
	<title>{data.tech.title} — Buildlog</title>
	<meta name="description" content={data.tech.content ? data.tech.content.substring(0, 120) + '...' : ''} />
	<meta property="og:title" content={data.tech.title} />
	<meta property="og:description" content={data.tech.content ? data.tech.content.substring(0, 120) + '...' : ''} />
	<meta property="og:type" content="article" />
	<meta property="og:image" content="https://buildlog.dev/ogp-tech.png" />
	<meta name="twitter:title" content={data.tech.title} />
	<meta name="twitter:description" content={data.tech.content ? data.tech.content.substring(0, 120) + '...' : ''} />
</svelte:head>

<div class="editorial-container mx-auto px-gutter pt-24 pb-20 max-w-[800px]">
	<!-- 戻るリンク -->
	<a
		href={resolve('/tech')}
		class="font-label-md text-label-md mb-8 inline-flex items-center gap-2 text-outline hover:text-primary transition-colors duration-200"
	>
		<span class="material-symbols-outlined text-[18px]">arrow_back</span>
		技術録へ戻る
	</a>

	<article class="flex flex-col gap-6">
		<header class="flex flex-col gap-4 border-b border-outline-variant/20 pb-8">
			<!-- メタ情報 -->
			<div class="flex flex-wrap items-center gap-stack-sm">
				<span class="font-label-sm text-label-sm rounded-full bg-secondary-container px-3 py-1 text-on-secondary-container">
					{data.tech.category}
				</span>
				<span class="font-label-sm text-label-sm text-on-surface-variant">
					{formatDate(data.tech.createdAt)}
				</span>
				{#if data.tech.views}
					<span class="h-1.5 w-1.5 rounded-full bg-outline-variant/30"></span>
					<span class="font-label-sm text-label-sm flex items-center gap-1 text-on-surface-variant">
						<span class="material-symbols-outlined text-[14px]">trending_up</span>
						{data.tech.views}
					</span>
				{/if}
			</div>

			<!-- タイトル -->
		<h1 class="font-display-lg text-[clamp(2rem,8vw,2.5rem)] leading-tight text-primary font-bold tracking-tight md:text-[40px]">
				{data.tech.title}
			</h1>
		</header>

		<!-- 本文 (Content) -->
		<section class="font-body-md text-body-md prose dark:prose-invert max-w-none break-words pt-4 mb-8 text-[1.0625rem] leading-8 text-on-surface md:text-body-md md:leading-relaxed">
			<!-- eslint-disable-next-line svelte/no-at-html-tags -->
			{@html parsedHtml}
		</section>
		<div aria-live="polite" class="min-h-6 text-label-sm text-primary">{copyStatus}</div>

		<div class="flex items-center gap-4 border-t border-outline-variant/10 pt-6">
			<LikeButton postId={data.tech.id} initialLikesCount={data.tech.likesCount} initialHasLiked={data.tech.hasLiked} />
		</div>
	</article>
</div>

	<style>
	:global(.prose pre) {
		position: relative;
		max-width: 100%;
		overflow-x: auto;
		padding-top: 3.25rem;
	}

	:global(.prose code) {
		overflow-wrap: anywhere;
	}

	:global(.code-copy-button) {
		position: absolute;
		right: 0.75rem;
		top: 0.75rem;
		min-height: 2.75rem;
		padding: 0.5rem 0.75rem;
		border-radius: 0.5rem;
		background: var(--color-surface-container-high);
		color: var(--color-on-surface);
	}
</style>
