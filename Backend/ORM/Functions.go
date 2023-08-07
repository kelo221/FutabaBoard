package DataModels

import (
	"gorm.io/gorm"
)

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
		for i := range threadsToMove {
			if i%threadsPerPage == 0 {
				pageCount++
			}
			tx.Model(&threadsToMove[i]).Update("page", pageCount)
		}
	}
	return nil
}
