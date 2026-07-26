<script lang="ts">
	import { beforeNavigate } from '$app/navigation';
	import { onMount } from 'svelte';

	type Props = {
		isDirty: boolean;
		isSubmitting: boolean;
	};

	let { isDirty, isSubmitting }: Props = $props();
	const message = '保存されていない変更があります。このページを離れますか？';

	beforeNavigate(({ cancel }) => {
		if (isDirty && !isSubmitting && !confirm(message)) cancel();
	});

	onMount(() => {
		function handleBeforeUnload(event: BeforeUnloadEvent) {
			if (!isDirty || isSubmitting) return;
			event.preventDefault();
			event.returnValue = '';
		}

		window.addEventListener('beforeunload', handleBeforeUnload);
		return () => window.removeEventListener('beforeunload', handleBeforeUnload);
	});
</script>
