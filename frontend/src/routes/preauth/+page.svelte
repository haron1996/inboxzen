<script lang="ts">
	import { getAuthURL } from '../../utils';
	import Button from '../Button.svelte';
</script>

<svelte:head>
	<title>Pre-auth | Gmail Inbox Zero</title>
</svelte:head>

<section>
	<div class="card">
		<div class="title">
			<p>Zenn will request the following permissions.</p>
		</div>
		<ul>
			<li>
				<p>1. See and edit your email labels.</p>
				<span
					>We will use this permission to organize your emails by creating a folder in your Gmail
					account specifically for snoozing incoming emails.</span
				>
			</li>
			<li>
				<p>2. Read, compose, and send emails from your Gmail account.</p>
				<span>We will use this permission to move snoozed emails to your inbox folder.</span>
			</li>
			<li>
				<p>3. See, edit, create, or change your email settings and filters in Gmail</p>
				<span
					>We will use this permission to manage incoming emails by redirecting them to a snoozed
					folder.</span
				>
			</li>
		</ul>
		<Button
			height={4}
			width={45}
			borderRadius={0.3}
			color="#0d1b2a"
			padding={0.5}
			text="continue"
			onClick={getAuthURL}
		/>
	</div>
</section>

<style lang="scss">
	section {
		width: 100dvw;
		min-height: 100dvh;
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 2rem;

		.card {
			width: 50rem;
			display: flex;
			flex-direction: column;
			gap: 2rem;
			line-height: 1.5;
			padding: 1rem;
			color: #0d1b2a;
			font-family: 'Jost', sans-serif;
			font-optical-sizing: auto;

			.title {
				display: flex;
				align-items: center;
				justify-content: flex-start;

				p {
					font-size: 1.3rem;
					font-weight: 500;
					text-transform: uppercase;
				}
			}

			ul {
				display: flex;
				flex-direction: column;
				gap: 2rem;

				li {
					list-style: none;
					display: flex;
					flex-direction: column;
					gap: 0.5rem;

					p {
						font-size: 1.2rem;
						font-weight: 500;
						text-transform: none;
						text-transform: uppercase;
					}

					span {
						font-size: 1.3rem;
					}
				}
			}
		}
	}
</style>
