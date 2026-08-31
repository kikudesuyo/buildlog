<script lang="ts">
	let { data } = $props();
</script>

<svelte:head><title>目標履歴 — Buildlog</title></svelte:head>

<section class="mx-auto max-w-4xl px-gutter" aria-labelledby="goal-history-heading">
	<header class="mb-8">
		<h1 id="goal-history-heading" class="font-headline-lg text-headline-lg text-primary">過去の目標</h1>
		<p class="font-body-md text-body-md mt-2 text-on-surface-variant">月ごとの目標と達成状況を振り返れます。</p>
	</header>
	{#if data.goalHistory.length === 0}
		<p class="font-body-md text-body-md text-on-surface-variant">過去の目標はまだありません。</p>
	{:else}
		<div class="grid gap-4 md:grid-cols-2">
			{#each data.goalHistory as period (period.startsAt)}
				<article class="rounded-xl border border-outline-variant/30 bg-surface-container-lowest p-5">
					<h2 class="font-headline-md text-headline-md text-primary">{period.startsAt.slice(0, 7)}</h2>
					<ul class="mt-4 space-y-3">
						{#each period.goals as goal (goal.id)}
							<li>
								<div class="flex items-center justify-between gap-3">
									<span class="font-body-md text-body-md text-on-surface">{goal.title}</span>
									<span class="font-label-md text-label-md text-primary">{goal.progressValue} / {goal.targetValue}</span>
								</div>
								<div class="mt-2 h-2 overflow-hidden rounded-full bg-outline-variant/20"><div class="h-full bg-primary" style={`width: ${Math.min((goal.progressValue / goal.targetValue) * 100, 100)}%`}></div></div>
							</li>
						{/each}
					</ul>
				</article>
			{/each}
		</div>
	{/if}
</section>
