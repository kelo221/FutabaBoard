import { writable } from 'svelte/store';
import type { Thread, Post, UserPrefences } from './app';
import type { Writable } from 'svelte/store';
import { localStorageStore } from '@skeletonlabs/skeleton';

const replyState = {
	threadID: 0,
	content: '',
	open: false,
	newThread: false
};

interface PostPreview {
	open: boolean;
	postData: Post;
}

const currentPage = 0;

export const currentPageStore = writable(currentPage);
export const replyBoxStore = writable(replyState);
export const currentThreadStore: Writable<Thread> = writable();
export const postPreview: Writable<PostPreview> = writable({ open: false });
export const userSettings: Writable<UserPrefences> = localStorageStore('userPrefs', {
	Theme: 'modern',
	Key: '',
	useHashFileName: false
});
