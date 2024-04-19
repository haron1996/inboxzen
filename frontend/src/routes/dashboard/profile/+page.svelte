<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import { loading, session } from '../../../store';
	import { getUserAccount } from '../../../utils';
	import CardSkeleton from '../CardSkeleton.svelte';
	import ProfileCard from '../ProfileCard.svelte';

	afterNavigate(() => {
		getUserAccount();
	});
</script>

<svelte:head>
	<title>Profile Settings | Gmail Inbox Zero</title>
</svelte:head>

<section>
	<!--Active accout-->
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
					{/if}
				{/if}
			</div>
		{/if}
	</div>
	<!-- End of active account account -->

	<!-- Primary account -->
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
							{/if}
						{/each}
					{/if}
				{/if}
			</div>
		{/if}
	</div>

	<!-- End of primary account -->

	<!--Secondary accounts-->
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
							</div>
						{/if}
					{/each}
				{:else}
					<div class="no-content">
						<p>No secondary emails</p>
					</div>
				{/if}
			{/if}
		</div>
	</div>

	<!-- End of secondary accounts -->
</section>

<style lang="scss">
	section {
		width: 100%;
		height: 100%;

		.card {
			display: flex;
			flex-direction: column;
			width: 70%;
			margin: 1rem auto;
			min-height: max-content;
			border: 0.1rem solid #e5e5e5;
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
					font-weight: 600;
					color: #0d1b2a;
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
			width: 70%;
			margin: 1rem auto;
			min-height: max-content;
			border: 0.1rem solid #e5e5e5;
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
					font-weight: 600;
					color: #0d1b2a;
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

				.no-content {
					min-height: 10rem;
					height: 1rem;
					display: flex;
					align-items: center;
					justify-content: space-between;
					padding: 1rem;

					p {
						font-family: $spline;
						font-size: 1.1rem;
						font-weight: 500;
						padding: 1rem;
					}
				}
			}
		}
	}
</style>
