<script lang="ts">
	import { URL, errors } from '../../store';

	async function getAuthURL() {
		const url = `${$URL}/account/getauthurl`;

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
	<title>Pre-auth | Inboxzen</title>
</svelte:head>

<section>
	<div class="card">
		<div class="title">
			<p>InboxZen will request the following permissions. 😊</p>
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
		<button on:click|preventDefault={getAuthURL}>
			<span>agree and continue 💪</span>
		</button>
	</div>
</section>

<style lang="scss">
	section {
		width: 100dvw;
		min-height: 90dvh;
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 2rem;

		.card {
			width: 50rem;
			display: flex;
			flex-direction: column;
			gap: 4rem;
			line-height: 1.5;
			padding: 0.5rem;

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
			width: 100%;
			display: flex;
			align-items: center;
			justify-content: center;
			gap: 1rem;

			span {
				font-size: 1.3rem;
				font-family: 'Spline Sans', sans-serif;
				color: $white;
				font-weight: 500;
				text-transform: uppercase;
				letter-spacing: 0.1rem;
			}
		}
	}
</style>
