<script lang="ts">
	// export let borderColor = '#e5e5e5';
	// export let backgroundColor = '#f8f9fa';
	export let width = 0;
	export let height = 0;
	export let padding = 0;
	export let borderRadius = 0;
	export let color = '';
	export let text = '';
	export let fontWeight = 500;
	export let onClick: (e: MouseEvent) => void;

	// function handleMouseEnter(e: MouseEvent) {
	// 	const t = e.currentTarget as HTMLButtonElement;

	// 	t.style.backgroundColor = backgroundColor;

	// 	const span = t.lastElementChild as HTMLSpanElement | null;

	// 	if (span === null) return;

	// 	span.style.color = '#fff';
	// }

	// function handleMouserLeave(e: MouseEvent) {
	// 	const t = e.currentTarget as HTMLButtonElement;

	// 	t.style.backgroundColor = '#fff';

	// 	const span = t.lastElementChild as HTMLSpanElement | null;

	// 	if (span === null) return;

	// 	span.style.color = color;
	// }
</script>

<button
	style="width: {width}rem; height: {height}rem; padding: {padding}rem; border-radius: {borderRadius}rem; color: {color};"
	on:click|preventDefault={onClick}
	on:focus
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
		transition: all 0.3s linear;
		border: 0.2rem solid #e5e5e5;
		background-color: #f8f9fa;

		span {
			font-size: 1.2rem;
			font-family: $spline;
			text-transform: uppercase;
			letter-spacing: 0.1rem;
			color: #757575;
		}

		&:hover {
			border-color: #adb5bd;
			background-color: #e5e5e5;
		}
	}
</style>
