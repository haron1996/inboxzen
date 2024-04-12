<script lang="ts">
	import { URL, errorMessages } from '../../store';
	import { updateErrorMessages } from '../../utils';
	import Button from '../Button.svelte';

	let hour = '';
	let minutes = '';
	let am_pm = '';

	function handleHourInputBlur() {
		if (hour === '') return;

		if (hour.length > 1) return;

		hour = '0' + hour;
	}

	function handleMinutesInputBlur() {
		if (minutes === '') return;

		if (minutes.length > 1) return;

		minutes = '0' + minutes;
	}

	async function setDeliveryTime() {
		const url = `${$URL}/private/setdeliverytime`;

		if (hour === '' || minutes === '' || am_pm === '') {
			updateErrorMessages('Set a valid time before submitting');
			return;
		}

		const data = {
			hour: hour,
			minutes: minutes,
			am_pm: am_pm
		};

		console.log(data);

		await fetch(url, {
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			},
			redirect: 'follow',
			referrerPolicy: 'no-referrer',
			body: JSON.stringify(data)
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error('Network response was not ok');
				}

				return response.json();
			})
			.then((data) => {
				console.log(data);
			})
			.catch((error) => {
				// console.error('Error:', error);
				// if (error.message === 'Failed to fetch') {
				// 	console.error('Connection error: Unable to reach the server.');
				// 	updateErrorMessages('Unable to reach the server.');
				// }
				console.log(error);
				errorMessages.update(error.message);
			});
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
				bind:value={hour}
				on:blur={handleHourInputBlur}
			/>
			<input
				type="text"
				name="hour"
				id="hour"
				autocomplete="off"
				placeholder="Minutes"
				bind:value={minutes}
				on:blur={handleMinutesInputBlur}
			/>
			<input
				type="text"
				name="hour"
				id="hour"
				autocomplete="off"
				placeholder="am/pm"
				bind:value={am_pm}
			/>
			<Button
				height={4}
				width={10}
				backgroundColor="#00a6fb"
				borderRadius={0.3}
				color="rgb(255, 255, 255)"
				padding={0.5}
				text="set"
				onClick={setDeliveryTime}
			/>
		</div>
		<div class="set-times">
			<div class="time">
				<input
					type="text"
					name="hour"
					id="hour"
					autocomplete="off"
					placeholder="Hour"
					value="06"
					disabled
				/>
				<input
					type="text"
					name="hour"
					id="hour"
					autocomplete="off"
					placeholder="Minutes"
					value="00"
					disabled
				/>
				<input
					type="text"
					name="hour"
					id="hour"
					autocomplete="off"
					placeholder="am/pm"
					value="am"
					disabled
				/>
				<Button
					height={4}
					width={10}
					backgroundColor="#ff4d6d"
					borderRadius={0.3}
					color="rgb(255, 255, 255)"
					padding={0.5}
					text="delete"
					onClick={() => {}}
				/>
			</div>
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
