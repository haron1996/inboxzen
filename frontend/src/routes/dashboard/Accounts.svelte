<script lang="ts">
	import { page } from '$app/stores';
	import { userSession } from '../../store';
	import { hideAccountSwitcher, showMenu } from '../../utils';
	import Spinner from '../Spinner.svelte';

	function handleBackClick() {
		hideAccountSwitcher();
		showMenu();
	}
</script>

<div id="accountSwitcher">
	<div class="back" role="none" on:click={handleBackClick}>
		<svg
			width="24px"
			height="24px"
			viewBox="0 0 24 24"
			stroke-width="1.5"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			color="#000000"
			><path
				d="M11 6L5 12L11 18"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path><path
				d="M19 6L13 12L19 18"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path></svg
		>
		<span>Menu</span>
	</div>
	{#if $userSession.accounts}
		{#each $userSession.accounts as account}
			<div class="account">
				<div class="left">
					<div class="image">
						{#if account.profile_picture}
							<img src={account.profile_picture} alt="profile" />
						{:else}
							<Spinner />
						{/if}
					</div>
				</div>
				<div class="center">
					<span class="name">{account.account_name}</span>
					<span class="email">{account.email}</span>
				</div>
			</div>
		{/each}
	{/if}
	<div class="add-account" role="none">
		<svg
			width="24px"
			height="24px"
			stroke-width="1.5"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			color="#000000"
			><path
				d="M9 12H12M15 12H12M12 12V9M12 12V15"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path><path
				d="M11.7 1.1732C11.8856 1.06603 12.1144 1.06603 12.3 1.17321L21.2263 6.3268C21.4119 6.43397 21.5263 6.63205 21.5263 6.84641V17.1536C21.5263 17.3679 21.4119 17.566 21.2263 17.6732L12.3 22.8268C12.1144 22.934 11.8856 22.934 11.7 22.8268L2.77372 17.6732C2.58808 17.566 2.47372 17.3679 2.47372 17.1536V6.84641C2.47372 6.63205 2.58808 6.43397 2.77372 6.32679L11.7 1.1732Z"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path></svg
		>
		<span>Add New Account</span>
	</div>
</div>

<style lang="scss">
	#accountSwitcher {
		position: absolute;
		background-color: $white;
		box-shadow:
			rgba(255, 255, 255, 0.2) 0rem 0rem 0rem 0.1rem inset,
			rgba(0, 0, 0, 0.9) 0rem 0rem 0rem 0.1rem;
		display: none;
		flex-direction: column;
		min-height: max-content;
		cursor: pointer;
		border-radius: 0.3rem;

		.back {
			display: flex;
			gap: 1rem;
			padding: 1rem;
			border-radius: 0.3rem;
			cursor: pointer;
			position: relative;
			background-color: $white;
			align-items: center;
			border-bottom: 0.1rem solid rgb(0 0 0 / 0.1);
			min-height: 4.5rem;

			svg {
				height: 2rem;
				width: 2rem;

				path {
					stroke: $black-1;
				}
			}

			span {
				font-size: 1.2rem;
				font-weight: 500;
				font-family: $spline;
				color: $black;
				white-space: nowrap;
				overflow: hidden;
				text-overflow: ellipsis;
				text-transform: uppercase;
				max-width: 30rem;
			}

			&:hover {
				background-color: $gray;
			}
		}

		.account {
			display: flex;
			gap: 1rem;
			padding: 1rem;
			border-radius: 0.3rem;
			cursor: pointer;
			position: relative;
			background-color: $white;
			min-height: 4.5rem;
			align-items: center;

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

			&:hover {
				background-color: $gray;
			}
		}

		.add-account {
			display: flex;
			gap: 1rem;
			padding: 1rem;
			border-radius: 0.3rem;
			cursor: pointer;
			position: relative;
			background-color: $white;
			align-items: center;
			border-top: 0.1rem solid rgb(0 0 0 / 0.1);
			min-height: 4.5rem;

			svg {
				height: 2rem;
				width: 2rem;

				path {
					stroke: $black-1;
				}
			}

			span {
				font-size: 1.2rem;
				font-family: $spline;
				color: $black;
				white-space: nowrap;
				overflow: hidden;
				text-overflow: ellipsis;
				text-transform: uppercase;
				max-width: 30rem;
				font-weight: 500;
			}

			&:hover {
				background-color: $gray;
			}
		}
	}
</style>
