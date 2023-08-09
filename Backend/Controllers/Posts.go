package Controllers

import (
	db "backend/Database"
	DataModels "backend/ORM"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// postSalt is used to generate hash value for each post.
var postSalt string

func init() {
	var err error
	postSalt = generateRandomString(20)
	if err != nil {
		panic("Error generating random string: " + err.Error())
	}
}

func Register(c *fiber.Ctx) error {
	return nil
}

func Login(c *fiber.Ctx) error {
	return nil
}

func TestFunction(c *fiber.Ctx) error {
	return nil
}

func checkImagePrivilege(hash string) bool {
	return true
}

func stripHTMLTags(input string) string {
	reader := strings.NewReader(input)
	tokenizer := html.NewTokenizer(reader)
	var result strings.Builder

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return result.String()
		case html.TextToken:
			result.WriteString(tokenizer.Token().Data)
		}
	}
}

func formatPostText(text string) string {
	lines := strings.Split(text, "\n")
	var result strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			if strings.HasPrefix(line, ">") {
				result.WriteString("<p style=\"color: green;\">")
				result.WriteString("&gt;" + line[1:])
				result.WriteString("</p>\n")
			} else if strings.HasPrefix(line, "<") {
				result.WriteString("<p style=\"color: red;\">")
				result.WriteString("&lt;" + line[1:])
				result.WriteString("</p>\n")
			} else {
				result.WriteString("<p>")
				result.WriteString(line)
				result.WriteString("</p>\n")
			}
		}
	}

	return result.String()
}

func Ban(c *fiber.Ctx) error {

	var newBan DataModels.Bans
	if err := json.Unmarshal(c.Body(), &newBan); err != nil {
		return err
	}

	if newBan.IP == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing IP address.",
		})
	}

	if newBan.ExpiringTimeUnix.Before(time.Now()) {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Your expire date is in the past.",
		})
	}

	if time.Time.IsZero(newBan.ExpiringTimeUnix) {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing expiring date",
		})
	}

	if newBan.Reason == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing ban reason.",
		})
	}

	db.GetDB().Create(&newBan)

	// Retrieve all posts by the banned IP address
	var posts []DataModels.Post
	if result := db.GetDB().Where("ip_address = ?", newBan.IP).Find(&posts); result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	for _, post := range posts {
		err := deletePostHelper(post.ID)
		if err != nil {
			return err
		}
	}

	// Retrieve all posts by the banned IP address
	var threads []DataModels.Thread
	if result := db.GetDB().Where("ip_address = ?", newBan.IP).Find(&threads); result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	for _, thread := range threads {
		err := deletePostHelper(thread.ID)
		if err != nil {
			return err
		}
	}

	// TODO remove posts too!

	return c.Status(http.StatusOK).JSON("User banned.")
}

// Todo Remove individual files too, not only whole threads.
func deletePostHelper(postID int64) error {

	var post DataModels.Post
	if result := db.GetDB().First(&post, postID).Delete(&post); result.Error != nil {
		// If post is not found, assume if it is a thread
		var thread DataModels.Thread
		if result = db.GetDB().First(&thread, postID).Delete(&post); result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
		}
		if result = db.GetDB().Where("parent_thread = ?", thread.ID).Delete(&DataModels.Post{}); result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
		}
		removeErr := os.RemoveAll("public/ThreadContent/" + strconv.FormatInt(thread.ID, 10))
		if removeErr != nil {
			log.Debug(removeErr)
		}

	} else {
		return result.Error
	}

	return nil
}

func DeletePost(c *fiber.Ctx) error {
	postIDParse := c.Params("*")
	if postID, err := strconv.ParseInt(postIDParse, 10, 64); err != nil {
		return err
	} else {
		delErr := deletePostHelper(postID)
		if delErr != nil {
			return delErr
		}
		return c.Status(http.StatusOK).JSON("Deleted Posts.")
	}
}

