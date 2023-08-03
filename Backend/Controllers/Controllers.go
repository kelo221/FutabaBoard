package Controllers

import (
	db "backend/Database"
	DataModels "backend/ORM"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zeebo/blake3"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var postSalt string

const saltFileName = ".env"

var salt string

func generateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func writeSaltToFile(salt string) error {
	return os.WriteFile(saltFileName, []byte(salt), 0644)
}

func readSaltFromFile() (string, error) {
	data, err := os.ReadFile(saltFileName)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func init() {

	if _, saltErr := os.Stat(saltFileName); os.IsNotExist(saltErr) {

		_salt, err := generateSalt()
		if err != nil {
			fmt.Println("Error generating salt:", err)
			os.Exit(1)
		}

		// Write the salt to the environment file
		err = writeSaltToFile(_salt)
		if err != nil {
			fmt.Println("Error writing salt to file:", err)
			os.Exit(1)
		}

		fmt.Println("Salt has been generated and saved to", saltFileName)
	} else {
		// Read the existing salt from the environment file
		var err error
		salt, err = readSaltFromFile()
		if err != nil {
			fmt.Println("Error reading salt from file:", err)
			os.Exit(1)
		}

	}
}

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

func imageCheck(c *fiber.Ctx, file *multipart.FileHeader) (error, string) {

	if !checkImagePrivilege("test") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials, or you lack the permission to post images.",
		}), ""
	}

	imageHash, err := generateHash(file, salt)
	fmt.Println(imageHash)
	if err != nil {
		return err, ""
	}

	maxFileSize := 5 * 1024 * 1024 // 5MB in bytes

	// Check file size
	if file.Size > int64(maxFileSize) {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{
			"error": "File size exceeds the maximum limit (5MB).",
		}), ""
	}

	// Check file type (example: allow only JPEG, PNG, and GIF images)
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}

	if !allowedTypes[file.Header.Get("Content-Type")] {
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
			"error": "Only JPEG, PNG, and GIF images are allowed.",
		}), ""
	}

	return err, imageHash
}

func TestFunction(c *fiber.Ctx) error {

	var newThread DataModels.Thread

	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		if post := form.Value["jsonFile"]; len(post) > 0 {

			if marshErr := json.Unmarshal([]byte(post[0]), &newThread); marshErr != nil {
				return marshErr
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

			var hashErr error
			newThread.Hash, hashErr = generateHashShort(c.IP(), postSalt)
			if hashErr != nil {
				return hashErr
			}

			newThread.IpAddress = c.IP()

		}
		if file := form.File["image"]; len(file) > 0 {

			var imageErr error
			imageErr, newThread.PostImage.ImageHash = imageCheck(c, file[0])
			if imageErr != nil {
				return imageErr
			}

		}

	}

	db.GetDB().Create(&newThread)

	return nil
}

func checkImagePrivilege(hash string) bool {
	return true
}

func PostImage(c *fiber.Ctx) error {
	// Set a file size limit (5MB in this case)
	maxFileSize := 5 * 1024 * 1024 // 5MB in bytes

	// Parse the form file
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file provided or invalid file field name.",
		})
	}

	// Check file size
	if file.Size > int64(maxFileSize) {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{
			"error": "File size exceeds the maximum limit (5MB).",
		})
	}

	// Check file type (example: allow only JPEG, PNG, and GIF images)
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}
	if !allowedTypes[file.Header.Get("Content-Type")] {
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
			"error": "Only JPEG, PNG, and GIF images are allowed.",
		})
	}

	// Sanitize file name
	fileName := sanitizeFileName(file.Filename)

	// Save the uploaded file
	err = c.SaveFile(file, filepath.Join("public", fileName))
	if err != nil {
		return err
	}

	return c.SendString("File uploaded successfully!")
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

	return c.Status(http.StatusOK).JSON("User banned.")
}

func DeletePost(c *fiber.Ctx) error {
	threadIDString := c.Params("*")
	if threadID, err := strconv.ParseInt(threadIDString, 10, 64); err != nil {
		return err
	} else {
		var post DataModels.Post
		if result := db.GetDB().First(&post, threadID).Delete(&post); result.Error != nil {
			// If post is not found, assume if it is a thread
			var thread DataModels.Thread
			if result = db.GetDB().First(&thread, threadID).Delete(&post); result.Error != nil {
				return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
			}
			if result = db.GetDB().Where("parent_thread = ?", thread.ID).Delete(&DataModels.Post{}); result.Error != nil {
				return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
			}

			return c.JSON(thread)
		}
		return c.JSON(post)
	}
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

	var err error
	newThread.Hash, err = generateHashShort(c.IP(), postSalt)
	if err != nil {
		return err
	}

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

	if newPost.Country == "" {
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

	var err error
	newPost.Hash, err = generateHashShort(c.IP(), postSalt)
	if err != nil {
		return err
	}

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

func FetchThreadPreviews(c *fiber.Ctx) error {
	threadIDString := c.Params("*")

	if threadID, err := strconv.ParseInt(threadIDString, 10, 64); err != nil {
		return err
	} else {
		var collection []DataModels.Thread
		if err = db.GetDB().Where("page = ?", threadID).Find(&collection).Error; err != nil {
			return err
		} else {
			var colWithPosts []DataModels.ThreadPreview

			for _, thread := range collection {
				var posts []DataModels.Post
				db.GetDB().Where("parent_thread = ?", thread.ID).Order("id asc").Limit(2).Find(&posts)

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

func generateHashShort(input, salt string) (string, error) {
	hash := blake3.New()
	_, err := hash.Write([]byte(input + salt))
	return hex.EncodeToString(hash.Sum(nil))[:6], err
}

func generateHash(file *multipart.FileHeader, salt string) (string, error) {
	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	var buf []byte

	for {
		buffer := make([]byte, 1024)
		n, err := src.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			return "", err
		}
		buf = append(buf, buffer[:n]...)
	}

	data := append(buf, []byte(salt)...)

	hasher := blake3.New()
	_, err = hasher.Write(data)
	if err != nil {
		return "", err
	}
	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash, err
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

// SanitizeFileName removes any directory traversal characters from the given file name.
func sanitizeFileName(fileName string) string {
	// Remove any path components from the file name
	fileName = filepath.Base(fileName)

	// Replace any invalid characters with underscores
	invalidChars := []string{"\\", "/", ":", "*", "?", "\"", "<", ">", "|"}
	for _, invalidChar := range invalidChars {
		fileName = strings.ReplaceAll(fileName, invalidChar, "_")
	}

	return fileName
}
