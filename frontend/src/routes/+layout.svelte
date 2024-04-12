<script lang="ts">
	import './../global.scss';
	import SuccessAlerts from './SuccessAlerts.svelte';
	import ErrorAlerts from './ErrorAlerts.svelte';
</script>

<SuccessAlerts />
<ErrorAlerts />

<slot />
