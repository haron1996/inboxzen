<script lang="ts">
	import { onMount } from 'svelte';
	import { session } from '../../store';
	import { getUserEmailSettings, hideMenu, showMenu } from '../../utils';
	import CardSkeleton from './CardSkeleton.svelte';
	import Menu from './Menu.svelte';
	import ProfileCard from './ProfileCard.svelte';
	import VipDomains from './VipDomains.svelte';
	import VipEmails from './VipEmails.svelte';
	import VipKeywords from './VipKeywords.svelte';
	import DeliveryTime from './DeliveryTime.svelte';

	onMount(() => {
		setTimeout(() => {
			getUserEmailSettings();
		}, 1000);
	});
</script>

<svelte:head>
	<title>Dashboard | Zenn</title>
</svelte:head>

<svelte:window on:resize={hideMenu} />

<Menu />

<section role="none" on:click={hideMenu}>
	<div class="top">
		{#if $session && $session.email}
			<ProfileCard
				name={$session.email.account_name}
				email={$session.email.email_address}
				picture={$session.email.profile_picture}
				onClick={showMenu}
			/>
		{:else}
			<CardSkeleton height={4.5} />
		{/if}
	</div>
	<div class="bottom">
		<VipDomains />
		<VipEmails />
		<VipKeywords />
		<DeliveryTime />
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
			height: max-content;
			min-height: 10dvh;
			display: flex;
			align-items: center;
			justify-content: center;
			padding: 1rem;
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
