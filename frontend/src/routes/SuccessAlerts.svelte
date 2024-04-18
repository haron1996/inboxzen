<script lang="ts">
	import { successMessages } from '../store';
</script>

<div class="alerts">
	{#each $successMessages as msg}
		<div class="alert" id="success">
			<span>{msg}</span>
		</div>
	{/each}
</div>

<style lang="scss">
	.alerts {
		z-index: 1000;
		position: fixed;
		left: 50%;
		bottom: 2%;
		transform: translateX(-50%);
		min-width: max-content;
		display: flex;
		flex-direction: column;
		gap: 1rem;

		.alert {
			box-shadow:
				rgba(9, 30, 66, 0.25) 0px 4px 8px -2px,
				rgba(9, 30, 66, 0.08) 0px 0px 0px 1px;
			background-color: #fff;
			border-radius: 0.3rem;
			transition: all 0.3s ease-in-out;
			display: flex;
			align-items: center;
			justify-content: center;
			width: 30rem;
			height: 5rem;
			padding: 1rem;
			border: 0.2rem solid #20bf55;
			border-radius: 0.3rem;

			span {
				font-family: $spline;
				font-size: 1.1rem;
				font-weight: 600;
				color: #20bf55;
				text-transform: uppercase;
			}
		}
	}
</style>
