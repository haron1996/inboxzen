<script lang="ts">
	import { email, emails, loading } from '../../store';
	import { updateVipEmailAddresses } from '../../utils';
	import Button from '../Button.svelte';
	import CardSkeleton from './CardSkeleton.svelte';

	function handleEmailInputBlur() {
		if ($email.vip_email_address === undefined) return;

		$email.vip_email_address = $email.vip_email_address.trim();

		$email.vip_email_address = $email.vip_email_address.toLowerCase();

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
		width: 70%;
		margin: 1rem auto;
		min-height: max-content;
		background-color: $white;
		display: flex;
		flex-direction: column;
		border-radius: 0.3rem;
		border: 0.1rem solid #e5e5e5;

		.top {
			display: flex;
			flex-direction: column;
			justify-content: center;
			gap: 1rem;
			border-bottom: 0.1rem solid #e0e1dd;
			padding: 2rem;

			p {
				font-family: $spline;
				font-size: 1.1rem;
				text-transform: uppercase;
				font-weight: 600;
				color: #0d1b2a;
			}

			span {
				font-family: $spline;
				font-weight: 600;
				text-transform: uppercase;
				color: #757575;
			}
		}

		.emails {
			min-height: max-content;
			display: flex;
			flex-flow: row wrap;
			align-content: start;
			align-items: center;
			gap: 1rem;
			padding: 2rem;

			.email {
				min-height: max-content;
				display: flex;
				align-items: center;
				padding: 0.5rem 0.5rem 0.5rem 0.7rem;
				width: max-content;
				border-radius: 0.5rem;
				gap: 0.3rem;
				background-color: #fbfbff;
				border: 0.1rem solid #e5e5e5;

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
				text-transform: lowercase;
			}
		}

		.bottom {
			display: flex;
			align-items: center;
			justify-content: flex-end;
			padding: 2rem;
		}
	}
</style>
