<div class="spinner"></div>

<style lang="scss">
	.spinner {
		width: 3.5rem;
		height: 3.5rem;
		border-radius: 50%;
		border: 0.3rem solid rgba(0, 0, 0, 0.1);
		border-top-color: #3498db;
		animation: spin 1s linear infinite;

		@keyframes spin {
			0% {
				transform: rotate(0deg);
			}
			100% {
				transform: rotate(360deg);
			}
		}
	}
</style>
