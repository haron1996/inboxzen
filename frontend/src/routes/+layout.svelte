<script lang="ts">
	import './../global.scss';
	import SuccessAlerts from './SuccessAlerts.svelte';
	import ErrorAlerts from './ErrorAlerts.svelte';
</script>

<SuccessAlerts />
<ErrorAlerts />

<slot />

<style>
	@import url('https://fonts.googleapis.com/css2?family=Jost:ital,wght@0,100..900;1,100..900&display=swap');
</style>
