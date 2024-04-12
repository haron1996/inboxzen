<script lang="ts">
	export let width = 0;
	export let onClick: () => void;
</script>

<button style="width: {width}rem;" on:click={onClick}>
	<span>BACK TO DASHBOARD</span>
</button>

<style lang="scss">
	button {
		border: none;
		display: flex;
		align-items: center;
		justify-content: center;
		min-width: 10rem;
		gap: 1rem;
		height: 4rem;
		border-radius: 0.3rem;
		cursor: pointer;
		padding: 0rem 1rem;
		background-color: #4e54c8;
		background-image: linear-gradient(to right top, #4e54c8, #8f94fb, #b4b8fd, #cad3ff, #e5f0ff);
		box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2);
		transition: box-shadow 0.3s ease;

		span {
			color: $white;
			font-family: $spline;
			letter-spacing: 0.1rem;
			font-size: 1.2rem;
			font-weight: 500;
		}

		&:hover {
			box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.2);
		}
	}
</style>
