import { writable } from 'svelte/store';
import type { session } from './types';

export const URL = writable<string>('http://localhost:5000');

export const errors = writable<string[]>([]);

export const loading = writable<boolean>(false);

export const userSession = writable<session>({});
