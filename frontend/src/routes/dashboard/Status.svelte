<script lang="ts">
	import { loading, running } from '../../store';
	import { activate } from '../../utils';
	import CardSkeleton from './CardSkeleton.svelte';
</script>

<div class="status">
	<div class="top">
		<p>STATUS</p>
		<span>Use this button to toggle status</span>
	</div>
	<div class="buttons">
		{#if $loading}
			<CardSkeleton height={10} width={100} padding={0} borderRadius={0} />
		{:else}
			<button class:running={$running} on:click|preventDefault={activate}>
				<div class="pulse"></div>
				<span>{$running ? 'RUNNING' : 'ACTIVATE'}</span>
			</button>
		{/if}
	</div>
</div>

<style lang="scss">
	.status {
		width: 70%;
		min-height: max-content;
		background-color: $white;
		display: flex;
		flex-direction: column;
		border-radius: 0.3rem;
		border: 0.1rem solid #e5e5e5;
		margin: 1rem auto;

		.top {
			display: flex;
			flex-direction: column;
			justify-content: center;
			gap: 1rem;
			border-bottom: 0.1rem solid #e0e1dd;
			padding: 2rem;

			p {
				font-family: $spline;
				font-size: 1.1rem;
				text-transform: uppercase;
				font-weight: 600;
				color: #0d1b2a;
			}

			span {
				font-family: $spline;
				font-weight: 600;
				text-transform: uppercase;
				color: #757575;
			}
		}

		.buttons {
			min-height: max-content;
			display: flex;
			align-items: center;
			justify-content: space-between;
			padding: 2rem;

			button {
				background-color: #fff;
				border: 0.1rem solid #d90429;
				padding: 1rem;
				cursor: pointer;
				border-radius: 0.3rem;
				min-width: 13rem;
				display: flex;
				align-items: center;
				justify-content: center;
				gap: 1rem;
				background-color: #fff0f3;

				.pulse {
					display: inline-block;
					width: 1.5rem;
					height: 1.5rem;
					background-color: #d90429;
					border-radius: 50%;
				}

				span {
					font-family: $spline;
					font-weight: 500;
					color: #d90429;
					text-transform: uppercase;
					letter-spacing: 0.1rem;
				}
			}

			.running {
				border-color: #20bf55;
				background-color: #f1fffa;

				.pulse {
					background-color: #20bf55;
					animation: pulse 2s infinite;

					@keyframes pulse {
						0% {
							transform: scale(0.9);
							box-shadow: 0 0 0 0 rgba(76, 175, 80, 0.7);
						}
						70% {
							transform: scale(1);
							box-shadow: 0 0 0 15px rgba(76, 175, 80, 0);
						}
						100% {
							transform: scale(0.9);
							box-shadow: 0 0 0 15px rgba(76, 175, 80, 0);
						}
					}
				}

				span {
					color: #20bf55;
				}
			}
		}
	}
</style>
