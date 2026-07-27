<script lang="ts">
	import * as env from '$env/static/public';
	import { resolve } from '$app/paths';

	const githubUrl = env.PUBLIC_GITHUB_URL || 'https://github.com/kikudesuyo';
	const xUrl = env.PUBLIC_X_URL || 'https://x.com/kikudesuyo';
</script>

<!-- Footer -->
<footer class="bg-surface w-full py-section-gap border-t border-outline-variant/20">
	<div class="max-w-container-max mx-auto flex flex-col md:flex-row justify-between items-center gap-stack-md px-gutter">
		<div class="font-label-md text-label-md text-primary font-semibold">
			Buildlog
		</div>
		<div class="flex gap-stack-lg items-center">
			<a href="#rss" class="font-label-sm text-label-sm text-on-surface-variant hover:text-primary transition-colors">RSS</a>
			<a href={resolve('/profile')} class="font-label-sm text-label-sm text-on-surface-variant hover:text-primary transition-colors">Contact</a>
			{#if githubUrl}
				<a href={githubUrl} target="_blank" rel="noopener noreferrer" class="text-on-surface-variant hover:text-primary transition-colors flex items-center" title="GitHub">
					<svg class="h-4 w-4 fill-current" viewBox="0 0 24 24">
						<path d="M12 2A10 10 0 0 0 2 12c0 4.42 2.87 8.17 6.84 9.5.5.08.66-.23.66-.5v-1.69c-2.77.6-3.36-1.34-3.36-1.34-.46-1.16-1.11-1.47-1.11-1.47-.9-.62.07-.6.07-.6 1 .07 1.53 1.03 1.53 1.03.9 1.52 2.34 1.07 2.91.83.09-.65.35-1.09.63-1.34-2.22-.25-4.55-1.11-4.55-4.92 0-1.11.38-2 1.03-2.71-.1-.25-.45-1.29.1-2.64 0 0 .84-.27 2.75 1.02.79-.22 1.65-.33 2.5-.33.85 0 1.71.11 2.5.33 1.91-1.29 2.75-1.02 2.75-1.02.55 1.35.2 2.39.1 2.64.65.71 1.03 1.6 1.03 2.71 0 3.82-2.34 4.66-4.57 4.91.36.31.69.92.69 1.85V21c0 .27.16.59.67.5C19.14 20.16 22 16.42 22 12A10 10 0 0 0 12 2z"/>
					</svg>
				</a>
			{/if}
			{#if xUrl}
				<a href={xUrl} target="_blank" rel="noopener noreferrer" class="text-on-surface-variant hover:text-primary transition-colors flex items-center" title="X (Twitter)">
					<svg class="h-3.5 w-3.5 fill-current" viewBox="0 0 24 24">
						<path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z"/>
					</svg>
				</a>
			{/if}
		</div>
		<div class="font-label-sm text-label-sm text-on-surface-variant opacity-80">
			© 2024 Buildlog. Crafted for clarity.
		</div>
	</div>
</footer>
