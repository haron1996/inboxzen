<div class="vip-emails">
	<div class="top">
		<p>VIP DOMAINS</p>
		<span>Emails from any of these email addresses will be delivered immediately</span>
	</div>
	<div class="emails">
		<div class="email">
			<span>kibetharon41@gmail.com</span>
			<svg
				width="24px"
				height="24px"
				stroke-width="1.5"
				viewBox="0 0 24 24"
				fill="none"
				xmlns="http://www.w3.org/2000/svg"
				color="#000000"
				><path
					d="M6.75827 17.2426L12.0009 12M17.2435 6.75736L12.0009 12M12.0009 12L6.75827 6.75736M12.0009 12L17.2435 17.2426"
					stroke="#000000"
					stroke-width="1.5"
					stroke-linecap="round"
					stroke-linejoin="round"
				></path></svg
			>
		</div>
		<div class="email">
			<span>kibetharon63@gmail.com</span>
			<svg
				width="24px"
				height="24px"
				stroke-width="1.5"
				viewBox="0 0 24 24"
				fill="none"
				xmlns="http://www.w3.org/2000/svg"
				color="#000000"
				><path
					d="M6.75827 17.2426L12.0009 12M17.2435 6.75736L12.0009 12M12.0009 12L6.75827 6.75736M12.0009 12L17.2435 17.2426"
					stroke="#000000"
					stroke-width="1.5"
					stroke-linecap="round"
					stroke-linejoin="round"
				></path></svg
			>
		</div>
		<div class="email">
			<span>haronkibetrutoh@gmail.com</span>
			<svg
				width="24px"
				height="24px"
				stroke-width="1.5"
				viewBox="0 0 24 24"
				fill="none"
				xmlns="http://www.w3.org/2000/svg"
				color="#000000"
				><path
					d="M6.75827 17.2426L12.0009 12M17.2435 6.75736L12.0009 12M12.0009 12L6.75827 6.75736M12.0009 12L17.2435 17.2426"
					stroke="#000000"
					stroke-width="1.5"
					stroke-linecap="round"
					stroke-linejoin="round"
				></path></svg
			>
		</div>
		<input type="email" name="email" id="email" autocomplete="off" placeholder="add email..." />
	</div>
	<div class="bottom">
		<button>
			<span>SAVE</span>
		</button>
	</div>
</div>

<style lang="scss">
	.vip-emails {
		width: 80rem;
		min-height: max-content;
		background-color: $white;
		display: flex;
		flex-direction: column;
		border-radius: 0.3rem;
		padding: 2rem;
		gap: 2rem;
		border: 0.1rem solid $black-4;
		box-shadow: rgba(0, 0, 0, 0.16) 0rem 0.1rem 0.4rem;

		.top {
			display: flex;
			flex-direction: column;
			justify-content: center;
			gap: 0.5rem;

			p {
				font-family: $spline;
				font-size: 1.2rem;
				text-transform: uppercase;
				font-weight: 600;
			}

			span {
				font-size: 1.3rem;
				font-family: $spline;
				font-weight: 500;
				color: $black-1;
			}
		}

		.emails {
			min-height: max-content;
			display: flex;
			flex-flow: row wrap;
			align-content: start;
			align-items: center;
			gap: 1rem;

			.email {
				background-color: #f6f5f5;
				min-height: max-content;
				display: flex;
				align-items: center;
				padding: 0.5rem 0.5rem 0.5rem 0.7rem;
				width: max-content;
				border-radius: 0.5rem;
				gap: 0.3rem;
				box-shadow:
					rgba(0, 0, 0, 0.05) 0rem 0rem 0rem 0.1rem,
					rgb(209, 213, 219) 0rem 0rem 0rem 0.1rem inset;

				span {
					font-size: 1.3rem;
					font-family: $spline;
					font-weight: 500;
					color: $black-1;
				}

				svg {
					height: 1.5rem;
					width: 1.5rem;
					cursor: pointer;
				}
			}

			input {
				border: none;
				outline: none;
			}
		}

		.bottom {
			display: flex;
			align-items: center;
			justify-content: flex-end;

			button {
				background-color: #525fe1;
				border: none;
				height: 4rem;
				width: 10rem;
				border-radius: 0.2rem;
				cursor: pointer;

				span {
					font-family: $spline;
					font-size: 1.2rem;
					text-transform: uppercase;
					font-weight: 500;
					color: $white;
					letter-spacing: 0.1rem;
				}
			}
		}
	}
</style>
