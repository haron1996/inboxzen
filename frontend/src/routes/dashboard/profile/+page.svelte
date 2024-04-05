<script lang="ts">
	import { onMount } from 'svelte';
	import { userSession } from '../../../store';
	import Spinner from '../../Spinner.svelte';
	import CardSkeleton from '../CardSkeleton.svelte';
	import { scrollPageToTop } from '../../../utils';
	import BackButton from '../BackButton.svelte';
	import Button from '../../Button.svelte';

	function goToPreAuth() {
		localStorage.setItem('request', 'add');
		location.href = '/preauth';
	}

	function handleBackButton() {
		location.href = '/dashboard';
	}

	function deleteAccount() {}

	onMount(() => {
		scrollPageToTop();
	});
</script>

<svelte:head>
	<title>Account Settings | Inboxzen</title>
</svelte:head>

<section>
	<div class="top">
		<BackButton width={80} onClick={handleBackButton} />
	</div>
	<div class="bottom">
		<div class="primary-account">
			<span class="title">ACTIVE ACCOUNT</span>
			<div class="account-wrapper">
				{#if $userSession.account}
					<div class="account">
						<div class="left">
							<div class="image">
								{#if $userSession.account.profile_picture}
									<img src={$userSession.account.profile_picture} alt="profile" />
								{:else}
									<Spinner />
								{/if}
							</div>
						</div>
						<div class="center">
							<span class="name">{$userSession.account.account_name}</span>
							<span class="email">{$userSession.account.email}</span>
						</div>
					</div>
				{:else}
					<CardSkeleton height={4.5} width={30} padding={1} borderRadius={0.5} />
				{/if}
				<!-- <button>
					<span>DELETE ACCOUNT</span>
				</button> -->
				<Button
					height={4.5}
					width={15}
					backgroundColor="#E74646"
					borderRadius={0.3}
					color="rgb(255, 255, 255)"
					padding={0.5}
					text="delete account"
					onClick={deleteAccount}
				/>
			</div>
		</div>

		<div class="primary-account">
			<span class="title">PRIMARY ACCOUNT</span>
			<div class="account-wrapper">
				{#if $userSession.accounts}
					{#each $userSession.accounts as account}
						{#if account.primaryaccount}
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
						{/if}
					{/each}
				{:else}
					<CardSkeleton height={4.5} width={30} padding={1} borderRadius={0.5} />
				{/if}
				<Button
					height={4.5}
					width={15}
					backgroundColor="#E74646"
					borderRadius={0.3}
					color="rgb(255, 255, 255)"
					padding={0.5}
					text="delete account"
					onClick={deleteAccount}
				/>
			</div>
		</div>
		<div class="secondary-accounts">
			<span class="title">SECONDARY ACCOUNTS</span>
			{#if $userSession.accounts}
				{#each $userSession.accounts as account}
					<div class="account-wrapper">
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
						<Button
							height={4.5}
							width={15}
							backgroundColor="#E74646"
							borderRadius={0.3}
							color="rgb(255, 255, 255)"
							padding={0.5}
							text="delete account"
							onClick={deleteAccount}
						/>
					</div>
				{/each}
			{:else}
				<CardSkeleton height={4.5} width={75} padding={1} borderRadius={0.5} />
			{/if}
		</div>
		<Button
			height={4}
			width={80}
			backgroundColor="#525FE1"
			borderRadius={0.3}
			color="rgb(255, 255, 255)"
			padding={0.5}
			text="add new account"
			onClick={goToPreAuth}
		/>
	</div>
</section>

<style lang="scss">
	section {
		min-height: 100dvh;
		width: 100dvw;
		display: flex;
		flex-direction: column;
		background-color: $gray;
		gap: 3rem;

		.top {
			min-height: 10dvh;
			display: flex;
			align-items: center;
			padding: 2rem;
			justify-content: center;
		}

		.bottom {
			width: 100%;
			min-height: 90dvh;
			display: flex;
			flex-direction: column;
			align-items: center;
			padding: 2rem;
			gap: 5rem;

			.primary-account {
				display: flex;
				flex-direction: column;
				width: 80rem;
				border: 0.1rem solid $black-4;
				box-shadow: rgba(0, 0, 0, 0.16) 0rem 0.1rem 0.4rem;
				background-color: $white;
				border-radius: 0.3rem;
				padding: 2rem;
				gap: 3rem;
				min-height: max-content;

				span.title {
					font-family: $spline;
					font-size: 1.2rem;
					text-transform: uppercase;
					border-bottom: 0.1rem solid rgb(0 0 0 / 0.1);
					height: 3rem;
					font-weight: 600;
					color: $black;
				}

				.account-wrapper {
					display: flex;
					align-items: center;
					justify-content: space-between;

					.account {
						display: flex;
						gap: 1rem;
						padding: 1rem;
						border-radius: 0.5rem;
						cursor: pointer;
						position: relative;
						background-color: $white;
						min-height: 4.5rem;
						align-items: center;
						width: 30rem;
						background-color: $gray;

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
					}
				}
			}

			.secondary-accounts {
				display: flex;
				flex-direction: column;
				width: 80rem;
				border: 0.1rem solid $black-4;
				box-shadow: rgba(0, 0, 0, 0.16) 0rem 0.1rem 0.4rem;
				background-color: $white;
				border-radius: 0.3rem;
				padding: 2rem;
				gap: 3rem;
				min-height: max-content;

				span.title {
					font-family: $spline;
					font-size: 1.2rem;
					text-transform: uppercase;
					border-bottom: 0.1rem solid rgb(0 0 0 / 0.1);
					height: 3rem;
					font-weight: 600;
					color: $black;
				}

				.account-wrapper {
					display: flex;
					align-items: center;
					justify-content: space-between;

					.account {
						display: flex;
						gap: 1rem;
						padding: 1rem;
						border-radius: 0.5rem;
						cursor: pointer;
						position: relative;
						background-color: $white;
						min-height: 4.5rem;
						align-items: center;
						width: 30rem;
						background-color: $gray;

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
					}
				}
			}
		}
	}
</style>
