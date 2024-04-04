<script lang="ts">
	import BackButton from '../BackButton.svelte';
	import VipDomains from '../VipDomains.svelte';
	import VipEmails from '../VipEmails.svelte';
	import VipKeywords from '../VipKeywords.svelte';

	function handleBackButton() {
		location.href = '/dashboard';
	}
</script>

<svelte:head>
	<title>Settings | Inboxzen</title>
</svelte:head>

<section>
	<div class="top">
		<BackButton width={80} onClick={handleBackButton} />
	</div>
	<div class="bottom">
		<VipDomains />
		<VipEmails />
		<VipKeywords />
	</div>
</section>

<style lang="scss">
	section {
		width: 100dvw;
		min-height: 100dvh;
		display: flex;
		flex-direction: column;
		background-color: $gray;
		gap: 4rem;
		padding: 2rem;

		.top {
			width: 100%;
			min-height: 10dvh;
			display: flex;
			align-items: center;
			padding: 2rem;
			justify-content: center;
		}

		.bottom {
			min-height: 90dvh;
			width: 100%;
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			gap: 5rem;
		}
	}
</style>
