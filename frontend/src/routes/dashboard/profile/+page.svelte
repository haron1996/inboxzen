<script lang="ts">
	import { onMount } from 'svelte';
	import { URL, session } from '../../../store';
	import CardSkeleton from '../CardSkeleton.svelte';
	import { scrollPageToTop, updateErrorMessages } from '../../../utils';
	import BackButton from '../BackButton.svelte';
	import Button from '../../Button.svelte';
	import ProfileCard from '../ProfileCard.svelte';

	function handleBackButton() {
		location.href = '/dashboard';
	}

	async function getUserAccount() {
		const url = `${$URL}/private/getuseraccount`;

		const response = await fetch(url, {
			method: 'GET',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer'
		});

		const result = await response.json();

		if (!response.ok) {
			const message = result.message;
			switch (message) {
				case 'The access token provided is invalid or has expired. Please log in again.':
					updateErrorMessages(message);
					setTimeout(() => {
						location.href = '/';
					}, 3000);
					return;
				default:
					updateErrorMessages(message);
					return;
			}
		}

		session.set(result);
	}

	onMount(() => {
		scrollPageToTop();
		getUserAccount();
	});
</script>

<svelte:head>
	<title>Account Settings | Inbox Check</title>
</svelte:head>

<section>
	<div class="top">
		<BackButton width={80} onClick={handleBackButton} />
	</div>
	<div class="bottom">
		{#if $session.email && $session.emails}
			{#if $session.email.primaryaccount}
				<div class="card">
					<div class="title">
						<span>PRIMARY ACCOUNT</span>
					</div>
					<div class="content">
						<ProfileCard
							name={$session.email.account_name}
							email={$session.email.email_address}
							picture={$session.email.profile_picture}
							width="30rem"
							SVGDisplay="none"
							onClick={() => {}}
						/>
						<Button
							height={4.5}
							width={15}
							backgroundColor="#E74646"
							borderRadius={0.3}
							color="rgb(255, 255, 255)"
							padding={0.5}
							text="delete account"
							onClick={() => {}}
						/>
					</div>
				</div>
			{:else}
				{#each $session.emails as email}
					{#if email.primaryaccount}
						<div class="card">
							<div class="title">
								<span>PRIMARY ACCOUNT</span>
							</div>
							<div class="content">
								<ProfileCard
									name={email.account_name}
									email={email.email_address}
									picture={email.profile_picture}
									width="30rem"
									SVGDisplay="none"
									onClick={() => {}}
								/>
								<Button
									height={4.5}
									width={15}
									backgroundColor="#E74646"
									borderRadius={0.3}
									color="rgb(255, 255, 255)"
									padding={0.5}
									text="delete account"
									onClick={() => {}}
								/>
							</div>
						</div>
					{/if}
				{/each}
			{/if}
		{:else}
			<CardSkeleton height={20} width={80} padding={0} borderRadius={0.5} />
		{/if}

		{#if $session.email && $session.emails}
			<div class="secondary-accounts">
				<div class="title">
					<span>SECONDARY ACCOUNTS</span>
				</div>
				<div class="contents">
					{#if !$session.email.primaryaccount}
						<div class="content">
							<ProfileCard
								name={$session.email.account_name}
								email={$session.email.email_address}
								picture={$session.email.profile_picture}
								width="30rem"
								SVGDisplay="none"
								onClick={() => {}}
							/>
							<Button
								height={4.5}
								width={15}
								backgroundColor="#E74646"
								borderRadius={0.3}
								color="rgb(255, 255, 255)"
								padding={0.5}
								text="delete account"
								onClick={() => {}}
							/>
						</div>
					{/if}
					{#each $session.emails as email}
						{#if !email.primaryaccount}
							<div class="content">
								<ProfileCard
									name={email.account_name}
									email={email.email_address}
									picture={email.profile_picture}
									width="30rem"
									SVGDisplay="none"
									onClick={() => {}}
								/>
								<Button
									height={4.5}
									width={15}
									backgroundColor="#E74646"
									borderRadius={0.3}
									color="rgb(255, 255, 255)"
									padding={0.5}
									text="delete account"
									onClick={() => {}}
								/>
							</div>
						{/if}
					{/each}
				</div>
			</div>
		{:else}
			<CardSkeleton height={40} width={80} padding={0} borderRadius={0.5} />
		{/if}

		{#if $session.email}
			<div class="card">
				<div class="title">
					<span>ACTIVE ACCOUNT</span>
				</div>
				<div class="content">
					<ProfileCard
						name={$session.email.account_name}
						email={$session.email.email_address}
						picture={$session.email.profile_picture}
						width="30rem"
						SVGDisplay="none"
						onClick={() => {}}
					/>
					<Button
						height={4.5}
						width={15}
						backgroundColor="#E74646"
						borderRadius={0.3}
						color="rgb(255, 255, 255)"
						padding={0.5}
						text="delete account"
						onClick={() => {}}
					/>
				</div>
			</div>
		{:else}
			<CardSkeleton height={20} width={80} padding={0} borderRadius={0.5} />
		{/if}
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

			.card {
				display: flex;
				flex-direction: column;
				width: 80rem;
				min-height: max-content;
				border: 0.1rem solid $black-4;
				box-shadow: rgba(0, 0, 0, 0.16) 0rem 0.1rem 0.4rem;
				background-color: $white;
				border-radius: 0.3rem;

				.title {
					min-height: 5rem;
					display: flex;
					align-items: center;
					border-bottom: 0.1rem solid rgba(0, 0, 0, 0.1);

					span {
						font-family: $spline;
						font-size: 1.1rem;
						text-transform: uppercase;
						font-weight: 500;
						padding: 1rem;
					}
				}

				.content {
					min-height: 10rem;
					height: 1rem;
					display: flex;
					align-items: center;
					justify-content: space-between;
					padding: 1rem;
				}
			}

			.secondary-accounts {
				display: flex;
				flex-direction: column;
				width: 80rem;
				min-height: max-content;
				border: 0.1rem solid $black-4;
				box-shadow: rgba(0, 0, 0, 0.16) 0rem 0.1rem 0.4rem;
				background-color: $white;
				border-radius: 0.3rem;

				.title {
					min-height: 5rem;
					display: flex;
					align-items: center;
					border-bottom: 0.1rem solid rgba(0, 0, 0, 0.1);

					span {
						font-family: $spline;
						font-size: 1.1rem;
						text-transform: uppercase;
						font-weight: 500;
						padding: 1rem;
					}
				}

				.contents {
					display: flex;
					flex-direction: column;

					.content {
						min-height: 10rem;
						height: 1rem;
						display: flex;
						align-items: center;
						justify-content: space-between;
						padding: 1rem;
					}
				}
			}
		}
	}
</style>
