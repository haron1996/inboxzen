<script lang="ts">
	import { URL, session } from '../../store';
	import { hideMenu, showMenu, updateErrorMessages } from '../../utils';
	import ProfileCard from './ProfileCard.svelte';

	function goToAccount() {
		location.href = '/dashboard/profile';
	}

	function goToSettings() {
		location.href = '/dashboard/settings';
	}

	function goToPreAuth() {
		location.href = '/preauth';
	}

	async function switchAccount(e: MouseEvent) {
		const target = e.currentTarget as HTMLDivElement;

		const email = target.children[1].children[1].textContent;

		if (email === null) return;

		const url = `${$URL}/private/switchaccount/${email}`;

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

		session.set(result);

		hideMenu();
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
	<div class="support">
		<span>CONTACT SUPPORT</span>
	</div>
	<div class="settings" role="none" on:click={goToSettings}>
		<span>ACCOUNT SETTINGS</span>
	</div>
	<div class="account" role="none" on:click={goToAccount}>
		<span>PROFILE SETTINGS</span>
	</div>

	{#if $session && $session.emails}
		{#each $session.emails as { account_name, email_address, profile_picture }}
			<div class="account">
				<ProfileCard
					name={account_name}
					email={email_address}
					picture={profile_picture}
					SVGDisplay="none"
					width="30rem"
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
		box-shadow:
			rgba(255, 255, 255, 0.2) 0rem 0rem 0rem 0.1rem inset,
			rgba(0, 0, 0, 0.9) 0rem 0rem 0rem 0.1rem;
		display: none;
		flex-direction: column;
		min-height: max-content;
		overflow-y: auto;
		cursor: pointer;
		border-radius: 0.3rem;

		div {
			width: 100%;
			min-height: 4rem;
			display: flex;
			align-items: center;
			padding: 1rem;
			gap: 1rem;

			span {
				font-family: $spline;
				text-transform: uppercase;
				font-weight: 500;
			}

			&:hover {
				background-color: $gray;
			}
		}
	}
</style>
