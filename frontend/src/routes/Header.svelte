<script lang="ts">
</script>

<header>
	<a href="/" class="logo">
		<span>InboxZen</span>
	</a>
	<button>
		<svg
			width="24px"
			height="24px"
			stroke-width="1.5"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			color="#000000"
			><path
				d="M19 12H12M12 12L15 15M12 12L15 9"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path><path
				d="M19 6V5C19 3.89543 18.1046 3 17 3H7C5.89543 3 5 3.89543 5 5V19C5 20.1046 5.89543 21 7 21H17C18.1046 21 19 20.1046 19 19V18"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path></svg
		>
		<span>Login</span>
	</button>
</header>

<style lang="scss">
	header {
		width: 100dvw;
		min-height: 10vh;
		background-color: $white;
		height: max-content;
		padding: 2rem;
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
				font-size: 1.5rem;
				font-weight: 900;
				text-transform: uppercase;
				color: $black;
				letter-spacing: 0.1rem;
				letter-spacing: 0.1rem;
				font-family: $spline;
			}
		}

		button {
			border: 0.1rem solid $light-black;
			display: flex;
			align-items: center;
			justify-content: center;
			width: max-content;
			min-width: 10rem;
			padding: 0.5rem 1rem;
			border-radius: 0.3rem;
			background-color: $white;
			gap: 0.5rem;
			cursor: pointer;

			svg {
				height: 1.8rem;
				width: 1.8rem;
			}

			span {
				font-size: 1.5rem;
				font-family: $spline;
				text-transform: uppercase;
			}
		}
	}
</style>
