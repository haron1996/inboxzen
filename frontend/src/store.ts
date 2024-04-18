import { writable } from 'svelte/store';
import type { Domain, Email, Keyword, Session, Time } from './types';

// baseurl
export const URL = writable<string>('http://localhost:5000');

// error alerts
export const errorMessages = writable<string[]>([]);

// success alerts
export const successMessages = writable<string[]>([]);

// loading state
export const loading = writable<boolean>(true);

// user session
export const session = writable<Session>({});

// vip domains
export const domain = writable<Domain>({});
export let domains = writable<Domain[]>([]);
// end of vip domains

// vip  email addresses
export let email = writable<Email>({});
export let emails = writable<Email[]>([]);
// end of vip email addresses

// vip keywords
export let keyword = writable<Keyword>({});
export let keywords = writable<Keyword[]>([]);
// end of vip keywords

// delivery times
export let time = writable<Time>({});
export let times = writable<Time[]>([]);
// end of delivery times

// status
export const running = writable<boolean>(false);
