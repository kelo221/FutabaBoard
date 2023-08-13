package DataModels

import (
	"time"
)

const postsPerThread = 300
const threadsPerPage = 4

type SharedID struct {
	ID int64 `gorm:"primaryKey"`
}

type Thread struct {
	SharedID
	UnixTime   time.Time `json:"UnixTime" gorm:"autoCreateTime"`
	LastBump   time.Time `json:"LastBump" gorm:"autoCreateTime"`
	Name       string    `json:"Name"`
	Text       string    `json:"Text"`
	Topic      string    `json:"Topic"`
	Country    string    `json:"Country"`
	ExtraFlags string    `json:"ExtraFlags"`
	Sticky     bool      `json:"Sticky"`
	Locked     bool      `json:"Locked"`
	Page       int       `json:"Page"`
	PostCount  int       `json:"PostCount"`
	PostImage  PostImage `json:"PostImage,omitempty" gorm:"embedded"`
	UserInfo
}

type ThreadPreview struct {
	SharedID
	UnixTime   time.Time `json:"UnixTime" gorm:"autoCreateTime"`
	LastBump   time.Time `json:"LastBump" gorm:"autoCreateTime"`
	Name       string    `json:"Name"`
	Text       string    `json:"Text"`
	Topic      string    `json:"Topic"`
	Country    string    `json:"Country"`
	ExtraFlags string    `json:"ExtraFlags"`
	Sticky     bool      `json:"Sticky"`
	Page       int       `json:"Page"`
	PostCount  int       `json:"PostCount"`
	PostImage  PostImage `json:"PostImage,omitempty" gorm:"embedded"`
	UserInfo
	Posts []Post
}

type Post struct {
	SharedID
	UnixTime     time.Time `json:"UnixTime" gorm:"autoCreateTime"`
	Name         string    `json:"Name"`
	Text         string    `json:"Text"`
	Country      string    `json:"Country"`
	ExtraFlags   string    `json:"ExtraFlags"`
	ParentThread int64     `json:"ParentThread"`
	PostImage    PostImage `json:"PostImage,omitempty" gorm:"embedded"`
	UserInfo
}

type PostImage struct {
	Filename     string `json:"Filename"`
	ImageInfo    string `json:"ImageInfo"`
	ImageHash    string `json:"ImageHash"`
	OmitFilename bool   `json:"-"`
}
