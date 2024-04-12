import { goto } from '$app/navigation';
import {
	URL,
	domains,
	emails,
	errorMessages,
	keywords,
	loading,
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
			goto('/dashboard', { replaceState: true, invalidateAll: true });
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
	if (doms.length < 1) {
		updateErrorMessages('at least one domain required');
		return;
	}

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
	if (es.length < 1) {
		updateErrorMessages('at least one email required');
		return;
	}

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
export const setInboxDeliveryTime = async (params: {}) => {
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
			times.update((t) => [data, ...t]);
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
		.then((data) => {
			console.log(data);
		})
		.catch((error) => {
			message = error.message;
			updateErrorMessages(message);
		})
		.finally(() => {
			loading.set(false);
		});
};

// get user email settings (domains, emails, keywords, delivery times)
export const getUserEmailSettings = async () => {
	try {
		await Promise.all([
			getUserAccount(),
			getVipDomains(),
			getVipEmailAddresses(),
			getVipKeywords()
		]);
	} catch (error) {
		console.log('An error occured:', error);
	}
};
