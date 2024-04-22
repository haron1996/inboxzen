<script lang="ts">
	import { loading, time, times } from '../../store';
	import type { Time } from '../../types';
	import { setDeliveryTimes, updateErrorMessages } from '../../utils';
	import Button from '../Button.svelte';
	import CardSkeleton from './CardSkeleton.svelte';

	function isValidTimeFormat(input: string): boolean {
		var pattern = /^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$/;
		return pattern.test(input);
	}

	function handleInputBlur() {
		if ($time.delivery_time === undefined) return;

		const containsLetter = /[a-zA-Z]/.test($time.delivery_time);

		if (containsLetter) {
			$time.delivery_time = '';
			updateErrorMessages('USE 24-HOUR FORMAT ONLY');
			return;
		}

		if (!isValidTimeFormat($time.delivery_time)) {
			$time.delivery_time = '';
			updateErrorMessages('USE 24-HOUR FORMAT ONLY');
			return;
		}

		if ($times) {
			let found = false;

			$times.forEach((t) => {
				if (t.delivery_time === $time.delivery_time) {
					found = true;
					return;
				}
			});

			if (found) {
				$time = {};
				return;
			}

			$times = [$time, ...$times];
		} else {
			const ts: Time[] = [$time];
			times.set(ts);
		}

		$time = {};
	}

	function handleInputKeydown(e: KeyboardEvent) {
		if (e.code !== 'Enter' && e.code !== 'Comma') return;

		e.preventDefault();

		handleInputBlur();
	}

	function handleSetDeliveryTimes() {
		let params: string[] = [];

		if ($times) {
			$times.forEach((t) => {
				if (t.delivery_time) {
					params = [t.delivery_time, ...params];
				}
			});
		} else {
			params = [];
		}

		setDeliveryTimes(params);
	}

	function removeDeliveryTime(e: MouseEvent) {
		const t = e.currentTarget as HTMLElement;

		const div = t.closest('div.time') as HTMLElement | null;

		if (div === null) return;

		const id = div.dataset.id;

		if (id === undefined) return;

		$times = $times.filter((t) => t.id !== id);
	}
</script>

<div class="delivery-times">
	<div class="top">
		<p>DELIVERY TIMES</p>
		<span>WE WILL DELIVER EMAILS TO YOUR INBOX AT SET</span>
	</div>
	<div class="times">
		{#if $loading}
			<CardSkeleton height={10} width={80} padding={0} borderRadius={0} />
		{/if}
		{#if $times}
			{#each $times as { id, delivery_time }}
				<div class="time" data-id={id}>
					<span>{delivery_time}</span>
					<svg
						width="24px"
						height="24px"
						stroke-width="1.5"
						viewBox="0 0 24 24"
						fill="none"
						xmlns="http://www.w3.org/2000/svg"
						color="#000000"
						role="none"
						on:click={removeDeliveryTime}
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

		<input
			type="text"
			name="time"
			id="time"
			autocomplete="off"
			placeholder="24-hour format..."
			bind:value={$time.delivery_time}
			on:blur={handleInputBlur}
			on:keydown={handleInputKeydown}
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
			onClick={handleSetDeliveryTimes}
		/>
	</div>
</div>

<style lang="scss">
	.delivery-times {
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

		.times {
			min-height: max-content;
			display: flex;
			flex-flow: row wrap;
			align-content: start;
			align-items: center;
			gap: 1rem;
			padding: 2rem;

			.time {
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
