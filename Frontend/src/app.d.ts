// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
// and what to do when importing types
declare namespace App {
	// interface Locals {}
	// interface PageData {}
	// interface Error {}
	// interface Platform {}
}

export interface Thread {
	ID: number;
	UnixTime: string;
	LastBump: string;
	Name: string;
	Text: string;
	Topic: string;
	Country: string;
	ExtraFlags: string;
	Sticky: boolean;
	Page: number;
	PostCount: number;
	PostImage: PostImage;
	UserHash: string;
	You: boolean;
	Posts: Post[];
}

export interface PostImage {
	Filename: string;
	ImageInfo: string;
	ImageHash: string;
}

export interface Post {
	ID: number;
	UnixTime: string;
	Name: string;
	Text: string;
	Country: string;
	ExtraFlags: string;
	ParentThread: number;
	PostImage: PostImage2;
	UserHash: string;
	You: boolean;
}

export interface PostImage2 {
	Filename: string;
	ImageInfo: string;
	ImageHash: string;
}
export interface UserPrefences {
	Theme: string;
	Key: string;
	useHashFileName: boolean;
}
