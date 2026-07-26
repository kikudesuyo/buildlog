<script lang="ts">
	import TechFeed from '$lib/components/TechFeed.svelte';
	let { data } = $props();
</script>

<svelte:head><title>Buildlog — Tech Feed</title></svelte:head>
<TechFeed {...data} />
