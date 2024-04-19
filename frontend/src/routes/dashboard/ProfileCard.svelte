<script lang="ts">
	export let name = '';
	export let email = '';
	export let picture = '';
	export let width = '';
	export let SVGDisplay = '';
	export let onClick: (e: MouseEvent) => void;
</script>

<div
	class="profile"
	id="profile"
	style="width: {width};"
	role="none"
	on:click|preventDefault|stopPropagation={onClick}
>
	<div class="left">
		<div class="image">
			<img src={picture} alt="profile" referrerPolicy="no-referrer" />
		</div>
	</div>
	<div class="center" id="center">
		<span class="name">{name}</span>
		<span class="email">{email}</span>
	</div>
	<div class="right">
		<svg
			id="profile-card-svg"
			width="24px"
			height="24px"
			stroke-width="1.5"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			color="#000000"
			style="display: {SVGDisplay};"
			><path
				d="M6 9L12 15L18 9"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path></svg
		>
	</div>
</div>

<style lang="scss">
	.profile {
		display: flex;
		align-items: center;
		width: max-content;
		height: max-content;
		gap: 1rem;
		padding: 0.5rem 1rem;
		border-radius: 0.3rem;
		cursor: pointer;
		position: relative;
		max-width: 40rem;
		background-color: #00a8e8;
		border: 0.2rem solid #0077b6;

		.left {
			display: flex;
			align-items: center;
			justify-content: center;
			height: 3.5rem;
			width: 3.5rem;
			border-radius: 50%;
			letter-spacing: 0.1rem;

			img {
				max-inline-size: 100%;
				object-fit: cover;
				border-radius: 50%;
			}
		}

		.center {
			display: flex;
			flex-direction: column;
			justify-content: center;

			span {
				font-size: 1.2rem;
				font-weight: 500;
				font-family: $spline;
				color: $black;
				white-space: nowrap;
				overflow: hidden;
				text-overflow: ellipsis;
				text-transform: uppercase;
				max-width: 35rem;
			}

			span.email {
				color: $light-black;
				text-transform: lowercase;
			}
		}

		.right {
			display: flex;
			align-items: center;
			justify-content: center;

			svg {
				border-radius: 50%;
				border: 0.1rem solid rgba(0, 0, 0, 0.1);
				height: 2rem;
				width: 2rem;
				transition: all 0.3s ease-in-out;

				path {
					stroke: $light-black;
				}
			}
		}

		&:hover {
			background-color: #1e91d6;
		}
	}
</style>
