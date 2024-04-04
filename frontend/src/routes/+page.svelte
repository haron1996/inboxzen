<script lang="ts">
	import Header from './Header.svelte';
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
	<button>
		<span>Sign up with Google</span>
	</button>
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
				font-size: 4rem;
				text-transform: capitalize;
				font-weight: 500;
				color: $black;
				max-width: 50%;
				font-family: 'Scholarly Ambition', sans-serif;
			}

			p {
				font-family: 'Spline Sans', sans-serif;
				font-size: 1.8rem;
				line-height: 1.6;
				color: $light-black;
				max-width: 50%;
			}
		}

		button {
			border: none;
			width: max-content;
			padding: 1.5rem;
			border: 0.2rem solid #008dda;
			border-radius: 0.3rem;
			cursor: pointer;
			background-color: #0079ff;
			box-shadow:
				rgba(0, 0, 0, 0.16) 0px 1px 4px,
				#0079ff 0px 0px 0px 2px;
			width: 30rem;
			display: flex;
			align-items: center;
			justify-content: center;
			gap: 1rem;

			span {
				font-size: 1.3rem;
				font-family: $spline;
				color: $white;
				font-weight: 500;
				text-transform: uppercase;
				letter-spacing: 0.1rem;
			}
		}
	}
</style>
