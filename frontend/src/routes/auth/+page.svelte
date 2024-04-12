<script lang="ts">
	import { onMount } from 'svelte';
	import Spinner from '../Spinner.svelte';
	import { page } from '$app/stores';
	import { finishAuth } from '../../utils';

	onMount(() => {
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
