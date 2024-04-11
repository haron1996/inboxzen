<script lang="ts">
	export let height = 40;
	export let width = 30;
	export let padding = 0;
	export let borderRadius = 0.5;
</script>

<div
	class="card"
	style="height: {height}rem; width: {width}rem; padding: {padding}rem; border-radius={borderRadius}rem"
>
	<div class="content"></div>
</div>

<style lang="scss">
	.card {
		background-color: #f0f0f0;
		overflow: hidden;
		position: relative;
		z-index: 100;

		.content {
			position: absolute;
			top: 0;
			left: 0;
			width: 100%;
			height: 100%;
			background-image: linear-gradient(90deg, #f0f0f0 25%, #e0e0e0 50%, #f0f0f0 75%);
			background-size: 200% 100%;
			animation: shimmer 1.5s infinite;

			@keyframes shimmer {
				0% {
					background-position: -200% 0;
				}
				100% {
					background-position: 200% 0;
				}
			}
		}
	}
</style>
