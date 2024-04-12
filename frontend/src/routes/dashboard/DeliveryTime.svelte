<script lang="ts">
	import { time, times } from '../../store';
	import { setInboxDeliveryTime, updateErrorMessages } from '../../utils';
	import Button from '../Button.svelte';

	// let hour = '';
	// let minutes = '';
	// let am_pm = '';

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

		setInboxDeliveryTime(data);
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
				color="rgb(255, 255, 255)"
				padding={0.5}
				text="set"
				onClick={handleSetInboxDeliveryTime}
			/>
		</div>
		<div class="set-times">
			{#each $times as t}
				<div class="time">
					<input
						type="text"
						name="set_hour"
						id="set_hour"
						autocomplete="off"
						placeholder="Hour"
						value={t.hour}
						disabled
					/>
					<input
						type="text"
						name="set_minutes"
						id="set_minutes"
						autocomplete="off"
						placeholder="Minutes"
						value={t.minutes}
						disabled
					/>
					<input
						type="text"
						name="set_am_pm"
						id="set_am_pm"
						autocomplete="off"
						placeholder="am/pm"
						value={t.am_pm}
						disabled
					/>
					<Button
						height={4}
						width={10}
						borderRadius={0.3}
						color="rgb(255, 255, 255)"
						padding={0.5}
						text="delete"
						onClick={() => {}}
					/>
				</div>
			{/each}
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
		padding: 2rem;
		gap: 2rem;
		border: 0.1rem solid $black-4;
		box-shadow: rgba(0, 0, 0, 0.16) 0rem 0.1rem 0.4rem;

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

		.times {
			display: flex;
			flex-direction: column;
			gap: 6rem;

			.setter {
				display: flex;
				justify-content: space-between;
				align-items: center;
				padding: 1rem;

				input {
					outline: none;
					border: none;
					border-bottom: 0.1rem solid $blue;
					padding: 0.5rem;
					text-transform: lowercase;
				}
			}

			.set-times {
				display: flex;
				flex-direction: column;
				gap: 3rem;

				.time {
					display: flex;
					align-items: center;
					justify-content: space-between;
					padding: 1rem;

					input {
						outline: none;
						border: none;
						padding: 0.5rem;
						text-transform: lowercase;
					}
				}
			}
		}
	}
</style>
