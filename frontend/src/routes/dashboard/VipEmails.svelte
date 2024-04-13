<script lang="ts">
	import { email, emails, loading } from '../../store';
	import { updateVipEmailAddresses } from '../../utils';
	import Button from '../Button.svelte';
	import CardSkeleton from './CardSkeleton.svelte';

	function handleEmailInputBlur() {
		if ($email.vip_email_address === undefined) return;

		let found = false;

		$emails.forEach((e) => {
			if (e.vip_email_address === $email.vip_email_address) {
				found = true;
				return;
			}
		});

		if (found) {
			$email = {};
			return;
		}

		$emails = [$email, ...$emails];

		$email = {};
	}

	function handleEmailInputKeydown(e: KeyboardEvent) {
		if (e.code !== 'Enter' && e.code !== 'Comma') return;

		e.preventDefault();

		handleEmailInputBlur();
	}

	function removeVIPEmail(e: MouseEvent) {
		const t = e.currentTarget as HTMLElement;

		const div = t.closest('.email') as HTMLDivElement | null;

		if (div === null) return;

		const emailAddress = div.innerText;

		$emails = $emails.filter((e) => e.vip_email_address !== emailAddress);
	}
</script>

<div class="vip-emails">
	<div class="top">
		<p>VIP EMAILS</p>
		<span>Emails from any of these email addresses will be delivered instantly</span>
	</div>
	<div class="emails">
		{#if $loading}
			<CardSkeleton height={10} width={80} padding={0} borderRadius={0} />
		{/if}
		{#each $emails as { vip_email_address }}
			<div class="email">
				<span>{vip_email_address}</span>
				<svg
					width="24px"
					height="24px"
					stroke-width="1.5"
					viewBox="0 0 24 24"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
					color="#000000"
					role="none"
					on:click={removeVIPEmail}
					><path
						d="M6.75827 17.2426L12.0009 12M17.2435 6.75736L12.0009 12M12.0009 12L6.75827 6.75736M12.0009 12L17.2435 17.2426"
						stroke="#000000"
						stroke-width="1.5"
						stroke-linecap="round"
						stroke-linejoin="round"
					></path></svg
				>
			</div>
		{/each}
		<input
			type="email"
			name="email"
			id="email"
			autocomplete="off"
			placeholder="add email..."
			bind:value={$email.vip_email_address}
			on:blur={handleEmailInputBlur}
			on:keydown={handleEmailInputKeydown}
		/>
	</div>
	<div class="bottom">
		<Button
			height={4}
			width={10}
			borderRadius={0.3}
			color="#0d1b2a"
			padding={0.5}
			text="update"
			onClick={() => {
				updateVipEmailAddresses($emails);
			}}
		/>
	</div>
</div>

<style lang="scss">
	.vip-emails {
		width: 80rem;
		min-height: max-content;
		background-color: $white;
		display: flex;
		flex-direction: column;
		border-radius: 0.3rem;
		padding: 2rem;
		gap: 2rem;
		border: 0.1rem solid #0d1b2a;

		.top {
			display: flex;
			flex-direction: column;
			justify-content: center;
			gap: 0.5rem;

			p {
				font-family: $spline;
				font-size: 1.1rem;
				text-transform: uppercase;
				font-weight: 500;
			}

			span {
				font-size: 1.3rem;
				font-family: $spline;
				font-weight: 500;
				color: $black-1;
			}
		}

		.emails {
			min-height: max-content;
			display: flex;
			flex-flow: row wrap;
			align-content: start;
			align-items: center;
			gap: 1rem;

			.email {
				background-color: #f6f5f5;
				min-height: max-content;
				display: flex;
				align-items: center;
				padding: 0.5rem 0.5rem 0.5rem 0.7rem;
				width: max-content;
				border-radius: 0.5rem;
				gap: 0.3rem;
				box-shadow:
					rgba(0, 0, 0, 0.05) 0rem 0rem 0rem 0.1rem,
					rgb(209, 213, 219) 0rem 0rem 0rem 0.1rem inset;

				span {
					font-size: 1.3rem;
					font-family: $spline;
					font-weight: 500;
					color: $black-1;
				}

				svg {
					height: 1.5rem;
					width: 1.5rem;
					cursor: pointer;
				}
			}

			input {
				border: none;
				outline: none;
				font-family: $spline;
				padding: 0.5rem;
			}
		}

		.bottom {
			display: flex;
			align-items: center;
			justify-content: flex-end;
			padding: 1rem;
		}
	}
</style>
