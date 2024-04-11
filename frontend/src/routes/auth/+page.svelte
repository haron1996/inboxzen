<script lang="ts">
	import { onMount } from 'svelte';
	import Spinner from '../Spinner.svelte';
	import { page } from '$app/stores';
	import { URL } from '../../store';
	import { updateErrorMessages } from '../../utils';

	onMount(() => {
		const code = $page.url.searchParams.get('code');

		if (code) {
			completeGoogleAuth(code);
		}
	});

	async function completeGoogleAuth(code: string) {
		const url = `${$URL}/user/comletegoogleauth`;

		const response = await fetch(url, {
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify({ code: code })
		});

		const result = await response.json();

		if (!response.ok) {
			const message = result.message;
			updateErrorMessages(message);
			setTimeout(() => {
				location.href = '/preauth';
			}, 3000);
			return;
		}

		location.href = '/dashboard';
	}
</script>

<svelte:head>
	<title>Continue with Google | Inbox Check</title>
</svelte:head>

<div class="wrapper">
	<Spinner />
</div>

<style lang="scss">
	.wrapper {
		min-height: 100dvh;
		width: 100dvw;
		display: flex;
		align-items: center;
		justify-content: center;
	}
</style>
