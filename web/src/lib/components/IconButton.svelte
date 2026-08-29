<script lang="ts">
	import type { Snippet } from 'svelte';
	import type { HTMLButtonAttributes } from 'svelte/elements';

	type Variant = 'default' | 'danger';

	type Props = HTMLButtonAttributes & {
		variant?: Variant;
		icon: string;
		element?: HTMLButtonElement;
		children?: Snippet;
	};

	let { variant = 'default', icon, element = $bindable<HTMLButtonElement | undefined>(), children, class: className = '', ...rest }: Props = $props();
</script>

<button
	bind:this={element}
	{...rest}
	class={`inline-flex min-h-11 min-w-11 items-center justify-center rounded-lg p-2 text-on-surface-variant transition-colors hover:bg-surface-container-high ${variant === 'danger' ? 'hover:text-error hover:bg-error/10' : 'hover:text-primary'} ${className}`}
>
	<span class="material-symbols-outlined text-[18px]" aria-hidden="true">{icon}</span>
	{#if children}{@render children()}{/if}
</button>
