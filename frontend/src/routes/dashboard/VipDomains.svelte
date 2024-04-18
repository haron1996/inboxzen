<script lang="ts">
	import { domain, domains, loading } from '../../store';
	import Button from '../Button.svelte';
	import { updateVipDomains } from '../../utils';
	import CardSkeleton from './CardSkeleton.svelte';

	function handleDomainInputBlur() {
		if ($domain.domain_name === undefined) return;

		$domain.domain_name = $domain.domain_name.trim();

		$domain.domain_name = $domain.domain_name.toLowerCase();

		let found = false;

		$domains.forEach((d) => {
			if (d.domain_name === $domain.domain_name) {
				found = true;
				return;
			}
		});

		if (found) {
			$domain = {};
			return;
		}

		$domains = [$domain, ...$domains];

		$domain = {};
	}

	function handleDomainInputKeydown(e: KeyboardEvent) {
		if (e.code !== 'Enter' && e.code !== 'Comma') return;

		e.preventDefault();

		handleDomainInputBlur();
	}

	function removeVIPDomain(e: MouseEvent) {
		const t = e.currentTarget as HTMLElement;

		const div = t.closest('.domain') as HTMLDivElement | null;

		if (div === null) return;

		const domainName = div.innerText;

		$domains = $domains.filter((d) => d.domain_name !== domainName);
	}
</script>

<div class="vip-domains">
	<div class="top">
		<p>VIP DOMAINS</p>
		<span>Emails from any of these domains will be delivered instantly</span>
	</div>
	<div class="domains">
		{#if $loading}
			<CardSkeleton height={10} width={80} padding={0} borderRadius={0} />
		{/if}
		{#each $domains as { domain_name }}
			<div class="domain">
				<span>{domain_name}</span>
				<svg
					width="24px"
					height="24px"
					stroke-width="1.5"
					viewBox="0 0 24 24"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
					color="#000000"
					role="none"
					on:click={removeVIPDomain}
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
			type="text"
			name="domain"
			id="domain"
			autocomplete="off"
			placeholder="add domain..."
			bind:value={$domain.domain_name}
			on:blur={handleDomainInputBlur}
			on:keydown={handleDomainInputKeydown}
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
				updateVipDomains($domains);
			}}
		/>
	</div>
</div>

<style lang="scss">
	.vip-domains {
		width: 80rem;
		min-height: max-content;
		background-color: $white;
		display: flex;
		flex-direction: column;
		border-radius: 0.3rem;
		border: 0.1rem solid #0d1b2a;

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

		.domains {
			min-height: max-content;
			display: flex;
			flex-flow: row wrap;
			align-content: start;
			align-items: center;
			gap: 1rem;
			padding: 2rem;

			.domain {
				min-height: max-content;
				display: flex;
				align-items: center;
				padding: 0.5rem 0.5rem 0.5rem 0.7rem;
				width: max-content;
				border-radius: 0.5rem;
				gap: 0.3rem;
				background-color: #ffd166;
				border: 0.2rem solid #fb8500;

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
