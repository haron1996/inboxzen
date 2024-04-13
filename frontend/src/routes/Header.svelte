<script lang="ts">
	import { checkUserLoginStatus } from '../utils';
</script>

<header>
	<div class="inner-container">
		<a href="/" class="logo">
			<span>ZENN</span>
		</a>
		<span role="none" on:click={checkUserLoginStatus}>Sign in</span>
	</div>
</header>

<style lang="scss">
	header {
		width: 100dvw;
		min-height: 10vh;
		background-color: $white;
		height: max-content;
		display: flex;
		align-items: center;
		justify-content: center;
		font-family: 'Jost', sans-serif;
		font-optical-sizing: auto;

		.inner-container {
			width: 70%;
			display: flex;
			align-items: center;
			justify-content: space-between;

			a.logo {
				display: flex;
				align-items: center;
				justify-content: center;
				width: max-content;
				gap: 0.5rem;
				text-decoration: none;

				span {
					font-size: 1.2rem;
					color: #0d1b2a;
					letter-spacing: 0.1rem;
					letter-spacing: 0.1rem;
					font-weight: 600;
					text-transform: uppercase;
				}
			}

			span {
				font-size: 1.2rem;
				color: #0d1b2a;
				letter-spacing: 0.1rem;
				letter-spacing: 0.1rem;
				text-transform: uppercase;
				font-weight: 500;
				cursor: pointer;
			}
		}
	}
</style>
