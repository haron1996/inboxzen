<script lang="ts">
	import { URL, session } from '../../store';
	import { switchAccount, updateErrorMessages } from '../../utils';
	import ProfileCard from './ProfileCard.svelte';

	function loadProfileSettings() {
		location.href = '/dashboard/profile';
	}

	function loadAccountSettings() {
		location.href = '/dashboard/settings';
	}

	function goToPreAuth() {
		location.href = '/preauth';
	}

	async function logout() {
		const url = `${$URL}/private/logout`;

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
						return;
					}, 3000);
				default:
					updateErrorMessages(message);
					return;
			}
		}

		location.href = '/';
	}
</script>

<menu id="menu">
	<div class="logout" role="none" on:click={logout}>
		<span>SIGN OUT</span>
	</div>
	<div class="add-account" role="none" on:click={goToPreAuth}>
		<span>ADD ACCOUNT</span>
	</div>
	<div class="account" role="none" on:click={loadProfileSettings}>
		<span>PROFILE SETTINGS</span>
	</div>
	<div class="account" role="none" on:click={loadAccountSettings}>
		<span>ACCOUNT SETTINGS</span>
	</div>
	{#if $session && $session.emails}
		{#each $session.emails as { account_name, email_address, profile_picture }}
			<div class="account">
				<ProfileCard
					name={account_name}
					email={email_address}
					picture={profile_picture}
					SVGDisplay="none"
					width="100%"
					onClick={(e) => {
						switchAccount(e);
					}}
				/>
			</div>
		{/each}
	{/if}
</menu>

<style lang="scss">
	menu {
		position: absolute;
		background-color: $white;
		display: none;
		min-height: max-content;
		overflow-y: auto;
		cursor: pointer;
		border-radius: 0.3rem;
		transition: all 0.3s linear;
		border: 0.2rem solid #0d1b2a;
		max-height: 90dvh;

		div {
			width: 100%;
			min-height: 4rem;
			display: flex;
			align-items: center;
			padding: 1rem;
			gap: 1rem;
			transition: all 0.3s linear;

			span {
				font-family: $spline;
				text-transform: uppercase;
				font-weight: 600;
				color: #0d1b2a;
			}

			&:hover {
				background-color: #ffd166;
			}
		}

		div:not(:last-child) {
			border-bottom: 0.1rem solid #dad7cd;
		}
	}
</style>
