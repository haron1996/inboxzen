<script lang="ts">
	import { onMount } from 'svelte';
	import BackButton from '../BackButton.svelte';
	import VipDomains from '../VipDomains.svelte';
	import VipEmails from '../VipEmails.svelte';
	import VipKeywords from '../VipKeywords.svelte';
	import { getUserEmailSettings } from '../../../utils';
	import DeliveryTime from '../DeliveryTime.svelte';

	function handleBackButton() {
		location.href = '/dashboard';
	}

	onMount(() => {
		setTimeout(() => {
			getUserEmailSettings();
		}, 1000);
	});
</script>

<svelte:head>
	<title>Account Settings | Zenn</title>
</svelte:head>

<section>
	<div class="top">
		<BackButton width={80} onClick={handleBackButton} />
	</div>
	<div class="bottom">
		<DeliveryTime />
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
			gap: 5rem;
		}
	}
</style>
