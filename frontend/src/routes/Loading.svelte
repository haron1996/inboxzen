<div class="wrapper">
	<div class="loader" />
</div>

<style lang="scss">
	.wrapper {
		position: fixed;
		left: 0;
		top: 0;
		width: 100dvw;
		height: 100dvh;
		background-color: $white;
		z-index: 1000;
		display: flex;
		align-items: center;
		justify-content: center;

		.loader {
			border: 0.4rem solid #f3f3f3;
			border-top: 0.4rem solid $blue;
			border-radius: 50%;
			width: 4rem;
			height: 4rem;
			animation: spin 1s linear infinite;
			margin: auto;

			@keyframes spin {
				0% {
					transform: rotate(0deg);
				}
				100% {
					transform: rotate(360deg);
				}
			}
		}
	}
</style>
