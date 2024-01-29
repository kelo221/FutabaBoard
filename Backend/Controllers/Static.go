package Controllers

import (
	"backend/Blueprints"
	db "backend/Database"
	DataModels "backend/ORM"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"strconv"
)

func ReturnStatic(c fiber.Ctx) error {
	return Render(&c, Blueprints.Empty())
}

func ReturnPage(c fiber.Ctx) error {
	pageParse := c.Params("*")

	if threadID, err := strconv.ParseInt(pageParse, 10, 64); err != nil {
		return err
	} else {
		var collection []DataModels.Thread
		if err = db.GetDB().Where("page = ?", threadID).Order("last_bump desc").Find(&collection).Error; err != nil {
			return err
		} else {
			var colWithPosts []DataModels.ThreadPreview

			for _, thread := range collection {
				var posts []DataModels.Post
				db.GetDB().Where("parent_thread = ?", thread.ID).Order("id desc").Limit(2).Find(&posts)

				//Swap order for preview
				if len(posts) > 1 {
					posts[0], posts[1] = posts[1], posts[0]
				}

				colWithPosts = append(colWithPosts, DataModels.ThreadPreview{
					SharedID:   thread.SharedID,
					UnixTime:   thread.UnixTime,
					LastBump:   thread.LastBump,
					Name:       thread.Name,
					Topic:      thread.Topic,
					Country:    thread.Country,
					ExtraFlags: thread.ExtraFlags,
					Sticky:     thread.Sticky,
					Page:       thread.Page,
					PostCount:  thread.PostCount,
					PostImage:  thread.PostImage,
					UserInfo:   thread.UserInfo,
					Posts:      posts,
					TextRaw:    thread.TextRaw,
				})
			}

			return Render(&c, Blueprints.ShowAll(colWithPosts))
		}
	}
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(*c)
}
