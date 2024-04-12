<script lang="ts">
	import { checkUserLoginStatus } from '../utils';
	import Button from './Button.svelte';
	import Features from './Features.svelte';
	import Header from './Header.svelte';
</script>

<svelte:head>
	<title>ZENN.EMAIL: BECOME A GMAIL PRO</title>
	<meta name="description" content="Reclaim control of your inbox with InboxZen" />
</svelte:head>

<section>
	<div class="offer">
		<h1>Become a Gmail Pro</h1>
		<p>
			Goodbye inbox clutter. Say hello to inbox zero. Zenn helps you take control of who makes it to
			your inbox and at what time. This way, you reclaim the time wasted on your inbox, increase
			your productivity, and make Gmail fun again,
		</p>
		<Button
			height={4.5}
			width={30}
			borderRadius={0.6}
			color="rgb(255, 255, 255)"
			padding={0.5}
			text="Get Started Now"
			fontWeight={900}
			onClick={checkUserLoginStatus}
		/>
	</div>
</section>

<Features />

<style lang="scss">
	section {
		width: 100dvw;
		min-height: 100dvh;
		height: max-content;
		overflow-y: auto;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 2rem;
		background-color: #4e54c8;
		background-image: linear-gradient(to right top, #4e54c8, #8f94fb, #b4b8fd, #cad3ff, #e5f0ff);

		.offer {
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			text-align: center;
			gap: 3rem;
			width: 50rem;
			padding: 1rem;

			h1 {
				font-size: 4rem;
				font-weight: 900;
				color: black;
				font-family: 'Arial MT Black', sans-serif;
			}

			p {
				font-size: 2rem;
				line-height: 1.6;
				color: $light-black;
				font-family: 'Product Sans', sans-serif;
			}
		}
	}
</style>
