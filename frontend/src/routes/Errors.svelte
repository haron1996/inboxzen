<script lang="ts">
	import { errors } from '../store';
</script>

<div class="errors">
	{#each $errors as error}
		<div class="error">
			<span>{error}</span>
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
			min-height: 4rem;
			height: max-content;
			padding: 1rem;
			display: flex;
			align-items: center;
			justify-content: center;
			border-radius: 0.3rem;
			position: relative;
			gap: 1rem;
			box-shadow: rgba(0, 0, 0, 0.16) 0rem 0.1rem 0.4rem;
			border-radius: 0.5rem;

			span {
				font-size: 1.3rem;
				font-family: $spline;
				text-transform: capitalize;
				color: $white;
				font-weight: 500;
			}
		}
	}
</style>
