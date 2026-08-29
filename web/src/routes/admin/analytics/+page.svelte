<script lang="ts">
	import { resolve } from '$app/paths';

	let { data } = $props();
	const analytics = $derived(data.analyticsData);

	// ランキング表示用のタブ管理 ('views' または 'likes')
	let activeTab = $state<'views' | 'likes'>('views');

	// グラフ描画用の最大値を計算
	const maxActivityCount = $derived(
		Math.max(...analytics.monthlyActivities.map((a) => a.count), 1)
	);

	const currentRanking = $derived(
		activeTab === 'views' ? analytics.topViewsArticles : analytics.topLikesArticles
	);
</script>

<svelte:head>
	<title>Buildlog — Analytics Dashboard</title>
</svelte:head>

<div class="flex flex-col gap-8 p-6">
	<!-- ヘッダー -->
	<header class="flex flex-col gap-2">
		<h1 class="font-headline-lg text-headline-lg font-bold text-primary flex items-center gap-3">
			<span class="material-symbols-outlined text-3xl">analytics</span>
			Analytics Dashboard
		</h1>
		<p class="font-body-md text-body-md text-outline">
			コンテンツ全体の閲覧数、いいね数、投稿推移、および人気記事の統計情報です。
		</p>
	</header>

	<!-- サマリーカード -->
	<section class="grid grid-cols-1 md:grid-cols-3 gap-6">
		<!-- 総閲覧数 -->
		<div class="flex flex-col justify-between p-6 rounded-2xl border border-outline-variant/30 bg-surface-container-lowest shadow-sm hover:shadow-md transition-shadow">
			<div class="flex items-center justify-between mb-4">
				<span class="font-label-md text-label-md font-semibold text-outline uppercase tracking-wider">総閲覧数</span>
				<span class="material-symbols-outlined text-primary text-2xl bg-primary/10 p-2 rounded-xl">visibility</span>
			</div>
			<div>
				<div class="font-headline-lg text-headline-lg font-bold text-on-surface">
					{analytics.totalViews.toLocaleString()} <span class="font-body-sm text-body-sm text-outline">views</span>
				</div>
			</div>
		</div>

		<!-- 総いいね数 -->
		<div class="flex flex-col justify-between p-6 rounded-2xl border border-outline-variant/30 bg-surface-container-lowest shadow-sm hover:shadow-md transition-shadow">
			<div class="flex items-center justify-between mb-4">
				<span class="font-label-md text-label-md font-semibold text-outline uppercase tracking-wider">総いいね数</span>
				<span class="material-symbols-outlined text-primary text-2xl bg-primary/10 p-2 rounded-xl">favorite</span>
			</div>
			<div>
				<div class="font-headline-lg text-headline-lg font-bold text-on-surface">
					{analytics.totalLikes.toLocaleString()} <span class="font-body-sm text-body-sm text-outline">likes</span>
				</div>
			</div>
		</div>

		<!-- 総投稿数と比率 -->
		<div class="flex flex-col justify-between p-6 rounded-2xl border border-outline-variant/30 bg-surface-container-lowest shadow-sm hover:shadow-md transition-shadow">
			<div class="flex items-center justify-between mb-4">
				<span class="font-label-md text-label-md font-semibold text-outline uppercase tracking-wider">総コンテンツ数</span>
				<span class="material-symbols-outlined text-primary text-2xl bg-primary/10 p-2 rounded-xl">article</span>
			</div>
			<div>
				<div class="font-headline-lg text-headline-lg font-bold text-on-surface mb-2">
					{analytics.totalPosts} <span class="font-body-sm text-body-sm text-outline">posts</span>
				</div>
				<!-- プログレスバー風にTech/Diary比率を表示 -->
				<div class="flex flex-col gap-1.5 mt-2">
					<div class="flex h-2 w-full overflow-hidden rounded-full bg-outline-variant/20">
						<div
							class="bg-primary transition-all"
							style="width: {analytics.totalPosts > 0 ? (analytics.techCount / analytics.totalPosts) * 100 : 0}%"
							title="Tech Articles"
						></div>
						<div
							class="bg-secondary transition-all"
							style="width: {analytics.totalPosts > 0 ? (analytics.diaryCount / analytics.totalPosts) * 100 : 0}%"
							title="Diaries"
						></div>
					</div>
					<div class="flex items-center justify-between text-[11px] font-medium text-outline">
						<span class="flex items-center gap-1">
							<span class="h-1.5 w-1.5 rounded-full bg-primary"></span>
							Tech: {analytics.techCount}
						</span>
						<span class="flex items-center gap-1">
							<span class="h-1.5 w-1.5 rounded-full bg-secondary"></span>
							Diary: {analytics.diaryCount}
						</span>
					</div>
				</div>
			</div>
		</div>
	</section>

	<div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
		<!-- 月間投稿活動推移 (グラフ) -->
		<section class="lg:col-span-2 p-6 rounded-2xl border border-outline-variant/30 bg-surface-container-lowest flex flex-col gap-6 shadow-sm">
			<h2 class="font-title-md text-title-md font-bold text-on-surface">過去12ヶ月の活動推移 / Monthly Activity</h2>
			
			<div class="flex flex-col justify-end h-64 w-full mt-4">
				<div class="flex items-end justify-between h-48 w-full gap-2 border-b border-outline-variant/20 pb-1">
					{#each analytics.monthlyActivities as activity (activity.month)}
						<div class="flex flex-col items-center flex-1 group">
							<!-- ツールチップ -->
							<span class="opacity-0 group-hover:opacity-100 transition-opacity bg-on-surface text-surface-container-lowest text-[10px] font-bold px-1.5 py-0.5 rounded shadow mb-1.5 absolute -translate-y-8 select-none pointer-events-none z-10">
								{activity.count} posts
							</span>
							
							<!-- バー本体 -->
							<div
								class="w-full rounded-t-md bg-primary-container group-hover:bg-primary transition-all duration-300"
								style="height: {Math.max((activity.count / maxActivityCount) * 100, 4)}%"
							></div>
						</div>
					{/each}
				</div>
				<!-- 月ラベル列 -->
				<div class="flex justify-between w-full mt-2 text-[10px] text-outline font-medium">
					{#each analytics.monthlyActivities as activity (activity.month)}
						<span class="flex-1 text-center truncate">{activity.month.substring(5)}月</span>
					{/each}
				</div>
			</div>
		</section>

		<!-- 記事ランキング -->
		<section class="p-6 rounded-2xl border border-outline-variant/30 bg-surface-container-lowest flex flex-col gap-6 shadow-sm">
			<header class="flex flex-col gap-4">
				<h2 class="font-title-md text-title-md font-bold text-on-surface">人気コンテンツ / Content Ranking</h2>
				<!-- タブ選択 -->
				<div class="flex rounded-lg bg-surface-container-low p-1 w-full border border-outline-variant/10">
					<button
						type="button"
						onclick={() => (activeTab = 'views')}
						class="flex-1 py-1.5 rounded-md font-label-md text-label-md font-bold transition-all text-center cursor-pointer {activeTab === 'views' ? 'bg-surface-container-lowest text-primary shadow-sm' : 'text-outline hover:text-on-surface'}"
					>
						閲覧数
					</button>
					<button
						type="button"
						onclick={() => (activeTab = 'likes')}
						class="flex-1 py-1.5 rounded-md font-label-md text-label-md font-bold transition-all text-center cursor-pointer {activeTab === 'likes' ? 'bg-surface-container-lowest text-primary shadow-sm' : 'text-outline hover:text-on-surface'}"
					>
						いいね数
					</button>
				</div>
			</header>

			<!-- ランキングリスト -->
			<div class="flex flex-col gap-3.5 overflow-y-auto max-h-[300px] pr-1">
				{#if currentRanking.length === 0}
					<p class="text-center text-outline text-body-sm py-8">データがありません</p>
				{:else}
					{#each currentRanking as item, i (item.id)}
						<div class="flex items-center justify-between gap-4 p-2.5 rounded-xl border border-outline-variant/10 bg-surface-container-low/30 hover:bg-surface-container-low/60 transition-colors">
							<div class="flex items-center gap-3 truncate">
								<!-- 順位バッジ -->
								<span class="flex items-center justify-center h-6 w-6 shrink-0 rounded-full font-bold text-xs {i === 0 ? 'bg-primary text-on-primary' : i === 1 ? 'bg-secondary text-on-secondary' : 'bg-outline-variant/30 text-on-surface-variant'}">
									{i + 1}
								</span>
								<div class="flex flex-col truncate">
									<a
										href={item.type === 'tech' ? resolve('/tech') : resolve(`/diary/${item.id}`)}
										class="font-body-md text-body-md font-semibold text-on-surface hover:text-primary hover:underline transition-all truncate"
									>
										{item.title}
									</a>
									<span class="text-[10px] font-bold text-outline uppercase tracking-wider mt-0.5">
										{item.type}
									</span>
								</div>
							</div>
							<div class="flex items-center gap-1.5 shrink-0 text-body-sm font-semibold text-primary">
								{#if activeTab === 'views'}
									<span class="material-symbols-outlined text-[16px]">visibility</span>
									<span>{item.views}</span>
								{:else}
									<span class="material-symbols-outlined text-[16px] text-error">favorite</span>
									<span>{item.likes}</span>
								{/if}
							</div>
						</div>
					{/each}
				{/if}
			</div>
		</section>
	</div>
</div>
