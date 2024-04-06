<script lang="ts">
	import { URL } from '../store';
	import Button from './Button.svelte';

	async function handleButtonClick() {
		const url = `${$URL}/private/checkloginstatus`;

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

		if (!response.ok) {
			location.href = '/preauth';
			return;
		}

		location.href = '/dashboard';
	}
</script>

<header>
	<a href="/" class="logo">
		<span>InboxZen</span>
	</a>
	<Button
		height={3.5}
		width={10}
		backgroundColor="#525FE1"
		borderRadius={0.3}
		color="rgb(255, 255, 255)"
		padding={0.5}
		text="login"
		onClick={handleButtonClick}
	/>
</header>

<style lang="scss">
	header {
		width: 100dvw;
		min-height: 10vh;
		background-color: $white;
		height: max-content;
		padding: 2rem;
		display: flex;
		align-items: center;
		justify-content: space-between;

		a.logo {
			display: flex;
			align-items: center;
			justify-content: center;
			width: max-content;
			gap: 0.5rem;
			text-decoration: none;

			span {
				font-size: 1.5rem;
				font-weight: 900;
				text-transform: uppercase;
				color: $black;
				letter-spacing: 0.1rem;
				letter-spacing: 0.1rem;
				font-family: $spline;
			}
		}
	}
</style>
