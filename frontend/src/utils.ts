import { goto } from '$app/navigation';
import { URL, errorMessages, session, successMessages } from './store';
import type { Session } from './types';

let url: string = '';
let result: any;
let message = '';
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

function getURL() {
	const unsub = URL.subscribe((val) => {
		url = val;
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

	const svg = document.getElementById('profile-card-svg') as SVGAElement | null;

	if (svg === null) return;

	switch (menu.style.display) {
		case 'flex':
			menu.style.display = 'none';

			svg.style.transform = 'rotate(0)';
			break;
		default:
			menu.style.display = 'flex';

			svg.style.transform = 'rotate(0.5turn)';
			break;
	}
}

export function hideMenu() {
	const menu = document.getElementById('menu') as HTMLMenuElement | null;

	if (menu === null) return;

	if (!(menu.style.display === 'flex')) return;

	menu.style.display = 'none';

	const svg = document.getElementById('profile-card-svg') as SVGAElement | null;

	if (svg === null) return;

	svg.style.transform = 'rotate(0)';
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

// FETCH REQUESTS

// check user login status
export const checkUserLoginStatus = async () => {
	getURL();

	url = `${url}/private/checkloginstatus`;

	await fetch(url, {
		method: 'GET',
		mode: 'cors',
		cache: 'no-cache',
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json'
		},
		redirect: 'follow',
		referrerPolicy: 'no-referrer'
	})
		.then(async (response) => {
			if (!response.ok) {
				result = await response.json();
				message = result.message;
				throw new Error(message);
			}

			return response.json();
		})
		.then((data) => {
			goto('/dashboard', { replaceState: true, invalidateAll: true });
		})
		.catch((error) => {
			message = error.message;

			switch (message) {
				case 'token has expired':
					goto('/preauth', { replaceState: true, invalidateAll: true });
					return;
				default:
					updateErrorMessages(message);
					break;
			}
		});
};

// get auth url
export const getAuthURL = async () => {
	getURL();

	url = `${url}/user/getauthurl`;

	await fetch(url, {
		method: 'GET',
		mode: 'cors',
		cache: 'no-cache',
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json'
		},
		redirect: 'follow',
		referrerPolicy: 'no-referrer'
	})
		.then(async (response) => {
			if (!response.ok) {
				result = await response.json();
				message = result.message;
				throw new Error(message);
			}

			return response.json();
		})
		.then((data: string) => {
			location.replace(data);
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
		});
};

// finish auth
export const finishAuth = async (code: string) => {
	getURL();

	url = `${url}/user/comletegoogleauth`;

	await fetch(url, {
		method: 'POST',
		mode: 'cors',
		cache: 'no-cache',
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json'
		},
		redirect: 'follow',
		referrerPolicy: 'no-referrer',
		body: JSON.stringify({ code: code })
	})
		.then(async (response) => {
			if (!response.ok) {
				result = await response.json();
				message = result.message;
				throw new Error(message);
			}

			return response.json();
		})
		.then((data) => {
			goto('/dashboard', { replaceState: true, invalidateAll: true });
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
			goto('/preauth', { replaceState: true, invalidateAll: true });
		});
};

// get user account details
export const getUserAccount = async () => {
	getURL();

	url = `${url}/private/getuseraccount`;

	await fetch(url, {
		method: 'GET',
		mode: 'cors',
		cache: 'no-cache',
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json'
		},
		redirect: 'follow',
		referrerPolicy: 'no-referrer'
	})
		.then(async (response) => {
			if (!response.ok) {
				result = await response.json();
				message = result.message;
				throw new Error(message);
			}

			return response.json();
		})
		.then((data: Session) => {
			session.set(data);
		})
		.catch((error) => {
			message = error.message;
			switch (message) {
				case 'token has expired':
					updateErrorMessages(message);
					goto('/preauth', { replaceState: true, invalidateAll: true });
					return;
				default:
					updateErrorMessages(message);
					break;
			}
		});
};

// get user email settings (domains, emails, keywords, delivery times)
export const getUserEmailSettings = async () => {
	getURL();

	const urls = [
		`${url}/private/getuseraccount`,
		`${url}/private/getdomains`,
		`${url}/private/getemails`,
		`${url}/private/getkeywords`
	];

	urls.forEach((url) => {
		console.log(url);
	});
};
