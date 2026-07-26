<script lang="ts">
	import { resolve } from '$app/paths';
</script>

<!-- Footer -->
<footer class="bg-surface w-full py-section-gap border-t border-outline-variant/20">
	<div class="max-w-container-max mx-auto flex flex-col md:flex-row justify-between items-center gap-stack-md px-gutter">
		<div class="font-label-md text-label-md text-primary font-semibold">
			Buildlog
		</div>
		<div class="flex gap-stack-lg">
			<a href="#rss" class="font-label-sm text-label-sm text-on-surface-variant hover:text-primary transition-colors">RSS</a>
			<a href={resolve('/profile')} class="font-label-sm text-label-sm text-on-surface-variant hover:text-primary transition-colors">Contact</a>
		</div>
		<div class="font-label-sm text-label-sm text-on-surface-variant opacity-80">
			© 2024 Buildlog. Crafted for clarity.
		</div>
	</div>
</footer>
