<script lang="ts">
	import { URL, errors } from '../store';
	import Button from './Button.svelte';
	import Header from './Header.svelte';

	async function handleButtonClick() {
		const url = `${$URL}/private/checkuserloginstatus`;

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
			errors.update((errs) => [result.message, ...errs]);
			return;
		}

		location.replace(result);
	}
</script>

<svelte:head>
	<title>Inboxzen | Way To Inbox Zero</title>
	<meta name="description" content="Reclaim control of your inbox with InboxZen" />
</svelte:head>

<Header />

<section>
	<div class="offer">
		<h1>Empower Your Inbox: Control Senders and Delivery Times with Our Plugin</h1>
		<p>
			Take charge of your inbox with our plugin, allowing you to manage allowed senders and schedule
			message delivery times effortlessly.
		</p>
	</div>
	<Button
		height={4}
		width={40}
		backgroundColor="#525FE1"
		borderRadius={0.3}
		color="rgb(255, 255, 255)"
		padding={0.5}
		text="Sign in with Google"
		onClick={handleButtonClick}
	/>
</section>

<style lang="scss">
	section {
		width: 100dvw;
		min-height: 90dvh;
		height: max-content;
		overflow-y: auto;
		background-color: $white;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		gap: 4rem;
		padding: 2rem;

		.offer {
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			text-align: center;
			gap: 2rem;

			h1 {
				font-size: 3rem;
				text-transform: capitalize;
				font-weight: 600;
				color: $black;
				max-width: 50%;
				font-family: $spline;
			}

			p {
				font-family: 'Spline Sans', sans-serif;
				font-size: 1.8rem;
				line-height: 1.6;
				color: $light-black;
				max-width: 50%;
			}
		}
	}
</style>
