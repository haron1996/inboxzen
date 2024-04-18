<script lang="ts">
	import { loading, time, times } from '../../store';
	import type { Time } from '../../types';
	import { deleteInboxDeliveryTime, setInboxDeliveryTime, updateErrorMessages } from '../../utils';
	import Button from '../Button.svelte';
	import CardSkeleton from './CardSkeleton.svelte';

	function handleHourInputBlur() {
		if ($time.hour === '' || $time.hour === undefined) return;

		if ($time.hour && $time.hour.length > 1) return;

		$time.hour = '0' + $time.hour;
	}

	function handleMinutesInputBlur() {
		if ($time.minutes === '' || $time.minutes === undefined) return;

		if ($time.minutes && $time.minutes.length > 1) return;

		$time.minutes = '0' + $time.minutes;
	}

	function handleSetInboxDeliveryTime() {
		if (
			$time.hour === '' ||
			$time.hour === undefined ||
			$time.minutes === '' ||
			$time.minutes === undefined ||
			$time.am_pm === '' ||
			$time.am_pm === undefined
		) {
			updateErrorMessages('Set a valid time before submitting');
			return;
		}

		const data = {
			hour: $time.hour,
			minutes: $time.minutes,
			am_pm: $time.am_pm
		};

		setInboxDeliveryTime(data, $times);
	}

	function handleDeleteDeliveryTime(e: MouseEvent) {
		const t = e.currentTarget as SVGAElement | null;

		if (t === null) return;

		const div = t.closest('.domain') as HTMLDivElement | null;

		if (div === null) return;

		const id = div.dataset.id;

		const data = {
			id: id
		};

		deleteInboxDeliveryTime(data, $times);
	}
</script>

<div class="delivery-times">
	<div class="top">
		<p>DELIVERY TIMES</p>
		<span>Set what time your emails will be delivered to your inbox</span>
	</div>
	<div class="times">
		<div class="setter">
			<input
				type="text"
				name="hour"
				id="hour"
				autocomplete="off"
				placeholder="Hour"
				bind:value={$time.hour}
				on:blur={handleHourInputBlur}
			/>
			<input
				type="text"
				name="minutes"
				id="minutes"
				autocomplete="off"
				placeholder="Minutes"
				bind:value={$time.minutes}
				on:blur={handleMinutesInputBlur}
			/>
			<input
				type="text"
				name="am_pm"
				id="am_pm"
				autocomplete="off"
				placeholder="am/pm"
				bind:value={$time.am_pm}
			/>
			<Button
				height={4}
				width={10}
				borderRadius={0.3}
				color="#0d1b2a"
				padding={0.5}
				text="update"
				onClick={handleSetInboxDeliveryTime}
			/>
		</div>
		<div class="set-times">
			{#if $loading}
				<CardSkeleton height={10} width={80} padding={0} borderRadius={0} />
			{/if}
			{#if $times}
				{#each $times as t}
					<div class="domain" data-id={t.id}>
						<span>{t.hour}:{t.minutes} {t.am_pm}</span>
						<svg
							width="24px"
							height="24px"
							stroke-width="1.5"
							viewBox="0 0 24 24"
							fill="none"
							xmlns="http://www.w3.org/2000/svg"
							color="#000000"
							role="none"
							on:click={(e) => {
								handleDeleteDeliveryTime(e);
							}}
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
			{/if}
		</div>
	</div>
</div>

<style lang="scss">
	.delivery-times {
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

		.times {
			display: flex;
			flex-direction: column;
			gap: 1rem;
			padding: 2rem;

			.setter {
				display: flex;
				justify-content: space-between;
				align-items: center;

				input {
					outline: none;
					border: none;
					border-bottom: 0.1rem solid #0d1b2a;
					padding: 0.5rem;
					text-transform: lowercase;
					font-family: $spline;
					font-size: 1.2rem;
					font-weight: 500;
					letter-spacing: 0.1rem;
					text-transform: uppercase;
				}
			}

			.set-times {
				min-height: max-content;
				display: flex;
				flex-flow: row wrap;
				align-content: start;
				align-items: center;
				gap: 1rem;

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
						text-transform: uppercase;
					}

					svg {
						height: 1.5rem;
						width: 1.5rem;
						cursor: pointer;
					}
				}
			}
		}
	}
</style>
