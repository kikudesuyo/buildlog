<script lang="ts">
	import PostCalendar from '$lib/components/PostCalendar.svelte';
	let { data } = $props();
	let avatarFailed = $state(false);

	function handleAvatarError() {
		avatarFailed = true;
	}
</script>

<svelte:head>
	<title>Buildlog — Profile</title>
</svelte:head>

<div class="mx-auto w-full max-w-5xl px-gutter">
	<!-- Hero Section -->
	<section class="mb-section-gap flex flex-col gap-stack-lg">
		<div class="flex items-center gap-stack-md">
			<div class="h-16 w-16 shrink-0 overflow-hidden rounded-full border border-outline-variant/30">
				{#if avatarFailed}
					<span class="flex h-full w-full items-center justify-center bg-surface-container-low text-headline-md text-outline" aria-label="プロフィール画像なし">{data.profileData.name.slice(0, 1)}</span>
				{:else}
					<img
						class="aspect-square h-full w-full object-cover grayscale transition-all duration-300 hover:grayscale-0"
						alt={data.profileData.name}
						src="/profile.jpg"
						onerror={handleAvatarError}
						loading="lazy"
					/>
				{/if}
			</div>
			<div class="flex min-w-0 flex-col gap-stack-xs">
				<h1 class="font-display-lg text-display-lg leading-tight text-primary">
					{data.profileData.name}
				</h1>
				<span
					class="font-label-md text-label-md flex items-center tracking-[0.2em] text-on-surface-variant uppercase"
				>
					<span class="mr-3 inline-block h-px w-10 shrink-0 bg-current"></span>{data.profileData.title}
				</span>
			</div>
		</div>
		<div class="mt-stack-md w-full max-w-3xl">
			<p class="font-body-lg text-body-lg break-words leading-relaxed text-primary italic">
				{data.profileData.quote}
			</p>
		</div>
	</section>

	<!-- Bio Section -->
	<section id="biography" class="mb-section-gap scroll-mt-24 border-l border-outline-variant/30 pl-stack-lg">
		<h2 class="font-label-md text-label-md mb-stack-md tracking-widest text-outline uppercase">
			自己紹介 / self-introduction
		</h2>
		<div class="max-w-3xl space-y-stack-md text-on-surface">
			{#each data.profileData.bio as paragraph (paragraph)}
				<p class="font-body-md text-body-md leading-relaxed text-on-surface-variant">
					{paragraph}
				</p>
			{/each}
		</div>
	</section>

	<!-- Career Highlights -->
	<section id="highlights" class="mb-section-gap scroll-mt-24">
		<h2 class="font-label-md text-label-md mb-stack-lg tracking-widest text-outline uppercase">
			主要な実績 / Highlights
		</h2>
		<div class="grid grid-cols-1 gap-stack-lg">
			{#each data.profileData.highlights as item (item.title)}
				<div class="group border-b border-outline-variant/20 pb-stack-md">
					<div class="mb-2 flex flex-col gap-1">
						<span class="font-label-sm text-label-sm text-on-surface-variant uppercase">{item.period}</span>
						<h3 class="font-headline-md text-headline-md text-primary">{item.title}</h3>
					</div>
					<p class="font-body-md text-body-md max-w-[600px] text-on-surface-variant">
						{item.description}
					</p>
				</div>
			{/each}
			{#if data.profileData.award}
				<div class="pt-stack-sm">
					<div class="inline-flex items-center gap-stack-sm text-primary">
						<span class="material-symbols-outlined text-xl">workspace_premium</span>
						<span class="font-label-md text-label-md font-semibold">{data.profileData.award}</span>
					</div>
				</div>
			{/if}
			</div>
	</section>

	<!-- Skills & Expertise -->
	<section id="expertise" class="mb-section-gap scroll-mt-24">
		<h2 class="font-label-md text-label-md mb-stack-md tracking-widest text-outline uppercase">
			専門領域 / Expertise
		</h2>
		<div class="flex flex-wrap gap-x-12 gap-y-stack-sm">
			{#each data.profileData.expertise as skill (skill)}
				<span
					class="font-body-md text-body-md cursor-default border-b border-transparent pb-1 text-on-surface transition-colors hover:border-primary/30 hover:text-primary"
				>
					{skill}
				</span>
			{/each}
		</div>
	</section>

	<PostCalendar history={data.postHistory} />
</div>
