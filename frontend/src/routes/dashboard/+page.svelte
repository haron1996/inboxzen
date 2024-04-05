<script lang="ts">
	import { onMount } from 'svelte';
	import { userSession } from '../../store';
	import { hideAccountSwitcher, hideMenu, showMenu } from '../../utils';
	import Spinner from '../Spinner.svelte';
	import Accounts from './Accounts.svelte';
	import CardSkeleton from './CardSkeleton.svelte';
	import Menu from './Menu.svelte';

	function handleSectionClick() {
		hideMenu();
		hideAccountSwitcher();
	}

	onMount(() => {
		console.log('dashboard mounted');
	});
</script>

<svelte:head>
	<title>Dashboard | Inboxzen</title>
</svelte:head>

<svelte:window on:resize={showMenu} />

<Menu />
<Accounts />

<section role="none" on:click={handleSectionClick}>
	<div class="top">
		{#if $userSession && $userSession.account}
			<div
				class="profile"
				id="profile"
				role="none"
				on:click|preventDefault|stopPropagation={showMenu}
			>
				<div class="left">
					<div class="image">
						{#if $userSession.account?.profile_picture}
							<img src={$userSession.account?.profile_picture} alt="profile" />
						{:else}
							<Spinner />
						{/if}
					</div>
				</div>
				<div class="center">
					<span class="name">{$userSession.account?.account_name}</span>
					<span class="email">{$userSession.account?.email}</span>
				</div>
				<div class="right">
					<svg
						width="24px"
						height="24px"
						stroke-width="1.5"
						viewBox="0 0 24 24"
						fill="none"
						xmlns="http://www.w3.org/2000/svg"
						color="#000000"
						><path
							d="M6 9L12 15L18 9"
							stroke="#000000"
							stroke-width="1.5"
							stroke-linecap="round"
							stroke-linejoin="round"
						></path></svg
					>
				</div>
			</div>
		{:else}
			<CardSkeleton height={4.5} />
		{/if}
	</div>
	<div class="bottom"></div>
</section>

<style lang="scss">
	section {
		min-height: 100dvh;
		width: 100dvw;
		display: flex;
		flex-direction: column;
		background-color: $gray;

		.top {
			width: 100%;
			height: max-content;
			min-height: 10dvh;
			display: flex;
			align-items: center;
			justify-content: center;
			padding: 1rem;

			.profile {
				display: flex;
				align-items: center;
				justify-content: center;
				width: max-content;
				height: max-content;
				gap: 1rem;
				padding: 0.5rem 1rem;
				border-radius: 0.3rem;
				cursor: pointer;
				position: relative;
				max-width: 40rem;
				background-color: $white;
				box-shadow:
					rgba(255, 255, 255, 0.2) 0rem 0rem 0rem 0.1rem inset,
					rgba(0, 0, 0, 0.9) 0rem 0rem 0rem 0.1rem;

				.left {
					display: flex;
					align-items: center;
					justify-content: center;
					height: 3.5rem;
					width: 3.5rem;
					border-radius: 50%;
					letter-spacing: 0.1rem;

					img {
						max-inline-size: 100%;
						object-fit: cover;
						border-radius: 50%;
					}
				}

				.center {
					display: flex;
					flex-direction: column;
					justify-content: center;

					span {
						font-size: 1.3rem;
						font-weight: 500;
						font-family: $spline;
						color: $black;
						white-space: nowrap;
						overflow: hidden;
						text-overflow: ellipsis;
						text-transform: uppercase;
						max-width: 30rem;
					}

					span.email {
						color: $light-black;
						text-transform: lowercase;
					}
				}

				.right {
					display: flex;
					align-items: center;
					justify-content: center;

					svg {
						border-radius: 50%;
						border: 0.1rem solid rgba(0, 0, 0, 0.1);
						height: 2rem;
						width: 2rem;

						path {
							stroke: $light-black;
						}
					}
				}
			}
		}

		.bottom {
			min-height: 90dvh;
			width: 100%;
			padding: 1rem;
		}
	}
</style>
