<script lang="ts">
	export let width = 0;
	export let height = 0;
	export let backgroundColor = '';
	export let padding = 0;
	export let borderRadius = 0;
	export let color = '';
	export let text = '';
	export let onClick: () => void;
</script>

<button
	style="width: {width}rem; height: {height}rem; background-color: {backgroundColor}; padding: {padding}rem; border-radius: {borderRadius}rem; color: {color}"
	on:click|preventDefault={onClick}
>
	<span>{text}</span>
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

		span {
			font-size: 1.2rem;
			font-family: $spline;
			font-weight: 500;
			text-transform: uppercase;
			letter-spacing: 0.1rem;
		}
	}
</style>
