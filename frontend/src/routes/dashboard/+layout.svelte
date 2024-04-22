<script lang="ts">
	import Menu from './Menu.svelte';
	import { hideMenu, moveEmailsToInbox } from '../../utils';
	import { afterNavigate } from '$app/navigation';
	import TopBar from './TopBar.svelte';
	import { browser } from '$app/environment';
	import { session } from '../../store';

	afterNavigate(() => {
		hideMenu();
	});

	function handleMoveEmailsToInbox() {
		if ($session.email?.running) {
			moveEmailsToInbox();
		}
	}

	if (browser) {
		setInterval(handleMoveEmailsToInbox, 10000);
	}
</script>

<Menu />

<div class="app">
	<TopBar />
	<menu>
		<slot />
	</menu>
</div>

<style lang="scss">
	.app {
		height: 100dvh;
		width: 100dvw;
		background-color: #f9f9f9;

		menu {
			position: fixed;
			top: 10dvh;
			left: 0;
			width: inherit;
			height: 90dvh;
			background-color: inherit;
			overflow-y: auto;
		}
	}
</style>
