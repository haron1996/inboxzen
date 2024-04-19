<script lang="ts">
	import { session } from '../../store';
	import { hideMenu, showMenu } from '../../utils';
	import CardSkeleton from './CardSkeleton.svelte';
	import ProfileCard from './ProfileCard.svelte';
</script>

<div class="top-bar" role="none" on:click={hideMenu}>
	{#if $session && $session.email}
		<ProfileCard
			name={$session.email.account_name}
			email={$session.email.email_address}
			picture={$session.email.profile_picture}
			onClick={showMenu}
		/>
	{:else}
		<CardSkeleton height={5} />
	{/if}
</div>

<style lang="scss">
	div {
		width: inherit;
		height: max-content;
		height: 10dvh;
		display: flex;
		align-items: center;
		justify-content: center;
		position: fixed;
		top: 0;
		left: 0;
		background-color: inherit;
	}
</style>
