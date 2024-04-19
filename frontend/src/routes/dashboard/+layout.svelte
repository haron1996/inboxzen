<script lang="ts">
	import Menu from './Menu.svelte';
	import { hideMenu } from '../../utils';
	import { afterNavigate } from '$app/navigation';
	import TopBar from './TopBar.svelte';

	afterNavigate(() => {
		hideMenu();
	});
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
