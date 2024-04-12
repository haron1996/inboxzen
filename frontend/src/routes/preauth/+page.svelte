<script lang="ts">
	import { getAuthURL } from '../../utils';
	import Button from '../Button.svelte';
</script>

<svelte:head>
	<title>Pre-auth | Inbox Check</title>
</svelte:head>

<section>
	<div class="card">
		<div class="title">
			<p>InboxZen will request the following permissions.</p>
		</div>
		<ul>
			<li>
				<p>🚀 see and edit your email labels.</p>
				<span
					>We will use this permission to organize your emails by creating a folder in your Gmail
					account specifically for snoozing incoming emails.</span
				>
			</li>
			<li>
				<p>🚀 read, compose, and send emails from your Gmail account.</p>
				<span>We will use this permission to move snoozed emails to your inbox folder.</span>
			</li>
			<li>
				<p>🚀 see, edit, create, or change your email settings and filters in Gmail</p>
				<span
					>We will use this permission to manage incoming emails by redirecting them to a snoozed
					folder.</span
				>
			</li>
		</ul>
		<Button
			height={4}
			width={45}
			backgroundColor="#00a6fb"
			borderRadius={0.6}
			color="rgb(255, 255, 255)"
			padding={0.5}
			text="agree and continue"
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
		background-color: #4e54c8;
		background-image: linear-gradient(to right top, #4e54c8, #8f94fb, #b4b8fd, #cad3ff, #e5f0ff);

		.card {
			width: 50rem;
			display: flex;
			flex-direction: column;
			gap: 4rem;
			line-height: 1.5;
			padding: 1rem;

			.title {
				display: flex;
				align-items: center;
				justify-content: flex-start;

				p {
					font-family: $spline;
					font-size: 1.3rem;
					color: $black;
					text-transform: uppercase;
					font-weight: 500;
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
						font-family: $spline;
						font-size: 1.3rem;
						font-weight: 600;
						color: $black;
						text-transform: none;
					}

					span {
						font-family: $spline;
						font-size: 1.4rem;
						color: $black-1;
					}
				}
			}
		}
	}
</style>
