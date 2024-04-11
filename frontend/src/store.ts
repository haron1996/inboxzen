import { writable } from 'svelte/store';
import type { Session } from './types';

export const URL = writable<string>('http://localhost:5000');

export const errorMessages = writable<string[]>([]);

export const successMessages = writable<string[]>([]);

export const loading = writable<boolean>(false);

export const session = writable<Session>({});
