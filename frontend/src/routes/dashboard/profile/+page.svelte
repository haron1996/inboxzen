<script lang="ts">
	import { onMount } from 'svelte';
	import { loading, session } from '../../../store';
	import CardSkeleton from '../CardSkeleton.svelte';
	import { getUserAccount, scrollPageToTop } from '../../../utils';
	import BackButton from '../BackButton.svelte';
	import Button from '../../Button.svelte';
	import ProfileCard from '../ProfileCard.svelte';

	function handleBackButton() {
		location.href = '/dashboard';
	}

	onMount(() => {
		scrollPageToTop();
		getUserAccount();
	});
</script>

<svelte:head>
	<title>Profile Settings | Zenn</title>
</svelte:head>

<section>
	<div class="top">
		<BackButton width={80} onClick={handleBackButton} />
	</div>
	<div class="bottom">
		<!-- Beginning of primary account -->
		<div class="card">
			<div class="title">
				<span>PRIMARY ACCOUNT</span>
			</div>

			{#if $loading}
				<CardSkeleton height={10} width={80} padding={0} borderRadius={0} />
			{:else if $session}
				<div class="content">
					{#if $session.email}
						{#if $session.email.primaryaccount}
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
								borderRadius={0.3}
								color="rgb(255, 255, 255)"
								padding={0.5}
								text="delete account"
								onClick={() => {}}
							/>
						{:else if $session.emails}
							{#each $session.emails as email}
								{#if email.primaryaccount}
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
										borderRadius={0.3}
										color="rgb(255, 255, 255)"
										padding={0.5}
										text="delete account"
										onClick={() => {}}
									/>
								{/if}
							{/each}
						{/if}
					{/if}
				</div>
			{/if}
		</div>

		<!-- End of primary account -->

		<div class="secondary-accounts">
			<div class="title">
				<span>SECONDARY ACCOUNTS</span>
			</div>
			<div class="contents">
				{#if $loading}
					<CardSkeleton height={10} width={80} padding={0} borderRadius={0} />
				{:else if $session}
					{#if $session.email}
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
									borderRadius={0.3}
									color="rgb(255, 255, 255)"
									padding={0.5}
									text="delete account"
									onClick={() => {}}
								/>
							</div>
						{/if}
					{/if}
					{#if $session.emails}
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
										borderRadius={0.3}
										color="rgb(255, 255, 255)"
										padding={0.5}
										text="delete account"
										onClick={() => {}}
									/>
								</div>
							{/if}
						{/each}
					{/if}
				{/if}
			</div>
		</div>

		<!-- End of secondary accounts -->

		<div class="card">
			<div class="title">
				<span>ACTIVE ACCOUNT</span>
			</div>
			{#if $loading}
				<CardSkeleton height={10} width={80} padding={0} borderRadius={0} />
			{:else}
				<div class="content">
					{#if $session}
						{#if $session.email}
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
								borderRadius={0.3}
								color="rgb(255, 255, 255)"
								padding={0.5}
								text="delete account"
								onClick={() => {}}
							/>
						{/if}
					{/if}
				</div>
			{/if}
		</div>
		<!-- End of active account account -->
	</div>
</section>

<style lang="scss">
	section {
		min-height: 100dvh;
		width: 100dvw;
		display: flex;
		flex-direction: column;
		gap: 3rem;
		background-color: #4e54c8;
		background-image: linear-gradient(to right top, #4e54c8, #8f94fb, #b4b8fd, #cad3ff, #e5f0ff);

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
