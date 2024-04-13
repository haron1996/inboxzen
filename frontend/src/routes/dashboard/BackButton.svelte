<script lang="ts">
	export let width = 0;
	export let onClick: () => void;
</script>

<button style="width: {width}rem;" on:click={onClick}>
	<span>BACK TO DASHBOARD</span>
</button>

<style lang="scss">
	button {
		display: flex;
		align-items: center;
		justify-content: center;
		min-width: 10rem;
		gap: 1rem;
		height: 4rem;
		border-radius: 0.3rem;
		cursor: pointer;
		padding: 0rem 1rem;
		background-color: #fff;
		transition: box-shadow 0.3s ease;
		border: 0.2rem solid #0d1b2a;

		span {
			color: #0d1b2a;
			font-family: $spline;
			letter-spacing: 0.1rem;
			font-size: 1.2rem;
			font-weight: 500;
		}
	}
</style>