// Thread TODO Do not allow regular users to sticky
func Thread(c *fiber.Ctx) error {

	banCheck, banExpires := isUserBanned(c.IP())
	if banCheck {
		return c.Status(http.StatusUnauthorized).JSON("You are banned. Ban expiring on: " + banExpires.String())
	}

	var newThread DataModels.Thread

	if form, err := c.MultipartForm(); err == nil {
		if post := form.Value["jsonFile"]; len(post) > 0 {

			if marshErr := json.Unmarshal([]byte(post[0]), &newThread); marshErr != nil {
				return marshErr
			}

			if newThread.Topic == "" {
				c.Status(400)
				return c.Status(400).JSON(fiber.Map{
					"message": "Missing topic for the thread.",
				})
			}

			if newThread.Name == "" {
				newThread.Name = "Anon"
			}

			if len(newThread.Text) > 3000 {
				return c.Status(http.StatusBadRequest).JSON("Post longer than 3000 characters.")
			}

			// Remove HTML tags and add quote coloring
			newThread.Text = formatPostText(stripHTMLTags(newThread.Text))

			var hashErr error
			newThread.UserHash, hashErr = generateHashShort(c.IP(), postSalt)
			if hashErr != nil {
				return hashErr
			}

			newThread.IpAddress = c.IP()
		}

		if file := form.File["image"]; len(file) > 0 {

			// ImageCheck function checks if the file is correct format and size. User privilege is verified.
			// Returns a UserHash of the file or an error.
			var imageErr error
			imageErr, newThread.PostImage.ImageHash = imageCheck(c, file[0])
			if imageErr != nil {
				return imageErr
			}

			imgPostErr := PostOPImage(c, newThread.PostImage.ImageHash, file[0])
			if imgPostErr != nil {
				return imgPostErr
			}

			// Filename is the original filename that is shown in the thread
			newThread.PostImage.Filename = sanitizeFileName(file[0].Filename)

			newThread.PostImage.ImageInfo = fmt.Sprintf("%.3f", float64(file[0].Size)/(1024*1024)) + " Mb " + filepath.Ext(file[0].Filename)
		}
	}

	db.GetDB().Create(&newThread)

	return c.Status(http.StatusOK).JSON("Created a new thread Successfully.")
}

// Post Creates a new post.
func Post(c *fiber.Ctx) error {

	banCheck, banExpires := isUserBanned(c.IP())
	if banCheck {
		return c.Status(http.StatusUnauthorized).JSON("You are banned. Ban expiring on: " + banExpires.String())
	}

	var newPost DataModels.Post

	if form, err := c.MultipartForm(); err == nil {
		if post := form.Value["jsonFile"]; len(post) > 0 {

			if marshErr := json.Unmarshal([]byte(post[0]), &newPost); marshErr != nil {
				return marshErr
			}

			if newPost.Name == "" {
				newPost.Name = "Anon"
			}

			if len(newPost.Text) > 3000 {
				return c.Status(http.StatusBadRequest).JSON("Post longer than 3000 characters.")
			}

			newPost.Text = formatPostText(stripHTMLTags(newPost.Text))
			var hashErr error
			newPost.UserHash, hashErr = generateHashShort(c.IP(), postSalt)
			if hashErr != nil {
				return hashErr
			}

			newPost.IpAddress = c.IP()
		}

		if file := form.File["image"]; len(file) > 0 {

			// ImageCheck function checks if the file is correct format and size. User privilege is verified.
			// Returns a UserHash of the file or an error.
			var imageErr error
			imageErr, newPost.PostImage.ImageHash = imageCheck(c, file[0])
			if imageErr != nil {
				return imageErr
			}

			var imageHashCheck DataModels.Post

			//Check if image already exists
			result := db.GetDB().Where("image_hash = ? AND parent_thread = ?", newPost.PostImage.ImageHash, newPost.ParentThread).Limit(1).Find(&imageHashCheck)
			if result.RowsAffected < 0 || imageHashCheck.ID == 0 {
				imgPostErr := PostImage(c, newPost.PostImage.ImageHash, strconv.FormatInt(newPost.ParentThread, 10), file[0])
				if imgPostErr != nil {
					return imgPostErr
				}
			} else {
				return c.Status(http.StatusBadRequest).JSON("Image already exists in this thread.")
			}

			// Filename is the original filename that is shown in the thread
			newPost.PostImage.Filename = sanitizeFileName(file[0].Filename)

			newPost.PostImage.ImageInfo = fmt.Sprintf("%.3f", float64(file[0].Size)/(1024*1024)) + " Mb " + filepath.Ext(file[0].Filename)
		}
	}

	db.GetDB().Create(&newPost)

	var lastPostDate DataModels.Thread
	if err := db.GetDB().Where("id = ?", newPost.ParentThread).Find(&lastPostDate).Error; err != nil {
		return err
	} else {
		db.GetDB().Model(&lastPostDate).Update("last_bump", time.Now())
	}

	return c.Status(http.StatusOK).JSON("Post Successful.")

}

