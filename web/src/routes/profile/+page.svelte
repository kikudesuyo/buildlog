<script lang="ts">
	let { data } = $props();
	let copied = $state(false);

	function copyEmail() {
		navigator.clipboard.writeText(data.profileData.contactEmail);
		copied = true;
		setTimeout(() => {
			copied = false;
		}, 2000);
	}
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

	<!-- Contact & SNS Section -->
	<section class="mb-section-gap border-t border-outline-variant/30 pt-stack-lg">
		<h2 class="font-label-md text-label-md mb-stack-md tracking-widest text-outline uppercase">
			連絡先と繋がり / Contact & Socials
		</h2>
		<div class="flex flex-col gap-6 md:flex-row md:items-center md:gap-12 mt-6">
			<!-- Email Link -->
			<button
				type="button"
				onclick={copyEmail}
				class="flex items-center gap-3 font-body-md text-body-md text-on-surface-variant hover:text-primary transition-colors duration-200 cursor-pointer bg-transparent border-none p-0 text-left"
				title="メールアドレスをコピー"
			>
				<span class="material-symbols-outlined text-[20px]">
					{copied ? 'check' : 'mail'}
				</span>
				<span class="flex items-center gap-2">
					{data.profileData.contactEmail}
					{#if copied}
						<span class="text-primary font-label-sm text-label-sm tracking-wider">(コピーしました)</span>
					{/if}
				</span>
			</button>

			<!-- Social Links -->
			<div class="flex items-center gap-6">
				{#if data.profileData.githubUrl}
					<a
						href={data.profileData.githubUrl}
						target="_blank"
						rel="noopener noreferrer"
						class="flex items-center gap-2 font-body-md text-body-md text-on-surface-variant hover:text-primary transition-colors duration-200"
						title="GitHub"
					>
						<svg class="h-5 w-5 fill-current" viewBox="0 0 24 24">
							<path d="M12 2A10 10 0 0 0 2 12c0 4.42 2.87 8.17 6.84 9.5.5.08.66-.23.66-.5v-1.69c-2.77.6-3.36-1.34-3.36-1.34-.46-1.16-1.11-1.47-1.11-1.47-.9-.62.07-.6.07-.6 1 .07 1.53 1.03 1.53 1.03.9 1.52 2.34 1.07 2.91.83.09-.65.35-1.09.63-1.34-2.22-.25-4.55-1.11-4.55-4.92 0-1.11.38-2 1.03-2.71-.1-.25-.45-1.29.1-2.64 0 0 .84-.27 2.75 1.02.79-.22 1.65-.33 2.5-.33.85 0 1.71.11 2.5.33 1.91-1.29 2.75-1.02 2.75-1.02.55 1.35.2 2.39.1 2.64.65.71 1.03 1.6 1.03 2.71 0 3.82-2.34 4.66-4.57 4.91.36.31.69.92.69 1.85V21c0 .27.16.59.67.5C19.14 20.16 22 16.42 22 12A10 10 0 0 0 12 2z"/>
						</svg>
						GitHub
					</a>
				{/if}

				{#if data.profileData.xUrl}
					<a
						href={data.profileData.xUrl}
						target="_blank"
						rel="noopener noreferrer"
						class="flex items-center gap-2 font-body-md text-body-md text-on-surface-variant hover:text-primary transition-colors duration-200"
						title="X (Twitter)"
					>
						<svg class="h-4 w-4 fill-current" viewBox="0 0 24 24">
							<path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z"/>
						</svg>
						X
					</a>
				{/if}
			</div>
		</div>
	</section>

</div>

