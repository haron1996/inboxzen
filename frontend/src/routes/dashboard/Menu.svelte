<script lang="ts">
	function showAccountSwitcher() {
		const profile = document.getElementById('profile') as HTMLDivElement | null;

		if (profile === null) return;

		const domRect = profile.getBoundingClientRect();

		const accountSwitcher = document.getElementById('accountSwitcher') as HTMLMenuElement | null;

		if (accountSwitcher === null) return;

		accountSwitcher.style.left = `${domRect.x}px`;
		accountSwitcher.style.left = `${domRect.x}px`;
		accountSwitcher.style.top = `${domRect.bottom + 5}px`;
		accountSwitcher.style.width = `${domRect.width}px`;

		const menu = document.getElementById('menu') as HTMLMenuElement | null;

		if (menu === null) return;

		menu.style.display = 'none';

		accountSwitcher.style.display = 'flex';
	}

	function goToAccount() {
		location.href = '/dashboard/profile';
	}

	function goToSettings() {
		location.href = '/dashboard/settings';
	}

	function logout() {
		alert('log out');
	}
</script>

<menu id="menu">
	<div class="support">
		<svg
			width="24px"
			height="24px"
			stroke-width="1.5"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			color="#000000"
			><path
				d="M12 22C17.5228 22 22 17.5228 22 12C22 6.47715 17.5228 2 12 2C6.47715 2 2 6.47715 2 12C2 17.5228 6.47715 22 12 22Z"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path><path
				d="M9 9C9 5.49997 14.5 5.5 14.5 9C14.5 11.5 12 10.9999 12 13.9999"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path><path
				d="M12 18.01L12.01 17.9989"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path></svg
		>
		<span>Support</span>
	</div>
	<div class="settings" role="none" on:click={goToSettings}>
		<svg
			width="24px"
			height="24px"
			stroke-width="1.5"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			color="#000000"
			><path
				d="M12 15C13.6569 15 15 13.6569 15 12C15 10.3431 13.6569 9 12 9C10.3431 9 9 10.3431 9 12C9 13.6569 10.3431 15 12 15Z"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path><path
				d="M19.6224 10.3954L18.5247 7.7448L20 6L18 4L16.2647 5.48295L13.5578 4.36974L12.9353 2H10.981L10.3491 4.40113L7.70441 5.51596L6 4L4 6L5.45337 7.78885L4.3725 10.4463L2 11V13L4.40111 13.6555L5.51575 16.2997L4 18L6 20L7.79116 18.5403L10.397 19.6123L11 22H13L13.6045 19.6132L16.2551 18.5155C16.6969 18.8313 18 20 18 20L20 18L18.5159 16.2494L19.6139 13.598L21.9999 12.9772L22 11L19.6224 10.3954Z"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path></svg
		>
		<span>Settings</span>
	</div>
	<div class="account" role="none" on:click={goToAccount}>
		<svg
			width="24px"
			height="24px"
			stroke-width="1.5"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			color="#000000"
			><path
				d="M11.6065 2.34184C11.8322 2.14578 12.1678 2.14578 12.3935 2.34184L14.3418 4.03445C14.4646 4.14111 14.6254 4.19336 14.7874 4.17924L17.3586 3.9551C17.6564 3.92913 17.9279 4.1264 17.9953 4.41768L18.5766 6.93224C18.6133 7.09069 18.7126 7.22748 18.852 7.31129L21.0639 8.64123C21.3201 8.79529 21.4238 9.11447 21.307 9.3897L20.2994 11.7657C20.2359 11.9155 20.2359 12.0845 20.2994 12.2343L21.307 14.6103C21.4238 14.8855 21.3201 15.2047 21.0639 15.3588L18.852 16.6887C18.7126 16.7725 18.6133 16.9093 18.5766 17.0678L17.9953 19.5823C17.9279 19.8736 17.6564 20.0709 17.3586 20.0449L14.7874 19.8208C14.6254 19.8066 14.4646 19.8589 14.3418 19.9655L12.3935 21.6582C12.1678 21.8542 11.8322 21.8542 11.6065 21.6582L9.65816 19.9655C9.53539 19.8589 9.37458 19.8066 9.21256 19.8208L6.64142 20.0449C6.34359 20.0709 6.07208 19.8736 6.00474 19.5823L5.42338 17.0678C5.38675 16.9093 5.28736 16.7725 5.14798 16.6887L2.93615 15.3588C2.67993 15.2047 2.57622 14.8855 2.69295 14.6103L3.70065 12.2343C3.76414 12.0845 3.76414 11.9155 3.70065 11.7657L2.69295 9.3897C2.57622 9.11447 2.67993 8.79529 2.93615 8.64123L5.14798 7.31129C5.28736 7.22748 5.38675 7.09069 5.42338 6.93224L6.00474 4.41768C6.07208 4.1264 6.34359 3.92913 6.64142 3.9551L9.21256 4.17924C9.37458 4.19336 9.53539 4.14111 9.65816 4.03445L11.6065 2.34184Z"
				stroke="#000000"
				stroke-width="1.5"
			></path><path
				d="M9 13L11 15L16 10"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path></svg
		>
		<span>Account</span>
	</div>
	<div class="logout" role="none" on:click={logout}>
		<svg
			width="24px"
			height="24px"
			stroke-width="1.5"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			color="#000000"
			><path
				d="M12 12H19M19 12L16 15M19 12L16 9"
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
		<span>Sign Out</span>
	</div>
	<div class="switch-account" role="none" on:click={showAccountSwitcher}>
		<svg
			width="24px"
			height="24px"
			viewBox="0 0 24 24"
			stroke-width="1.5"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			color="#000000"
			><path
				d="M13 6L19 12L13 18"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path><path
				d="M5 6L11 12L5 18"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path></svg
		>
		<span>Switch Account</span>
	</div>
	<div class="add-account" role="none">
		<svg
			width="24px"
			height="24px"
			stroke-width="1.5"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			color="#000000"
			><path
				d="M9 12H12M15 12H12M12 12V9M12 12V15"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path><path
				d="M11.7 1.1732C11.8856 1.06603 12.1144 1.06603 12.3 1.17321L21.2263 6.3268C21.4119 6.43397 21.5263 6.63205 21.5263 6.84641V17.1536C21.5263 17.3679 21.4119 17.566 21.2263 17.6732L12.3 22.8268C12.1144 22.934 11.8856 22.934 11.7 22.8268L2.77372 17.6732C2.58808 17.566 2.47372 17.3679 2.47372 17.1536V6.84641C2.47372 6.63205 2.58808 6.43397 2.77372 6.32679L11.7 1.1732Z"
				stroke="#000000"
				stroke-width="1.5"
				stroke-linecap="round"
				stroke-linejoin="round"
			></path></svg
		>
		<span>Add New Account</span>
	</div>
</menu>

<style lang="scss">
	menu {
		position: absolute;
		background-color: $white;
		box-shadow:
			rgba(255, 255, 255, 0.2) 0rem 0rem 0rem 0.1rem inset,
			rgba(0, 0, 0, 0.9) 0rem 0rem 0rem 0.1rem;
		display: none;
		flex-direction: column;
		min-height: max-content;
		cursor: pointer;
		border-radius: 0.3rem;

		div {
			width: 100%;
			min-height: 4.5rem;
			display: flex;
			align-items: center;
			padding: 1rem;
			gap: 1rem;

			svg {
				height: 2rem;
				width: 2rem;

				path {
					stroke: $black-1;
				}
			}

			span {
				font-family: $spline;
				font-size: 1.2rem;
				color: $black;
				text-transform: uppercase;
				font-weight: 500;
			}

			&:hover {
				background-color: $gray;
			}
		}
	}
</style>
