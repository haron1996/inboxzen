<script lang="ts">
	import './../global.scss';
	import Errors from './Errors.svelte';
	import Loading from './Loading.svelte';
</script>

<Errors />
<Loading />

<slot />

<style>
	@import url('https://fonts.cdnfonts.com/css/satoshi');
</style>
