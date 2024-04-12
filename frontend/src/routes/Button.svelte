<script lang="ts">
	export let width = 0;
	export let height = 0;
	export let backgroundColor = '';
	export let padding = 0;
	export let borderRadius = 0;
	export let color = '';
	export let text = '';
	export let fontWeight = 500;
	export let onClick: () => void;
</script>

<button
	style="width: {width}rem; height: {height}rem; background-color: {backgroundColor}; padding: {padding}rem; border-radius: {borderRadius}rem; color: {color}"
	on:click|preventDefault={onClick}
>
	<span style="font-weight: {fontWeight};">{text}</span>
</button>

<style lang="scss">
	button {
		border: none;
		cursor: pointer;
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 1rem;
		min-width: max-content;
		background-image: linear-gradient(to right top, #4e54c8, #8f94fb, #b4b8fd, #cad3ff, #e5f0ff);
		box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2);
		transition: box-shadow 0.3s ease;

		span {
			font-size: 1.2rem;
			font-family: $spline;
			text-transform: uppercase;
			letter-spacing: 0.1rem;
		}

		&:hover {
			box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.2);
		}
	}
</style>
