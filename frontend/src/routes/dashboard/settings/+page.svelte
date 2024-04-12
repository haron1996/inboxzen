<script lang="ts">
	import { onMount } from 'svelte';
	import { URL, session } from '../../../store';
	import BackButton from '../BackButton.svelte';
	import VipDomains from '../VipDomains.svelte';
	import VipEmails from '../VipEmails.svelte';
	import VipKeywords from '../VipKeywords.svelte';
	import { getUserEmailSettings, scrollPageToTop, updateErrorMessages } from '../../../utils';
	import DeliveryTime from '../DeliveryTime.svelte';

	function handleBackButton() {
		location.href = '/dashboard';
	}

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
		scrollPageToTop();
		//getUserAccount();
		getUserEmailSettings();
	});
</script>

<svelte:head>
	<title>SETTINGS | ZENN.EMAIL</title>
</svelte:head>

<section>
	<div class="top">
		<BackButton width={80} onClick={handleBackButton} />
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
		background-color: #4e54c8;
		background-image: linear-gradient(to right top, #4e54c8, #8f94fb, #b4b8fd, #cad3ff, #e5f0ff);

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
