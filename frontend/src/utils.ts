import { goto } from '$app/navigation';
import {
	URL,
	domains,
	emails,
	errorMessages,
	keywords,
	loading,
	running,
	session,
	successMessages,
	time,
	times
} from './store';
import type { Domain, Email, Keyword, Session, Time } from './types';

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

	// menu.style.left = `${domRect.x}px`;
	// menu.style.left = `${domRect.x}px`;
	// menu.style.top = `${domRect.bottom + 5}px`;
	// menu.style.width = `${domRect.width}px`;

	// Set menu width to be wider
	const menuWidth = domRect.width * 1.5; // Adjust the multiplier as needed
	menu.style.width = `${menuWidth}px`;

	// Calculate left position for the menu to span equally on both sides
	const leftPosition = domRect.x - (menuWidth - domRect.width) / 2;
	menu.style.left = `${leftPosition}px`;

	menu.style.top = `${domRect.bottom + 5}px`;

	const svg = target.lastElementChild?.firstElementChild as SVGAElement | null;

	if (svg === null) return;

	switch (menu.style.display) {
		case 'block':
			menu.style.display = 'none';
			svg.style.transform = 'rotate(0)';
			break;
		default:
			menu.style.display = 'block';
			menu.classList.add('animate__bounceIn');
			svg.style.transform = 'rotate(0.5turn)';
			break;
	}
}

export function hideMenu() {
	const menu = document.getElementById('menu') as HTMLMenuElement | null;

	if (menu === null) return;

	if (!(menu.style.display === 'block')) return;

	menu.style.display = 'none';

	const profileCards = document.querySelectorAll(
		'div.profile'
	) as NodeListOf<HTMLDivElement> | null;

	if (profileCards === null) return;

	const lastProfileCard = profileCards[profileCards.length - 1];

	const svg = lastProfileCard.lastElementChild?.firstElementChild as SVGAElement | null;

	if (svg === null) return;

	svg.style.transform = 'rotate(0)';
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
			location.href = '/dashboard';
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
		})
		.finally(() => {
			loading.set(false);
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
		})
		.finally(() => {
			loading.set(false);
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
			location.href = '/dashboard';
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
			goto('/preauth', { replaceState: true, invalidateAll: true });
		})
		.finally(() => {
			loading.set(false);
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
		})
		.finally(() => {
			loading.set(false);
		});
};

// update vip domains
export const updateVipDomains = async (doms: Domain[]) => {
	getURL();

	url = `${url}/private/updatedomains`;

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
		body: JSON.stringify({
			domain_names: doms.map((d) => d.domain_name)
		})
	})
		.then(async (response) => {
			if (!response.ok) {
				result = await response.json();
				message = result.message;
				throw new Error(message);
			}

			return response.json();
		})
		.then((data: Domain[]) => {
			domains.set(data);
			updateSuccessMessages('domains updated successfully');
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
		})
		.finally(() => {
			loading.set(false);
		});
};

// get vip domains
export const getVipDomains = async () => {
	getURL();

	url = `${url}/private/getdomains`;

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
		.then((data: Domain[]) => {
			domains.set(data);
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
		})
		.finally(() => {
			loading.set(false);
		});
};

// update vip email addresses
export const updateVipEmailAddresses = async (es: Email[]) => {
	getURL();

	url = `${url}/private/updateemails`;

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
		body: JSON.stringify({
			emails: es.map((e) => e.vip_email_address)
		})
	})
		.then(async (response) => {
			if (!response.ok) {
				result = await response.json();
				message = result.message;
				throw new Error(message);
			}

			return response.json();
		})
		.then((data: Email[]) => {
			emails.set(data);
			updateSuccessMessages('emails updated successfully');
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
		})
		.finally(() => {
			loading.set(false);
		});
};

// get vip email addresses
export const getVipEmailAddresses = async () => {
	getURL();

	url = `${url}/private/getemails`;

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
			emails.set(data);
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
		})
		.finally(() => {
			loading.set(false);
		});
};

