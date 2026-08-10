<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { deleteDiary, saveCurrentGoals } from '$lib/api/client';
	import type { GoalPeriod } from '$lib/api/types';
	import DiaryFeed from '$lib/components/DiaryFeed.svelte';
	let { data } = $props();
	let goals = $state<GoalPeriod>(data.goals);
	let isGoalDialogOpen = $state(false);
	let isSavingGoals = $state(false);
	let goalError = $state('');
	let editingGoalId = $state<number | null>(null);
	let isAddingGoal = $state(false);
	let draftGoal = $state({ title: '', targetValue: 1, progressValue: 0 });

	async function handleDelete(id: number) {
		try {
			await deleteDiary(id);
		} catch {
			return false;
		}
		return true;
	}

	function openGoalDialog(goal?: GoalPeriod['goals'][number]) {
		editingGoalId = goal?.id ?? null;
		isAddingGoal = !goal;
		draftGoal = goal
			? { title: goal.title, targetValue: goal.targetValue, progressValue: goal.progressValue }
			: { title: '', targetValue: 1, progressValue: 0 };
		goalError = '';
		isGoalDialogOpen = true;
	}

	async function removeGoal(goal: GoalPeriod['goals'][number]) {
		if (goals.goals.length === 1) {
			goalError = '最後の目標は削除できません。';
			return;
		}
		if (!confirm(`「${goal.title}」を削除しますか？`)) return;
		await persistGoals(goals.goals.filter((item) => item.id !== goal.id));
	}

	async function persistGoals(goalList: Array<{ title: string; targetValue: number; progressValue: number }>) {
		isSavingGoals = true;
		try {
			goals = await saveCurrentGoals(goalList);
			isGoalDialogOpen = false;
		} catch {
			goalError = '目標を保存できませんでした。';
		} finally {
			isSavingGoals = false;
		}
	}

	async function submitGoal() {
		goalError = '';
		const normalizedGoal = {
			title: draftGoal.title.trim(),
			targetValue: Number(draftGoal.targetValue),
			progressValue: Math.min(Number(draftGoal.progressValue), Number(draftGoal.targetValue))
		};
		if (!normalizedGoal.title || normalizedGoal.targetValue < 1 || normalizedGoal.progressValue < 0) {
			goalError = '目標名と1以上の目標値を入力してください。';
			return;
		}
		const nextGoals = isAddingGoal
			? [...goals.goals, normalizedGoal]
			: goals.goals.map((goal) => (goal.id === editingGoalId ? normalizedGoal : goal));
		await persistGoals(nextGoals);
	}

	function updateGoalProgress(goalId: number, progressValue: number) {
		goals = {
			...goals,
			goals: goals.goals.map((goal) => (goal.id === goalId ? { ...goal, progressValue } : goal))
		};
	}

	async function saveGoalProgress(goalId: number, progressValue: number) {
		const nextGoals = goals.goals.map((goal) =>
			goal.id === goalId ? { ...goal, progressValue } : goal
		);
		await persistGoals(nextGoals);
	}

	function closeGoalDialog() {
		if (!isSavingGoals) {
			isGoalDialogOpen = false;
			goalError = '';
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
		<button type="button" onclick={() => openGoalDialog()} class="font-label-md text-label-md inline-flex min-h-11 min-w-32 shrink-0 items-center justify-center whitespace-nowrap rounded-lg border border-primary px-4 py-2 text-primary hover:bg-primary hover:text-on-primary">目標を追加</button>
	</div>
	<div class="grid gap-4 md:grid-cols-2">
		{#each goals.goals as goal (goal.id)}
			{@const percentage = Math.round((goal.progressValue / goal.targetValue) * 100)}
			{@const isCompleted = goal.progressValue >= goal.targetValue}
			<div
				class="rounded-xl border {isCompleted ? 'border-primary/40 bg-primary/5' : 'border-outline-variant/30 bg-surface-container-lowest'} p-4 transition-colors hover:border-primary/50 hover:bg-surface-container-low"
			>
				<div class="mb-3 flex items-start justify-between gap-3">
					<div>
						<h2 class="font-label-md text-label-md font-semibold text-primary">{goal.title}</h2>
						<p class="font-label-sm text-label-sm mt-1 text-on-surface-variant">{goal.progressValue} / {goal.targetValue} ({Math.min(percentage, 100)}%)</p>
					</div>
					<div class="flex items-center gap-2">
						{#if isCompleted}
							<span class="font-label-sm text-label-sm flex items-center gap-1 rounded-full bg-primary/10 px-2.5 py-1 font-semibold text-primary" role="status">
								<span class="material-symbols-outlined text-[16px]">check_circle</span>
								達成
							</span>
						{/if}
						<button type="button" onclick={() => openGoalDialog(goal)} class="min-h-11 min-w-11 rounded-lg p-2 text-on-surface-variant hover:bg-surface-container-high hover:text-primary" aria-label={`${goal.title}を編集`}><span class="material-symbols-outlined text-[18px]">edit</span></button>
						<button type="button" onclick={() => removeGoal(goal)} class="min-h-11 min-w-11 rounded-lg p-2 text-on-surface-variant hover:bg-error/10 hover:text-error" aria-label={`${goal.title}を削除`}><span class="material-symbols-outlined text-[18px]">delete</span></button>
					</div>
				</div>
				<div class="flex flex-col gap-1.5">
					<input
						type="range"
						min="0"
						max={goal.targetValue}
						value={goal.progressValue}
						oninput={(event) => updateGoalProgress(goal.id, Number(event.currentTarget.value))}
						onchange={(event) => saveGoalProgress(goal.id, Number(event.currentTarget.value))}
						class="h-3 w-full cursor-pointer accent-primary"
						aria-label={`${goal.title}の進捗`}
						aria-valuemin="0"
						aria-valuemax={goal.targetValue}
						aria-valuenow={goal.progressValue}
					/>
					<p class="font-label-sm text-label-sm text-on-surface-variant">バーを動かして進捗を更新</p>
				</div>
				{#if isCompleted}<p class="font-body-sm text-body-sm mt-3 flex items-center gap-1 text-primary"><span class="material-symbols-outlined text-[17px]">celebration</span>目標達成おめでとうございます！</p>{/if}
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
					<h2 id="goal-dialog-title" class="font-headline-lg text-headline-lg text-primary">{isAddingGoal ? '目標を追加' : '目標を編集'}</h2>
					<p class="font-body-sm text-body-sm mt-1 text-on-surface-variant">目標値と現在の進捗を入力してください。</p>
				</div>
				<button type="button" onclick={closeGoalDialog} class="min-h-11 min-w-11 rounded-lg text-on-surface-variant hover:bg-surface-container-low" aria-label="目標入力を閉じる">×</button>
			</div>
			<div class="space-y-4">
				<div class="rounded-lg border border-outline-variant/30 p-4">
					<label class="font-label-sm text-label-sm mb-1 block text-on-surface-variant" for="goal-title">目標名</label>
					<input id="goal-title" bind:value={draftGoal.title} class="font-body-md text-body-md mb-3 min-h-11 w-full rounded-lg border border-outline-variant/40 bg-surface px-3 py-2 text-primary" placeholder="例: 個人開発" />
					<div class="grid gap-3 sm:grid-cols-2">
						<div><label class="font-label-sm text-label-sm mb-1 block text-on-surface-variant" for="goal-target">目標値</label><input id="goal-target" type="number" min="1" bind:value={draftGoal.targetValue} class="font-body-md text-body-md min-h-11 w-full rounded-lg border border-outline-variant/40 bg-surface px-3 py-2 text-primary" /></div>
						<div><label class="font-label-sm text-label-sm mb-1 block text-on-surface-variant" for="goal-progress">現在の進捗: {draftGoal.progressValue}</label><input id="goal-progress" type="range" min="0" max={draftGoal.targetValue || 1} bind:value={draftGoal.progressValue} class="mt-3 min-h-11 w-full accent-primary" /></div>
					</div>
				</div>
			</div>
			{#if goalError}<p class="font-body-sm text-body-sm mt-4 text-error" role="alert">{goalError}</p>{/if}
			<div class="mt-6 flex flex-wrap justify-between gap-3">
				<div class="flex w-full justify-end gap-3"><button type="button" onclick={closeGoalDialog} class="font-label-md text-label-md min-h-11 rounded-lg px-4 py-2 text-on-surface-variant" disabled={isSavingGoals}>キャンセル</button><button type="button" onclick={submitGoal} class="font-label-md text-label-md min-h-11 rounded-lg bg-primary px-5 py-2 text-on-primary" disabled={isSavingGoals}>{isSavingGoals ? '保存中…' : '保存する'}</button></div>
			</div>
		</section>
	</div>
{/if}

<DiaryFeed entries={data.diaryEntries} isAdmin onEdit={(id) => goto(resolve(`/admin/diary/${id}/edit`))} onDelete={handleDelete} />
