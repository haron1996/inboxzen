<script lang="ts">
	import { onMount } from 'svelte';
	import { URL, session } from '../../store';
	import { hideMenu, showMenu, updateErrorMessages } from '../../utils';
	import CardSkeleton from './CardSkeleton.svelte';
	import Menu from './Menu.svelte';
	import ProfileCard from './ProfileCard.svelte';

	async function getUserAccount() {
		const url = `${$URL}/private/getuseraccount`;

		const response = await fetch(url, {
			method: 'GET',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer'
		});

		const result = await response.json();

		if (!response.ok) {
			const message = result.message;
			switch (message) {
				case 'The access token provided is invalid or has expired. Please log in again.':
					updateErrorMessages(message);
					setTimeout(() => {
						location.href = '/';
					}, 3000);
					return;
				default:
					updateErrorMessages(message);
					return;
			}
		}

		session.set(result);
	}

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
