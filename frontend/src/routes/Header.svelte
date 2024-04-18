<script lang="ts">
	import { checkUserLoginStatus } from '../utils';
</script>

<header>
	<div class="inner-container">
		<a href="/" class="logo">
			<svg
				width="24px"
				height="24px"
				viewBox="0 0 24 24"
				stroke-width="1.5"
				fill="none"
				xmlns="http://www.w3.org/2000/svg"
				color="#011627"
				><path
					d="M3 20.4V3.6C3 3.26863 3.26863 3 3.6 3H20.4C20.7314 3 21 3.26863 21 3.6V20.4C21 20.7314 20.7314 21 20.4 21H3.6C3.26863 21 3 20.7314 3 20.4Z"
					stroke="#011627"
					stroke-width="1.5"
				></path><path
					d="M9.5 14V10C9.5 8.89543 10.3954 8 11.5 8H12.5C13.6046 8 14.5 8.89543 14.5 10V14C14.5 15.1046 13.6046 16 12.5 16H11.5C10.3954 16 9.5 15.1046 9.5 14Z"
					stroke="#011627"
					stroke-width="1.5"
				></path></svg
			>
			<span>GMAIL INBOX ZERO</span>
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
		border-bottom: 0.1rem solid #dad7cd;

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
				gap: 1rem;
				text-decoration: none;

				span {
					font-family: $spline;
					font-size: 1.5rem;
					font-weight: 500;
					color: #0d1b2a;
					letter-spacing: 0.1rem;
					letter-spacing: 0.1rem;
					text-transform: uppercase;

					&:hover {
						text-decoration: underline;
					}
				}
			}

			span {
				font-family: $spline;
				font-size: 1.5rem;
				font-weight: 500;
				color: #0d1b2a;
				letter-spacing: 0.1rem;
				letter-spacing: 0.1rem;
				text-transform: uppercase;
				cursor: pointer;
			}
		}
	}
</style>
