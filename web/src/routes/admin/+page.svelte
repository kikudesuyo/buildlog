<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { deleteDiary } from '$lib/api/client';
	import DiaryFeed from '$lib/components/DiaryFeed.svelte';
	let { data } = $props();

	async function handleDelete(id: number) {
		try {
			await deleteDiary(id);
		} catch {
			return false;
		}
		return true;
	}
</script>

<svelte:head><title>Buildlog — Admin Diary</title></svelte:head>
<DiaryFeed entries={data.diaryEntries} isAdmin onEdit={(id) => goto(resolve(`/admin/diary/${id}/edit`))} onDelete={handleDelete} />
