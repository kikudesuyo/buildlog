<script lang="ts">
	let { form } = $props();
</script>

<svelte:head>
	<title>認証 — Buildlog</title>
</svelte:head>

<div class="mx-auto flex min-h-[60vh] max-w-md items-center px-gutter">
	<section class="w-full rounded-2xl border border-outline-variant/30 bg-surface-container-lowest p-6 shadow-sm md:p-8" aria-labelledby="auth-heading">
		<p class="font-label-md text-label-md mb-3 tracking-[0.2em] text-outline uppercase">Private workspace</p>
		<h1 id="auth-heading" class="font-display-lg text-display-lg text-primary">認証</h1>
		<p class="font-body-md text-body-md mt-3 text-on-surface-variant">管理画面を開くにはパスワードを入力してください。</p>

		<form method="POST" class="mt-8 flex flex-col gap-4">
			<label for="password" class="font-label-md text-label-md font-semibold text-on-surface">パスワード</label>
			<input
				id="password"
				name="password"
				type="password"
				autocomplete="current-password"
				required
				class="font-body-md text-body-md min-h-11 rounded-lg border border-outline-variant/40 bg-surface px-3 py-2 text-primary outline-none focus:border-primary"
			/>
			{#if form?.error}<p class="font-body-sm text-body-sm text-error" role="alert">{form.error}</p>{/if}
			<button type="submit" class="font-label-md text-label-md min-h-11 rounded-lg bg-primary px-5 py-2 font-semibold text-on-primary transition-opacity hover:opacity-90">認証する</button>
		</form>
	</section>
</div>
