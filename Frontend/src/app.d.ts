declare global {
	declare namespace App {
		// interface Locals {}
		// interface PageData {}
		interface Error {
			code: string;
			id: string;
		}
		// interface Platform {}
	}
}
export {};
export interface Thread {
	ID: number;
	UnixTime: string;
	LastBump: string;
	Name: string;
	Topic: string;
	Country: string;
	ExtraFlags: string;
	Sticky: boolean;
	Page: number;
	PostCount: number;
	PostImage: PostImage;
	TextRaw: string;
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
	TextRaw: string;
	Country: string;
	ExtraFlags: string;
	ParentThread: number;
	PostImage: PostImage;
	UserHash: string;
	You: boolean;
}

export interface UserPrefences {
	Theme: string;
	Key: string;
	useHashFileName: boolean;
}
