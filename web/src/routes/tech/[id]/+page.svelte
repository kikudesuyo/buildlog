<script lang="ts">
	import { resolve } from '$app/paths';
	import LikeButton from '$lib/components/LikeButton.svelte';
	import CommentSection from '$lib/components/CommentSection.svelte';
	import { Marked, type Tokens } from 'marked';
	import { markedHighlight } from 'marked-highlight';
	import hljs from 'highlight.js';
	import 'highlight.js/styles/github-dark.css';

	let { data } = $props();
	let copyMessage = $state('');
	let copyMessageTimer: ReturnType<typeof setTimeout> | undefined;

	function escapeHtml(value: string) {
		return value
			.replaceAll('&', '&amp;')
			.replaceAll('<', '&lt;')
			.replaceAll('>', '&gt;')
			.replaceAll('"', '&quot;')
			.replaceAll("'", '&#39;');
	}

	const marked = new Marked(
		markedHighlight({
			emptyLangClass: 'hljs',
			langPrefix: 'hljs language-',
			highlight(code, lang) {
				const language = hljs.getLanguage(lang) ? lang : 'plaintext';
				return hljs.highlight(code, { language }).value;
			}
		})
	);

	marked.use({
		renderer: {
			html: () => '',
			code({ text, lang, escaped }: Tokens.Code) {
				const highlighted = escaped ? text : escapeHtml(text);
				const encodedCode = encodeURIComponent(text);
				const languageLabel = lang?.trim() || 'text';

				return `<div class="markdown-code-block"><div class="markdown-code-toolbar"><span class="markdown-code-language">${escapeHtml(languageLabel)}</span><button type="button" class="markdown-code-copy" data-copy-code="${encodedCode}" aria-label="${escapeHtml(languageLabel)}のコードをコピー">Copy</button></div><pre><code class="hljs language-${escapeHtml(languageLabel)}">${highlighted}</code></pre></div>`;
			}
		}
	});

	const parsedContent = $derived(marked.parse(data.tech.content || ''));

	async function copyCode(event: MouseEvent) {
		const target = event.target;
		if (!(target instanceof HTMLButtonElement) || !target.dataset.copyCode) return;

		const code = decodeURIComponent(target.dataset.copyCode);
		try {
			await navigator.clipboard.writeText(code);
			copyMessage = 'コードをコピーしました';
			target.textContent = 'Copied';
		} catch {
			copyMessage = 'コードをコピーできませんでした';
			target.textContent = 'Retry';
		}

		if (copyMessageTimer) clearTimeout(copyMessageTimer);
		copyMessageTimer = setTimeout(() => {
			copyMessage = '';
		}, 3000);
	}

	function handleProseKeydown(event: KeyboardEvent) {
		if (event.key === 'Enter' || event.key === ' ') {
			copyCode(event as unknown as MouseEvent);
		}
	}

	function formatDate(dateStr: string) {
		if (!dateStr) return '';
		const date = new Date(dateStr);
		return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`;
	}
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
		<section
			class="font-body-md text-body-md prose dark:prose-invert max-w-none break-words pt-4 mb-8 text-[1.0625rem] leading-8 text-on-surface md:text-body-md md:leading-relaxed"
			onclick={copyCode}
			onkeydown={handleProseKeydown}
			role="group"
			tabindex="-1"
		>
			{@html parsedContent}
		</section>
		{#if copyMessage}
			<p class="sr-only" role="status" aria-live="polite">{copyMessage}</p>
		{/if}

		<div class="flex items-center gap-4 border-t border-outline-variant/10 pt-6">
			<LikeButton postId={data.tech.id} initialLikesCount={data.tech.likesCount} initialHasLiked={data.tech.hasLiked} />
		</div>
		<CommentSection postId={data.tech.id} />
	</article>
</div>

	<style>
	:global(.prose pre) {
		position: relative;
		max-width: 100%;
		overflow-x: auto;
		padding-top: 3.25rem;
	}

	:global(.markdown-code-block) {
		margin: 1.5rem 0;
		overflow: hidden;
		border: 1px solid var(--color-outline-variant);
		border-radius: 0.75rem;
		background: #0d1117;
	}

	:global(.markdown-code-toolbar) {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 0.75rem;
		padding: 0.5rem 0.75rem;
		border-bottom: 1px solid rgb(255 255 255 / 12%);
		font-family: var(--font-hanken);
		font-size: 0.75rem;
		color: #c9d1d9;
	}

	:global(.markdown-code-language) {
		text-transform: uppercase;
		letter-spacing: 0.08em;
	}

	:global(.markdown-code-copy) {
		min-width: 4.5rem;
		min-height: 2.25rem;
		border: 1px solid rgb(255 255 255 / 25%);
		border-radius: 0.5rem;
		padding: 0.35rem 0.65rem;
		color: #f0f6fc;
		font-family: var(--font-hanken);
		font-size: 0.75rem;
		cursor: pointer;
	}

	:global(.markdown-code-copy:hover),
	:global(.markdown-code-copy:focus-visible) {
		background: rgb(255 255 255 / 12%);
		outline: 2px solid #b7cbbf;
		outline-offset: 2px;
	}

	:global(.markdown-code-block pre) {
		margin: 0;
		overflow-x: auto;
		padding: 1rem;
		font-size: 0.875rem;
		line-height: 1.6;
	}

	:global(.prose code) {
		overflow-wrap: anywhere;
	}
</style>
