<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { deleteDiary, saveCurrentGoals } from '$lib/api/client';
	import type { GoalPeriod } from '$lib/api/types';
	import DiaryFeed from '$lib/components/DiaryFeed.svelte';
	let { data } = $props();
	let goals = $state<GoalPeriod>(data.goals);
	let isGoalDialogOpen = $state(goals.goals.length === 0);
	let isSavingGoals = $state(false);
	let goalError = $state('');
	let draftGoals = $state(
		(goals.goals.length > 0 ? goals.goals : [{ title: '', targetValue: 1, progressValue: 0 }]).map((goal) => ({
			title: goal.title,
			targetValue: goal.targetValue,
			progressValue: goal.progressValue
		}))
	);

	async function handleDelete(id: number) {
		try {
			await deleteDiary(id);
		} catch {
			return false;
		}
		return true;
	}

	function openGoalDialog() {
		draftGoals = goals.goals.map((goal) => ({
			title: goal.title,
			targetValue: goal.targetValue,
			progressValue: goal.progressValue
		}));
		isGoalDialogOpen = true;
	}

	function openGoalDialogWithKeyboard(event: KeyboardEvent) {
		if (event.key !== 'Enter' && event.key !== ' ') return;
		event.preventDefault();
		openGoalDialog();
	}

	function addGoal() {
		draftGoals = [...draftGoals, { title: '', targetValue: 1, progressValue: 0 }];
	}

	function removeGoal(index: number) {
		if (draftGoals.length === 1) return;
		draftGoals = draftGoals.filter((_, goalIndex) => goalIndex !== index);
	}

	async function submitGoals() {
		goalError = '';
		const normalizedGoals = draftGoals.map((goal) => ({
			title: goal.title.trim(),
			targetValue: Number(goal.targetValue),
			progressValue: Math.min(Number(goal.progressValue), Number(goal.targetValue))
		}));
		if (normalizedGoals.some((goal) => !goal.title || goal.targetValue < 1 || goal.progressValue < 0)) {
			goalError = '目標名と1以上の目標値を入力してください。';
			return;
		}
		isSavingGoals = true;
		try {
			goals = await saveCurrentGoals(normalizedGoals);
			isGoalDialogOpen = false;
		} catch {
			goalError = '目標を保存できませんでした。';
		} finally {
			isSavingGoals = false;
		}
	}
</script>

