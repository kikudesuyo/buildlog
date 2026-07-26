<script lang="ts">
	import DiaryFeed from '$lib/components/DiaryFeed.svelte';
	let { data } = $props();
</script>

<svelte:head><title>Essence — Diary</title></svelte:head>
<DiaryFeed entries={data.diaryEntries} />
