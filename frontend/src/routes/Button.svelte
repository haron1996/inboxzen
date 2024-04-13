<script lang="ts">
	export let borderColor = '#0d1b2a';
	export let width = 0;
	export let height = 0;
	export let padding = 0;
	export let borderRadius = 0;
	export let color = '';
	export let text = '';
	export let fontWeight = 500;
	export let onClick: (e: MouseEvent) => void;
</script>

<button
	style="border: .2rem solid {borderColor};width: {width}rem; height: {height}rem; padding: {padding}rem; border-radius: {borderRadius}rem; color: {color};"
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
		transition: box-shadow 0.3s ease;
		background-color: #fff;

		span {
			font-size: 1.2rem;
			font-family: $spline;
			text-transform: uppercase;
			letter-spacing: 0.1rem;
		}
	}
</style>
