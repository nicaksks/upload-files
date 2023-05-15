package routes

import (
	"cdn/backend/utils"
	"math/rand"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	domain = utils.Domain()
)

const (
	VIDEO_DIR = "files/videos/"
	IMAGE_DIR = "files/images/"
	GIF_DIR   = "files/gifs/"
	AUDIO_DIR = "files/audios/"
)

func Files(c *gin.Context) {
	file, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Error getting the file.",
		})
		return
	}
	Save(file, c)
}

func Save(file *multipart.FileHeader, c *gin.Context) {
	var folder string

	switch filepath.Ext(file.Filename) {
	case ".mp4":
		folder = VIDEO_DIR
	case ".png", ".jpg":
		folder = IMAGE_DIR
	case ".gif":
		folder = GIF_DIR
	case ".mp3":
		folder = AUDIO_DIR
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Unsupported file type.",
		})
		return
	}

	path := folder + strings.Replace(file.Filename, ".", strconv.Itoa(rand.Intn(10000))+".", 1)

	err := c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Unsupported file type.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"path":    domain + strings.Replace(folder, "files/", "", 1) + path,
	})
}
