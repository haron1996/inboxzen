<script lang="ts">
	export let width = 0;
	export let onClick: () => void;
</script>

<button style="width: {width}rem;" on:click={onClick}>
	<span>BACK TO DASHBOARD</span>
</button>

<style lang="scss">
	button {
		background-color: $black-2;
		display: flex;
		align-items: center;
		justify-content: center;
		border: 0.1rem solid $black-1;
		box-shadow:
			rgba(67, 71, 85, 0.27) 0rem 0rem 0.25em,
			rgba(90, 125, 188, 0.05) 0rem 0.25em 1em;
		min-width: 10rem;
		gap: 1rem;
		height: 4rem;
		border-radius: 0.3rem;
		cursor: pointer;
		padding: 0rem 1rem;

		span {
			color: $white;
			font-family: $spline;
			letter-spacing: 0.1rem;
			font-size: 1.2rem;
			font-weight: 500;
		}

		&:hover {
			background-color: $black-1;
		}
	}
</style>
