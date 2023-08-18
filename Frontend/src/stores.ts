import { writable } from 'svelte/store';
import type { Thread } from './app';
import type { Writable } from 'svelte/store';

const replyState = {
	threadID: 0,
	content: '',
	open: false
};

export const replyBoxStore = writable(replyState);
export const currentThreadStore: Writable<Thread> = writable();
