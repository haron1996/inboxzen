<script lang="ts">
	import './../global.scss';
	import Loading from './Loading.svelte';
	import SuccessAlerts from './SuccessAlerts.svelte';
	import ErrorAlerts from './ErrorAlerts.svelte';
</script>

<Loading />
<SuccessAlerts />
<ErrorAlerts />

<slot />

<style>
	@import url('https://fonts.cdnfonts.com/css/satoshi');
</style>
