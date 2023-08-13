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
	Post: Post[];
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
}

export interface PostImage2 {
	Filename: string;
	ImageInfo: string;
	ImageHash: string;
}