// update vip keywords
export const updateVipKeywords = async (kws: Keyword[]) => {
	getURL();

	url = `${url}/private/updatekeywords`;

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
		body: JSON.stringify({
			keywords: kws.map((kw) => kw.keyword)
		})
	})
		.then(async (response) => {
			if (!response.ok) {
				result = await response.json();
				message = result.message;
				throw new Error(message);
			}

			return response.json();
		})
		.then((data: Keyword[]) => {
			keywords.set(data);
			updateSuccessMessages('keywords updated successfully');
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
		})
		.finally(() => {
			loading.set(false);
		});
};

// get vip keywords
export const getVipKeywords = async () => {
	getURL();

	url = `${url}/private/getkeywords`;

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
		.then((data: Keyword[]) => {
			keywords.set(data);
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
		})
		.finally(() => {
			loading.set(false);
		});
};

// set inbox delivery times
export const setInboxDeliveryTime = async (params: {}, dts: Time[]) => {
	getURL();

	url = `${url}/private/setdeliverytime`;

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
		body: JSON.stringify(params)
	})
		.then(async (response) => {
			if (!response.ok) {
				result = await response.json();
				message = result.message;
				throw new Error(message);
			}

			return response.json();
		})
		.then((data: Time) => {
			if (dts === null) {
				let ts: Time[] = [];
				ts = [data, ...ts];
				times.set(ts);
			} else {
				times.update((t) => [data, ...t]);
			}

			time.set({});
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
		})
		.finally(() => {
			loading.set(false);
		});
};

// get inbox delivery times
export const getInboxDeliveryTimes = async () => {
	getURL();

	url = `${url}/private/getdeliverytimes`;

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
		.then((data: Time[]) => {
			times.set(data);
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
		})
		.finally(() => {
			loading.set(false);
		});
};

// delete inbox delivery time
export const deleteInboxDeliveryTime = async (params: {}, dts: Time[]) => {
	getURL();

	url = `${url}/private/deletedeliverytime`;

	await fetch(url, {
		method: 'DELETE',
		mode: 'cors',
		cache: 'no-cache',
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json'
		},
		redirect: 'follow',
		referrerPolicy: 'no-referrer',
		body: JSON.stringify(params)
	})
		.then(async (response) => {
			if (!response.ok) {
				result = await response.json();
				message = result.message;
				throw new Error(message);
			}

			return response.json();
		})
		.then((data: Time) => {
			dts = dts.filter((dt) => dt.id !== data.id);
			times.set(dts);
			updateSuccessMessages('time deleted successfully');
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
		})
		.finally(() => {
			loading.set(false);
		});
};

// activate
export const activate = async () => {
	getURL();

	url = `${url}/private/activate`;

	await fetch(url, {
		method: 'PATCH',
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
			running.set(true);
			updateSuccessMessages('service started successfully');
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
		})
		.finally(() => {
			loading.set(false);
		});
};

// get status
export const getServiceStatus = async () => {
	getURL();

	url = `${url}/private/checkstatus`;

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
			if (data.running) {
				running.set(true);
			}
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
		})
		.finally(() => {
			loading.set(false);
		});
};

// switch account
export const switchAccount = async (e: MouseEvent) => {
	getURL();

	const target = e.currentTarget as HTMLDivElement;

	const email = target.children[1].children[1].textContent;

	if (email === null) return;

	url = `${url}/private/switchaccount/${email}`;

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
			updateErrorMessages(message);
		})
		.finally(() => {
			hideMenu();
			loading.set(false);
			location.reload();
		});
};

// get first time senders
export const getFirstTimeSenders = async () => {
	getURL();

	url = `${url}/private/getfirsttimesenders`;

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
			updateErrorMessages(message);
		})
		.finally(() => {
			hideMenu();
			loading.set(false);
		});
};

// get user email settings (domains, emails, keywords, delivery times)
export const getUserEmailSettings = async () => {
	try {
		await Promise.all([
			getUserAccount(),
			getServiceStatus(),
			getInboxDeliveryTimes(),
			getVipDomains(),
			getVipEmailAddresses(),
			getVipKeywords()
		]);
	} catch (error) {
		console.log('An error occured:', error);
	}
};
