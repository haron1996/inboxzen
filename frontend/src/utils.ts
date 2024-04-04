export function showMenu() {
	hideAccountSwitcher();

	const profile = document.getElementById('profile') as HTMLDivElement | null;

	if (profile === null) return;

	const domRect = profile.getBoundingClientRect();

	const menu = document.getElementById('menu') as HTMLMenuElement | null;

	if (menu === null) return;

	menu.style.left = `${domRect.x}px`;
	menu.style.left = `${domRect.x}px`;
	menu.style.top = `${domRect.bottom + 5}px`;
	menu.style.width = `${domRect.width}px`;
	switch (menu.style.display) {
		case 'flex':
			menu.style.display = 'none';
			break;
		default:
			menu.style.display = 'flex';
			break;
	}
}

export function hideMenu() {
	const menu = document.getElementById('menu') as HTMLMenuElement | null;

	if (menu === null) return;

	menu.style.display = 'none';
}

export function hideAccountSwitcher() {
	const accountSwitcher = document.getElementById('accountSwitcher') as HTMLMenuElement | null;

	if (accountSwitcher === null) return;

	accountSwitcher.style.display = 'none';
}

export function scrollPageToTop() {
	window.scrollTo({
		top: 0,
		behavior: 'smooth'
	});
}
