<script lang="ts">
	import { getUserEmailSettings, hideMenu } from '../../utils';
	import Status from './Status.svelte';
	import VipDomains from './VipDomains.svelte';
	import VipEmails from './VipEmails.svelte';
	import VipKeywords from './VipKeywords.svelte';
	import DeliveryTime from './DeliveryTime.svelte';
	import ScreenerCount from './ScreenerCount.svelte';
	import { afterNavigate } from '$app/navigation';

	afterNavigate(() => {
		getUserEmailSettings();
	});
</script>

<svelte:head>
	<title>Dashboard | Gmail Inbox Zero</title>
</svelte:head>

<svelte:window on:resize={hideMenu} />

<section role="none" on:click={hideMenu}>
	<ScreenerCount />
	<Status />
	<DeliveryTime />
	<VipDomains />
	<VipEmails />
	<VipKeywords />
</section>

<style lang="scss">
	section {
		width: 100%;
		height: 100%;
	}
</style>
