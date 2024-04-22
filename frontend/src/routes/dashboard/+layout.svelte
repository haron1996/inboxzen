<script lang="ts">
	import Menu from './Menu.svelte';
	import { hideMenu, moveEmailsToInbox } from '../../utils';
	import { afterNavigate } from '$app/navigation';
	import TopBar from './TopBar.svelte';
	import { browser } from '$app/environment';

	afterNavigate(() => {
		hideMenu();
	});

	function checkTimeAndRun() {
		const now = new Date();
		if (now.getHours() === 9 && now.getMinutes() === 0 && now.getSeconds() <= 59) {
			moveEmailsToInbox();
		}
	}

	// if (browser) {
	// 	setInterval(checkTimeAndRun, 30000);
	// }
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
