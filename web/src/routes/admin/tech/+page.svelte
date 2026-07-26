<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { deleteTech } from '$lib/api/client';
	import TechFeed from '$lib/components/TechFeed.svelte';
	let { data } = $props();

	async function handleDelete(id: number) {
		if (!confirm('この記事を削除してもよろしいですか？')) return;
		try {
			await deleteTech(id);
		} catch {
			alert('削除に失敗しました。');
			return false;
		}
		return true;
	}
</script>

<svelte:head><title>Essence — Admin Tech Feed</title></svelte:head>
<TechFeed
	featuredArticle={data.featuredArticle}
	techArticles={data.techArticles}
	isAdmin
	onEdit={(id) => goto(resolve(`/admin/tech/${id}/edit`))}
	onDelete={handleDelete}
/>
