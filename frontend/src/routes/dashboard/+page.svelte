<script lang="ts">
	import { onMount } from 'svelte';
	import { session } from '../../store';
	import { getUserAccount, hideMenu, showMenu } from '../../utils';
	import CardSkeleton from './CardSkeleton.svelte';
	import Menu from './Menu.svelte';
	import ProfileCard from './ProfileCard.svelte';

	onMount(() => {
		getUserAccount();
	});
</script>

<svelte:head>
	<title>Dashboard | Inbox Check</title>
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
	<div class="bottom"></div>
</section>

<style lang="scss">
	section {
		min-height: 100dvh;
		width: 100dvw;
		display: flex;
		flex-direction: column;
		background-color: $gray;
		background-color: #4e54c8;
		background-image: linear-gradient(to right top, #4e54c8, #8f94fb, #b4b8fd, #cad3ff, #e5f0ff);

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
			padding: 1rem;
		}
	}
</style>
