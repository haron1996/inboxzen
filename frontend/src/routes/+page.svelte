<script lang="ts">
	import { URL } from '../store';
	import Button from './Button.svelte';
	import Header from './Header.svelte';

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

<svelte:head>
	<title>Inbox Check | Way To Inbox Zero</title>
	<meta name="description" content="Reclaim control of your inbox with InboxZen" />
</svelte:head>

<Header />

<section>
	<div class="offer">
		<h1>Manage Senders and Delivery Times with Our Gmail Plugin.</h1>
		<p>
			Effortlessly take charge of your inbox: control approved senders and schedule message delivery
			with ease.
		</p>
	</div>
	<Button
		height={5}
		width={40}
		backgroundColor="#00a6fb"
		borderRadius={0.6}
		color="rgb(255, 255, 255)"
		padding={0.5}
		text="Get Started Now"
		onClick={handleButtonClick}
	/>
</section>

<style lang="scss">
	section {
		width: 100dvw;
		min-height: 70dvh;
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
			gap: 1rem;
			width: 40%;

			h1 {
				font-size: 2.4rem;
				font-weight: 900;
				text-transform: capitalize;
				color: $black;
				font-family: $spline;
			}

			p {
				font-family: 'Spline Sans', sans-serif;
				font-size: 1.8rem;
				line-height: 1.5;
				color: $light-black;
			}
		}
	}
</style>