<svelte:head><title>Buildlog — Admin Diary</title></svelte:head>
<section class="editorial-container mx-auto mb-10 px-gutter" aria-labelledby="goals-heading">
	<div class="mb-4 flex items-center justify-between gap-4">
		<div>
			<h1 id="goals-heading" class="font-headline-lg text-headline-lg text-primary">今月の目標</h1>
			<p class="font-body-sm text-body-sm text-on-surface-variant">{goals.startsAt} — {goals.endsAt}</p>
		</div>
		<button type="button" onclick={openGoalDialog} class="font-label-md text-label-md min-h-11 rounded-lg border border-primary px-4 py-2 text-primary hover:bg-primary hover:text-on-primary">目標を編集</button>
	</div>
	<div class="grid gap-4 md:grid-cols-2">
		{#each goals.goals as goal (goal.id)}
			{@const percentage = Math.round((goal.progressValue / goal.targetValue) * 100)}
			<div
				class="cursor-pointer rounded-xl border border-outline-variant/30 bg-surface-container-lowest p-4 transition-colors hover:border-primary/50 hover:bg-surface-container-low"
				role="button"
				tabindex="0"
				aria-label={`${goal.title}を編集`}
				onclick={openGoalDialog}
				onkeydown={openGoalDialogWithKeyboard}
			>
				<div class="mb-2 flex items-center justify-between gap-3">
					<h2 class="font-label-md text-label-md font-semibold text-primary">{goal.title}</h2>
					<span class="font-label-sm text-label-sm text-on-surface-variant">{goal.progressValue} / {goal.targetValue} ({percentage}%)</span>
				</div>
				<div class="h-3 overflow-hidden rounded-full bg-surface-container-high" role="progressbar" aria-label={`${goal.title}の進捗`} aria-valuemin="0" aria-valuemax={goal.targetValue} aria-valuenow={goal.progressValue}>
					<div class="h-full rounded-full bg-primary transition-all" style={`width: ${Math.min(percentage, 100)}%`}></div>
				</div>
			</div>
		{:else}
			<p class="font-body-md text-body-md text-on-surface-variant">今月の目標はまだ登録されていません。</p>
		{/each}
	</div>
</section>

{#if isGoalDialogOpen}
	<div class="fixed inset-0 z-50 flex items-start justify-center overflow-y-auto bg-primary/20 px-4 py-24 backdrop-blur-xs" role="presentation">
		<section role="dialog" aria-modal="true" aria-labelledby="goal-dialog-title" class="w-full max-w-2xl rounded-xl border border-outline-variant/30 bg-surface-container-lowest p-6 shadow-xl">
			<div class="mb-6 flex items-start justify-between gap-4">
				<div>
					<h2 id="goal-dialog-title" class="font-headline-lg text-headline-lg text-primary">今月の目標{goals.goals.length > 0 ? 'を編集' : 'を登録'}</h2>
					<p class="font-body-sm text-body-sm mt-1 text-on-surface-variant">目標値と現在の進捗を入力してください。</p>
				</div>
				{#if goals.goals.length > 0}<button type="button" onclick={() => (isGoalDialogOpen = false)} class="min-h-11 min-w-11 rounded-lg text-on-surface-variant hover:bg-surface-container-low" aria-label="目標入力を閉じる">×</button>{/if}
			</div>
			<div class="space-y-4">
				{#each draftGoals as goal, index (index)}
					<div class="rounded-lg border border-outline-variant/30 p-4">
						<div class="mb-3 flex items-center justify-between gap-3">
							<span class="font-label-sm text-label-sm text-outline">目標 {index + 1}</span>
							{#if draftGoals.length > 1}<button type="button" onclick={() => removeGoal(index)} class="font-label-sm text-label-sm min-h-11 px-2 text-error" aria-label={`目標${index + 1}を削除`}>削除</button>{/if}
						</div>
						<label class="font-label-sm text-label-sm mb-1 block text-on-surface-variant" for={`goal-title-${index}`}>目標名</label>
						<input id={`goal-title-${index}`} bind:value={goal.title} class="font-body-md text-body-md mb-3 min-h-11 w-full rounded-lg border border-outline-variant/40 bg-surface px-3 py-2 text-primary" placeholder="例: 個人開発" />
						<div class="grid gap-3 sm:grid-cols-2">
							<div><label class="font-label-sm text-label-sm mb-1 block text-on-surface-variant" for={`goal-target-${index}`}>目標値</label><input id={`goal-target-${index}`} type="number" min="1" bind:value={goal.targetValue} class="font-body-md text-body-md min-h-11 w-full rounded-lg border border-outline-variant/40 bg-surface px-3 py-2 text-primary" /></div>
							<div><label class="font-label-sm text-label-sm mb-1 block text-on-surface-variant" for={`goal-progress-${index}`}>現在の進捗: {goal.progressValue}</label><input id={`goal-progress-${index}`} type="range" min="0" max={goal.targetValue || 1} bind:value={goal.progressValue} class="mt-3 min-h-11 w-full accent-primary" /></div>
						</div>
					</div>
				{/each}
			</div>
			{#if goalError}<p class="font-body-sm text-body-sm mt-4 text-error" role="alert">{goalError}</p>{/if}
			<div class="mt-6 flex flex-wrap justify-between gap-3">
				<button type="button" onclick={addGoal} class="font-label-md text-label-md min-h-11 rounded-lg border border-outline-variant/50 px-4 py-2 text-primary">+ 目標を追加</button>
				<div class="flex gap-3"><button type="button" onclick={() => (isGoalDialogOpen = false)} class="font-label-md text-label-md min-h-11 rounded-lg px-4 py-2 text-on-surface-variant" disabled={isSavingGoals}>キャンセル</button><button type="button" onclick={submitGoals} class="font-label-md text-label-md min-h-11 rounded-lg bg-primary px-5 py-2 text-on-primary" disabled={isSavingGoals}>{isSavingGoals ? '保存中…' : '保存する'}</button></div>
			</div>
		</section>
	</div>
{/if}

<DiaryFeed entries={data.diaryEntries} isAdmin onEdit={(id) => goto(resolve(`/admin/diary/${id}/edit`))} onDelete={handleDelete} />
