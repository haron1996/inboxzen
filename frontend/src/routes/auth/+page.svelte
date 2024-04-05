<script lang="ts">
	import { onMount } from 'svelte';
	import Spinner from '../Spinner.svelte';
	import { page } from '$app/stores';
	import { URL, errors } from '../../store';

	onMount(() => {
		const code = $page.url.searchParams.get('code');

		if (code) {
			completeGoogleAuth(code);
		}
	});

	async function completeGoogleAuth(code: string) {
		const url = `${$URL}/account/comletegoogleauth`;

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
			errors.update((errs) => [result.message, ...errs]);
			location.href = '/preauth';
			return;
		}

		location.href = '/dashboard';
	}
</script>

<svelte:head>
	<title>Continue with Google | Inboxzen</title>
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
