<script lang="ts">
	import { onMount } from 'svelte';
	import Spinner from '../Spinner.svelte';
	import { page } from '$app/stores';
	import { finishAuth, updateErrorMessages } from '../../utils';
	import { goto } from '$app/navigation';

	let error: string | null = '';

	onMount(() => {
		error = $page.url.searchParams.get('error');

		if (error) {
			updateErrorMessages(error);
			goto('/preauth');
			return;
		}

		const code = $page.url.searchParams.get('code');

		if (code) {
			finishAuth(code);
		}
	});
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
