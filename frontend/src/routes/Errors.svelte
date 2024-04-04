<script lang="ts">
	import { errors } from '../store';

	function removeError(e: MouseEvent) {
		const target = e.currentTarget as HTMLElement;

		const div = target.closest('.error') as HTMLDivElement | null;

		if (div === null) return;

		const span = div.firstElementChild as HTMLSpanElement | null;

		if (span === null) return;

		const text = span.innerText;

		$errors = $errors.filter((value) => value === text);
	}
</script>

<div class="errors">
	{#each $errors as error}
		<div class="error">
			<span>{error}</span>
			<svg
				width="24px"
				height="24px"
				stroke-width="1.5"
				viewBox="0 0 24 24"
				fill="none"
				xmlns="http://www.w3.org/2000/svg"
				color="#000000"
				role="none"
				on:click={removeError}
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
</div>

<style lang="scss">
	.errors {
		position: fixed;
		left: 50%;
		top: 1rem;
		transform: translateX(-50%);
		z-index: 100;
		background-color: lightcyan;
		width: max-content;
		height: max-content;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		gap: 1rem;

		.error {
			background-color: $red;
			width: 40rem;
			min-height: 5rem;
			height: max-content;
			padding: 2rem;
			display: flex;
			align-items: center;
			justify-content: center;
			border-radius: 0.3rem;
			position: relative;
			gap: 1rem;
			box-shadow:
				rgba(0, 0, 0, 0.25) 0rem 0.0625em 0.0625em,
				rgba(0, 0, 0, 0.25) 0rem 0.125em 0.5em,
				rgba(255, 255, 255, 0.1) 0rem 0rem 0rem 0.1rem inset;

			span {
				font-size: 1.5rem;
				font-family: $spline;
				text-transform: capitalize;
				color: $white;
				font-weight: 500;
			}

			svg {
				position: absolute;
				top: 0.1rem;
				right: 0.1rem;
				cursor: pointer;

				path {
					stroke: $white;
				}
			}
		}
	}
</style>
