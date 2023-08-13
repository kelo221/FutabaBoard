package Controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/zeebo/blake3"
	"image"
	"image/png"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

const saltFileName = ".env"

// UserAccountSalt is used to salt API keys
var UserAccountSalt string

func generateSalt() (string, error) {
	genSalt := make([]byte, 16)
	_, err := rand.Read(genSalt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(genSalt), nil
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

	// Checks if there is a .env file present, if not generate UserAccountSalt for user keys
	if _, saltErr := os.Stat(saltFileName); os.IsNotExist(saltErr) {

		_salt, err := generateSalt()
		if err != nil {
			fmt.Println("Error generating UserAccountSalt:", err)
			os.Exit(1)
		}

		// Write the UserAccountSalt to the environment file
		err = writeSaltToFile(_salt)
		if err != nil {
			fmt.Println("Error writing UserAccountSalt to file:", err)
			os.Exit(1)
		}

		fmt.Println("Salt has been generated and saved to", saltFileName)
	} else {
		// Read the existing UserAccountSalt from the environment file
		var err error
		UserAccountSalt, err = readSaltFromFile()
		if err != nil {
			fmt.Println("Error reading UserAccountSalt from file:", err)
			os.Exit(1)
		}

	}

}

// GenThumbnail generates a thumbnail for the file passed into this function.
func GenThumbnail(file *multipart.FileHeader, id, imageHash string) error {

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		return err
	}

	thumbnail := imaging.Fit(img, 225, 150, imaging.Lanczos)

	dst, err := os.Create(filepath.Join("public/threadContent/"+id, imageHash+"_small.png"))
	if err != nil {
		return err
	}
	defer dst.Close()

	err = png.Encode(dst, thumbnail)

	if err != nil {
		return err
	}

	return nil
}

func PostImage(c *fiber.Ctx, imageHash string, parentThreadID string, file *multipart.FileHeader) error {

	extension := filepath.Ext(file.Filename)
	fileErr := c.SaveFile(file, filepath.Join("public/threadContent/"+parentThreadID, imageHash+extension))
	if fileErr != nil {
		log.Debug(fileErr)
	}

	thumbGenErr := GenThumbnail(file, parentThreadID, imageHash)
	if thumbGenErr != nil {
		log.Debug(thumbGenErr)
	}

	return nil
}

func PostOPImage(c *fiber.Ctx, imageHash string, file *multipart.FileHeader, ID string) error {

	// Image is saved in the thread's folder based on its hash value.
	if folderErr := os.Mkdir("public/threadContent/"+ID, os.ModePerm); folderErr != nil {
		log.Debug(folderErr)
	}

	extension := filepath.Ext(file.Filename)
	fileErr := c.SaveFile(file, filepath.Join("public/threadContent/"+ID, imageHash+extension))
	if fileErr != nil {
		log.Debug(fileErr)
	}

	thumbGenErr := GenThumbnail(file, ID, imageHash)
	if thumbGenErr != nil {
		log.Debug(thumbGenErr)
	}

	return nil
}

func imageCheck(c *fiber.Ctx, file *multipart.FileHeader) (error, string) {

	if !checkImagePrivilege("test") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials, or you lack the permission to post images.",
		}), ""
	}

	imageHash, err := generateHash(file, UserAccountSalt)
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

	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/gif":  true,
		/*"audio/aac":  true,
		"audio/ogg":  true,
		"audio/opus": true,*/

		/*		"video/mp4":  true,
				"video/webm": true,*/
	}

	if !allowedTypes[file.Header.Get("Content-Type")] {
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
			"error": "Not a JPEG, PNG, OGG,AAC, OPUS, GIF, MP4, WEBM or a WEBP.",
		}), ""
	}

	return err, imageHash
}

func generateHashShort(input, salt string) (string, error) {
	hash := blake3.New()
	_, err := hash.Write([]byte(input + salt))
	return hex.EncodeToString(hash.Sum(nil))[:6], err
}

func generateHash(file *multipart.FileHeader, salt string) (string, error) {

	src, openingErr := file.Open()
	if openingErr != nil {
		return "", openingErr
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {

		}
	}(src)

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
	_, err := hasher.Write(data)
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
