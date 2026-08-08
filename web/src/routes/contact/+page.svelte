<script lang="ts">
	let { data } = $props();
</script>

<svelte:head>
	<title>Contact — Buildlog</title>
</svelte:head>

<div class="mx-auto max-w-container-max px-gutter">
	<section class="max-w-2xl border-t border-outline-variant/30 pt-section-gap">
		<p class="font-label-md text-label-md mb-stack-md tracking-widest text-outline uppercase">Contact</p>
		<h1 class="font-display-lg text-display-lg text-primary">お問い合わせ</h1>
		<p class="font-body-lg text-body-lg mt-stack-md leading-relaxed text-on-surface-variant">
			ご連絡は、以下のメールアドレスまたはソーシャルアカウントからお願いします。
		</p>

		<div class="mt-section-gap flex flex-col gap-4">
			<a
				href={`mailto:${data.profileData.contactEmail}`}
				class="flex min-h-11 items-center gap-3 rounded-lg border border-outline-variant/30 px-4 py-3 text-body-md text-on-surface-variant transition-colors hover:border-primary hover:text-primary"
			>
				<span class="material-symbols-outlined">mail</span>
				{data.profileData.contactEmail}
			</a>
			{#if data.profileData.githubUrl}
				<a href={data.profileData.githubUrl} target="_blank" rel="noopener noreferrer" class="flex min-h-11 items-center gap-3 rounded-lg border border-outline-variant/30 px-4 py-3 text-body-md text-on-surface-variant transition-colors hover:border-primary hover:text-primary">
					GitHub
				</a>
			{/if}
			{#if data.profileData.xUrl}
				<a href={data.profileData.xUrl} target="_blank" rel="noopener noreferrer" class="flex min-h-11 items-center gap-3 rounded-lg border border-outline-variant/30 px-4 py-3 text-body-md text-on-surface-variant transition-colors hover:border-primary hover:text-primary">
					X (Twitter)
				</a>
			{/if}
		</div>
	</section>
</div>
