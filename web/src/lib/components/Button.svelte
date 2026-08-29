<script lang="ts">
	import type { Snippet } from 'svelte';
	import type { HTMLButtonAttributes } from 'svelte/elements';

	type Variant = 'primary' | 'outline' | 'danger' | 'ghost';

	type Props = HTMLButtonAttributes & {
		variant?: Variant;
		children: Snippet;
	};

	let { variant = 'primary', class: className = '', children, ...rest }: Props = $props();

	const variantClasses: Record<Variant, string> = {
		primary: 'bg-primary text-on-primary hover:bg-primary/90',
		outline: 'border border-primary text-primary hover:bg-primary hover:text-on-primary',
		danger: 'border border-error text-error hover:bg-error hover:text-on-error',
		ghost: 'text-on-surface-variant hover:bg-surface-container-low hover:text-primary'
	};
</script>

<button
	{...rest}
	class={`font-label-md text-label-md inline-flex min-h-11 items-center justify-center rounded-lg px-5 py-2 transition-colors focus-visible:outline focus-visible:outline-2 focus-visible:outline-primary disabled:cursor-wait disabled:opacity-60 ${variantClasses[variant]} ${className}`}
>
	{@render children()}
</button>
