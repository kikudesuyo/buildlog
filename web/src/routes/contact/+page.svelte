<script lang="ts">
	let { data } = $props();

	const contactTypes = ['仕事の依頼', 'サイトの不具合', 'その他'];
	let contactType = $state(contactTypes[0]);
	let message = $state('');

	function handleSubmit(event: SubmitEvent) {
		event.preventDefault();

		const subject = `お問い合わせ（${contactType}）`;
		const body = `お問い合わせ内容：${contactType}\n\n${message}`;
		window.location.href = `mailto:${data.profileData.contactEmail}?subject=${encodeURIComponent(subject)}&body=${encodeURIComponent(body)}`;
	}
</script>

<svelte:head>
	<title>Contact — Buildlog</title>
</svelte:head>

<div class="mx-auto max-w-container-max px-gutter">
	<section class="max-w-2xl">
		<p class="font-label-md text-label-md mb-stack-md tracking-widest text-outline uppercase">Contact</p>
		<h1 class="font-display-lg text-display-lg text-primary">お問い合わせ</h1>
		<p class="font-body-lg text-body-lg mt-stack-md leading-relaxed text-on-surface-variant">
			お問い合わせ内容を入力してください。送信ボタンを押すと、お使いのメールソフトが起動します。
		</p>

		<form class="mt-section-gap flex flex-col gap-6" onsubmit={handleSubmit}>
			<div class="flex flex-col gap-2">
				<label for="contact-type" class="font-label-md text-label-md font-bold text-on-surface">
					お問い合わせ内容
				</label>
				<select
					id="contact-type"
					bind:value={contactType}
					class="min-h-11 rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-body-md text-on-surface focus:border-primary focus:outline-none"
				>
					{#each contactTypes as type (type)}
						<option value={type}>{type}</option>
					{/each}
				</select>
			</div>

			<div class="flex flex-col gap-2">
				<label for="message" class="font-label-md text-label-md font-bold text-on-surface">内容</label>
				<textarea
					id="message"
					bind:value={message}
					required
					rows="8"
					placeholder="お問い合わせ内容を入力してください"
					class="resize-y rounded-lg border border-outline-variant bg-surface-container-high px-3 py-2 text-body-md text-on-surface focus:border-primary focus:outline-none"
				></textarea>
			</div>

			<button
				type="submit"
				class="min-h-11 self-start rounded-lg bg-primary px-5 py-2 font-label-md text-label-md font-bold text-on-primary transition-opacity hover:opacity-80"
			>
				メールで送信
			</button>
		</form>

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
