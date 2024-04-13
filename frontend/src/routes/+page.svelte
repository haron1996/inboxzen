<script lang="ts">
	import { checkUserLoginStatus } from '../utils';
	import Button from './Button.svelte';
	import Header from './Header.svelte';
</script>

<svelte:head>
	<title>ZENN.EMAIL: BECOME A GMAIL PRO</title>
	<meta name="description" content="Reclaim control of your inbox with InboxZen" />
</svelte:head>

<Header />

<section>
	<div class="offer">
		<p>Take Control of Your Gmail Inbox</p>
		<span>
			Goodbye inbox clutter. Say hello to inbox zero. Zenn helps you take control of who makes it to
			your inbox and at what time. This way, you reclaim the time wasted on your inbox, increase
			your productivity, and make Gmail fun again.
		</span>
		<Button
			height={4.5}
			width={30}
			borderRadius={0.6}
			color="#0d1b2a"
			padding={0.5}
			text="CONTINUE WITH GOOGLE"
			fontWeight={500}
			onClick={checkUserLoginStatus}
		/>
	</div>
</section>

<style lang="scss">
	section {
		width: 100dvw;
		min-height: 90dvh;
		height: max-content;
		overflow-y: auto;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 2rem;

		.offer {
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			text-align: center;
			gap: 3rem;
			width: 50rem;
			padding: 1rem;
			color: #0d1b2a;
			font-family: 'Jost', sans-serif;
			font-optical-sizing: auto;

			p {
				font-size: 2rem;
				font-weight: 600;
				text-transform: uppercase;
			}

			span {
				font-size: 1.5rem;
				color: $light-black;
			}
		}
	}
</style>
