export interface Root {
    threadPreviews: ThreadPreview[];
}

export interface ThreadPreview {
    ID: number;
    UnixTime: string;
    LastBump: string;
    Name: string;
    Text: string;
    Topic: string;
    Flags: string;
    Sticky: boolean;
    Page: number;
    PostCount: number;
    PostImage: PostImage;
    Hash: string;
    Posts: Post[];
}

export interface PostImage {
    Filename: string;
    ImageInfo: string;
}

export interface Post {
    ID: number;
    UnixTime: string;
    Name: string;
    Text: string;
    Flags: string;
    ParentThread: number;
    PostImage: PostImage2;
    Hash: string;
}

export interface PostImage2 {
    Filename: string;
    ImageInfo: string;
}