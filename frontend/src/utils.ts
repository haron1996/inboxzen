import { errorMessages, successMessages } from './store';

let successMsgs: string[] = [];
let errorMsgs: string[] = [];

function getSuccessMsgs() {
	const unsub = successMessages.subscribe((msgs) => {
		successMsgs = msgs;
	});

	unsub();
}

function getErrorMsgs() {
	const unsub = errorMessages.subscribe((msgs) => {
		errorMsgs = msgs;
	});

	unsub();
}

export function showMenu(e: MouseEvent) {
	const target = e.currentTarget as HTMLDivElement;

	const domRect = target.getBoundingClientRect();

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

export function scrollPageToTop() {
	window.scrollTo({
		top: 0,
		behavior: 'smooth'
	});
}

export function showError(message: string) {
	const div = document.createElement('div');

	div.innerText = message;

	div.classList.add('success');

	document.body.appendChild(div);

	const successAlerts = document.querySelectorAll('.success') as NodeListOf<HTMLDivElement> | null;

	if (successAlerts === null) return;

	const lastAlert = successAlerts[successAlerts.length - 1];

	const rect = lastAlert.getBoundingClientRect();

	console.log(rect);
}

export function updateSuccessMessages(message: string) {
	successMessages.update((msgs) => [message, ...msgs]);
	removeSuccessMessage();
}

function removeSuccessMessage() {
	getSuccessMsgs();

	setTimeout(() => {
		successMsgs.pop();
		successMessages.set(successMsgs);
	}, 3000);
}

export function updateErrorMessages(message: string) {
	errorMessages.update((msgs) => [message, ...msgs]);
	removeErrorMessage();
}

function removeErrorMessage() {
	getErrorMsgs();

	setTimeout(() => {
		errorMsgs.pop();
		errorMessages.set(errorMsgs);
	}, 3000);
}
