<script lang="ts">
	import { errorMessages } from '../store';
</script>

<div class="alerts">
	{#each $errorMessages as msg}
		<div class="alert">
			<span>{msg}</span>
		</div>
	{/each}
</div>

<style lang="scss">
	.alerts {
		z-index: 1000;
		position: fixed;
		left: 50%;
		top: 2%;
		transform: translateX(-50%);
		min-width: max-content;
		display: flex;
		flex-direction: column;
		gap: 1rem;

		.alert {
			box-shadow:
				rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
				rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
			background-color: #ffcccc;
			border-radius: 0.3rem;
			transition: all 0.3s ease-in-out;
			min-height: 5rem;
			min-width: 30rem;
			display: flex;
			align-items: center;
			width: 30rem;
			height: 5rem;
			padding: 1rem;
			border: 0.2rem solid #d90429;

			span {
				font-family: $spline;
				font-size: 1.3rem;
				font-weight: 500;
				color: rgb(155 28 28/1);
				text-transform: lowercase;

				&::first-letter {
					text-transform: uppercase;
				}
			}
		}
	}
</style>
