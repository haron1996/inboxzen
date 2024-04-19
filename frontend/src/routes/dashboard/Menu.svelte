<script lang="ts">
	import { URL, session } from '../../store';
	import { switchAccount, updateErrorMessages } from '../../utils';
	import ProfileCard from './ProfileCard.svelte';

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

<menu class="animate__animated" id="menu">
	<a href="/dashboard">
		<span>DASHBOARD</span>
	</a>
	<a href="/preauth">
		<span>ADD ACCOUNT</span>
	</a>
	<a href="/dashboard/profile">
		<span>PROFILE SETTINGS</span>
	</a>
	<a href="/dashboard/screener">
		<span>INBOX SCREENER</span>
	</a>
	{#if $session && $session.emails}
		{#each $session.emails as { account_name, email_address, profile_picture }}
			<a href="/" on:click|preventDefault={() => {}}>
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
			</a>
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
		max-height: 90dvh;
		box-shadow: rgba(0, 0, 0, 0.09) 0px 3px 12px;
		background-color: #00a8e8;
		border: 0.2rem solid #00a8e8;
		z-index: 2000;

		a {
			width: 100%;
			min-height: 5rem;
			display: flex;
			align-items: center;
			padding: 1rem;
			gap: 1rem;
			text-decoration: none;

			span {
				font-family: $spline;
				text-transform: uppercase;
				font-weight: 600;
				color: #0d1b2a;
			}

			&:hover {
				background-color: #1e91d6;
			}
		}

		a:not(:last-child) {
			border-bottom: 0.1rem solid #dad7cd;
		}
	}
</style>