func FetchThreadPreviews(c *fiber.Ctx) error {
	pageParse := c.Params("*")

	if threadID, err := strconv.ParseInt(pageParse, 10, 64); err != nil {
		return err
	} else {
		var collection []DataModels.Thread
		if err = db.GetDB().Where("page = ?", threadID).Find(&collection).Error; err != nil {
			return err
		} else {
			var colWithPosts []DataModels.ThreadPreview

			for _, thread := range collection {
				var posts []DataModels.Post
				db.GetDB().Where("parent_thread = ?", thread.ID).Order("id desc").Limit(2).Find(&posts)

				//Swap order for preview
				if len(posts) > 0 {
					posts[0], posts[1] = posts[1], posts[0]
				}

				colWithPosts = append(colWithPosts, DataModels.ThreadPreview{
					SharedID:   thread.SharedID,
					UnixTime:   thread.UnixTime,
					LastBump:   thread.LastBump,
					Name:       thread.Name,
					Text:       thread.Text,
					Topic:      thread.Topic,
					Country:    thread.Country,
					ExtraFlags: thread.ExtraFlags,
					Sticky:     thread.Sticky,
					Page:       thread.Page,
					PostCount:  thread.PostCount,
					PostImage:  thread.PostImage,
					UserInfo:   thread.UserInfo,
					Posts:      posts,
				})
			}

			return c.JSON(colWithPosts)
		}
	}
}

// FetchThread Returns a whole thread, input thread ID.
func FetchThread(c *fiber.Ctx) error {
	threadIDString := c.Params("*")

	if threadID, err := strconv.ParseInt(threadIDString, 10, 64); err != nil {
		return err
	} else {
		var collection []DataModels.Post
		if err = db.GetDB().Where("parent_thread = ?", threadID).Find(&collection).Error; err != nil {
			return err
		} else {
			return c.JSON(collection)
		}
	}
}

// FetchPost Returns a single post, used for seeing posts from a thread preview.
func FetchPost(c *fiber.Ctx) error {
	threadIDString := c.Params("*")
	if threadID, err := strconv.ParseInt(threadIDString, 10, 64); err != nil {
		return err
	} else {
		var post DataModels.Post
		if result := db.GetDB().First(&post, threadID); result.Error != nil {
			// If post is not found, assume if it is a thread
			var thread DataModels.Thread
			if result = db.GetDB().First(&thread, threadID); result.Error != nil {
				return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
			}
			return c.JSON(thread)
		}
		return c.JSON(post)
	}
}

func isUserBanned(ipAddress string) (bool, time.Time) {
	var banCheck DataModels.Bans

	result := db.GetDB().Where("ip = ?", ipAddress).Limit(1).Find(&banCheck)

	if result.RowsAffected > 0 {
		if banCheck.ExpiringTimeUnix.Before(time.Now()) {
			db.GetDB().Where("ip = ?", ipAddress).Limit(1).Delete(&banCheck)
			return false, banCheck.ExpiringTimeUnix
		}
	}
	return result.RowsAffected > 0, banCheck.ExpiringTimeUnix
}
