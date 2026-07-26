<script lang="ts">
	let { data } = $props();
</script>

<svelte:head>
	<title>Buildlog — Profile</title>
</svelte:head>

<div class="mx-auto max-w-container-max px-gutter">
	<!-- Hero Section -->
	<section class="mb-section-gap flex flex-col gap-stack-lg">
		<div class="flex items-center gap-stack-md">
			<div class="h-16 w-16 shrink-0 overflow-hidden rounded-full border border-outline-variant/30">
				<img
					class="h-full w-full object-cover grayscale transition-all duration-300 hover:grayscale-0"
					alt={data.profileData.name}
					src={data.profileData.avatarUrl}
				/>
			</div>
			<div class="h-px flex-grow bg-outline-variant/30"></div>
		</div>
		<div class="flex flex-col gap-stack-sm">
			<span
				class="font-label-md text-label-md flex items-center tracking-[0.2em] text-on-surface-variant uppercase"
			>
				<span class="mr-3 inline-block h-px w-10 bg-current"></span>{data.profileData.title}
			</span>
			<h1 class="font-display-lg text-display-lg mt-2 leading-tight text-primary">
				{data.profileData.name}
			</h1>
			<p class="font-label-sm text-label-sm text-outline italic">{data.profileData.subtitle}</p>
		</div>
		<div class="mt-stack-md max-w-[500px]">
			<p class="font-body-lg text-body-lg leading-relaxed text-primary italic">
				{data.profileData.quote}
			</p>
		</div>
	</section>

	<!-- Bio Section -->
	<section class="mb-section-gap border-l border-outline-variant/30 pl-stack-lg">
		<h2 class="font-label-md text-label-md mb-stack-md tracking-widest text-outline uppercase">
			経歴と哲学 / Biography
		</h2>
		<div class="space-y-stack-md text-on-surface">
			{#each data.profileData.bio as paragraph (paragraph)}
				<p class="font-body-md text-body-md leading-relaxed text-on-surface-variant">
					{paragraph}
				</p>
			{/each}
		</div>
	</section>

	<!-- Career Highlights -->
	<section class="mb-section-gap">
		<h2 class="font-label-md text-label-md mb-stack-lg tracking-widest text-outline uppercase">
			主要な実績 / Highlights
		</h2>
		<div class="grid grid-cols-1 gap-stack-lg">
			{#each data.profileData.highlights as item (item.title)}
				<div class="group border-b border-outline-variant/20 pb-stack-md">
					<div class="mb-2 flex items-baseline justify-between">
						<h3 class="font-headline-md text-headline-md text-primary">{item.title}</h3>
						<span class="font-label-sm text-label-sm text-on-surface-variant uppercase"
							>{item.period}</span
						>
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
	<section class="mb-section-gap">
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

</div>

