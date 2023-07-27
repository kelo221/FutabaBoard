package DataModels

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

const postsPerThread = 300
const threadsPerPage = 4

type SharedID struct {
	ID int64 `gorm:"primaryKey"`
}

type UserInfo struct {
	IpAddress string `json:"-"`
	Hash      string `json:"Hash"`
}

type Thread struct {
	SharedID
	UnixTime  time.Time `json:"UnixTime" gorm:"autoCreateTime"`
	LastBump  time.Time `json:"LastBump" gorm:"autoCreateTime"`
	Name      string    `json:"Name"`
	Text      string    `json:"Text"`
	Topic     string    `json:"Topic"`
	Flags     string    `json:"Flags"`
	Sticky    bool      `json:"Sticky"`
	Page      int       `json:"Page"`
	PostCount int       `json:"PostCount"`
	PostImage PostImage `json:"PostImage,omitempty" gorm:"embedded"`
	UserInfo
}

type Post struct {
	SharedID
	UnixTime     time.Time `json:"UnixTime" gorm:"autoCreateTime"`
	Name         string    `json:"Name"`
	Text         string    `json:"Text"`
	Flags        string    `json:"Flags"`
	ParentThread int64     `json:"ParentThread"`
	PostImage    PostImage `json:"PostImage,omitempty" gorm:"embedded"`
	UserInfo
}

type PostImage struct {
	Filename  string `json:"Filename"`
	ImageInfo string `json:"ImageInfo"`
}

type Bans struct {
	IP               string    `json:"IP"`
	ExpiringTimeUnix time.Time `json:"ExpiringTimeUnix" gorm:"autoCreateTime"`
}

func (t *Thread) BeforeCreate(tx *gorm.DB) error {
	var maxThreadID Thread
	if err := tx.Order("id desc").Limit(1).Find(&maxThreadID).Error; err != nil {
		return err
	}
	var maxPostID Post
	if err := tx.Order("id desc").Limit(1).Find(&maxPostID).Error; err != nil {
		return err
	}

	if maxPostID.SharedID.ID > maxThreadID.SharedID.ID {
		t.SharedID = SharedID{maxPostID.ID + 1}
	} else {
		t.SharedID = SharedID{maxThreadID.ID + 1}
	}

	return nil
}

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	var maxThreadID Thread
	if err := tx.Order("id desc").Limit(1).Find(&maxThreadID).Error; err != nil {
		return err
	}
	var maxPostID Post
	if err := tx.Order("id desc").Limit(1).Find(&maxPostID).Error; err != nil {
		return err
	}

	if maxPostID.SharedID.ID > maxThreadID.SharedID.ID {
		p.SharedID = SharedID{maxPostID.ID + 1}
	} else {
		p.SharedID = SharedID{maxThreadID.ID + 1}
	}

	return nil
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	var thread Thread

	// Increment Post Count
	if result := tx.First(&thread, p.ParentThread); result.Error != nil {
		return result.Error
	} else {
		tx.Model(&thread).Update("post_count", thread.PostCount+1)
	}

	err := calculatePages(tx)
	if err != nil {
		return err
	}

	// Bump the page
	if thread.Page != 0 && thread.PostCount < postsPerThread {
		tx.Model(&thread).Update("page", 0)
	}

	return nil
}

func (t *Thread) AfterCreate(tx *gorm.DB) error {
	if err := tx.Where("page > ?", 5).Delete(&Thread{}).Error; err != nil {
		return err
	}

	err := calculatePages(tx)
	if err != nil {
		return err
	}

	return nil
}

func calculatePages(tx *gorm.DB) error {
	var threadsToMove []Thread
	if err := tx.Order("last_bump desc").Find(&threadsToMove).Error; err != nil {
		return err
	} else {
		pageCount := -1
		fmt.Println(threadsToMove)
		for i := range threadsToMove {
			if i%threadsPerPage == 0 {
				pageCount++
			}
			tx.Model(&threadsToMove[i]).Update("page", pageCount)
		}
	}
	return nil
}
