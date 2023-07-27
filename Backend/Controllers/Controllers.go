package Controllers

import (
	db "backend/Database"
	DataModels "backend/ORM"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
	"strconv"
	"time"
)

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

func Ban(c *fiber.Ctx) error {
	return nil
}

func DeletePost(c *fiber.Ctx) error {
	return nil
}

// Thread TODO Do not allow regular users to sticky
func Thread(c *fiber.Ctx) error {

	banCheck, banExpires := isUserBanned(c.IP())
	if banCheck {
		return c.Status(http.StatusUnauthorized).JSON("You are banned. Ban expiring on: " + banExpires.String())
	}

	var newThread DataModels.Thread
	if err := json.Unmarshal(c.Body(), &newThread); err != nil {
		return err
	}

	if newThread.PostImage.Filename == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing post content, must include an image for OP.",
		})
	}

	if newThread.Topic == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing topic for the thread.",
		})
	}

	if newThread.Flags == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing flags.",
		})
	}

	newThread.Hash = generateMD5HashWithSalt(c.IP(), postSalt)
	newThread.IpAddress = c.IP()

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
	if err := json.Unmarshal(c.Body(), &newPost); err != nil {
		return err
	}

	if newPost.PostImage.Filename == "" && newPost.Text == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing post content, must include either an image or text.",
		})
	}

	if newPost.Flags == "" {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing flags.",
		})
	}

	if newPost.ParentThread == 0 {
		c.Status(400)
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing Parent Thread.",
		})
	}

	newPost.Hash = generateMD5HashWithSalt(c.IP(), postSalt)
	newPost.IpAddress = c.IP()

	db.GetDB().Create(&newPost)

	var lastPostDate DataModels.Thread
	if err := db.GetDB().Where("id = ?", newPost.ParentThread).Find(&lastPostDate).Error; err != nil {
		return err
	} else {
		fmt.Println(lastPostDate)
		db.GetDB().Model(&lastPostDate).Update("last_bump", time.Now())
		fmt.Println(lastPostDate)
	}

	return c.Status(http.StatusOK).JSON("Post Successful.")

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

// FetchThreadPreviews Returns OP post and three of the latest posts of a specific page.
func FetchThreadPreviews(c *fiber.Ctx) error {
	threadIDString := c.Params("*")

	if threadID, err := strconv.ParseInt(threadIDString, 10, 64); err != nil {
		return err
	} else {
		return c.SendString(strconv.FormatInt(threadID, 10))
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

func TestFunction(c *fiber.Ctx) error {
	ipAddress := c.IP()
	md5Hash := generateMD5HashWithSalt(c.IP(), postSalt)

	return c.SendString("Your IP address is: " + ipAddress + " And your hash is:" + md5Hash)
}

func isUserBanned(ipAddress string) (bool, time.Time) {
	var banCheck DataModels.Bans

	result := db.GetDB().Where("ip = ?", ipAddress).Limit(1).Find(&banCheck)

	return result.RowsAffected > 0, banCheck.ExpiringTimeUnix
}

func generateMD5HashWithSalt(input, salt string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(input + salt))
	return hex.EncodeToString(md5Hash.Sum(nil))[:5]
}

func generateRandomString(length int) string {
	byteSize := (length * 6) / 8
	if (length*6)%8 != 0 {
		byteSize++
	}

	randomBytes := make([]byte, byteSize)
	_, err := rand.Read(randomBytes)
	if err != nil {
		os.Exit(0)
	}

	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	randomString = randomString[:length]
	return randomString
}
