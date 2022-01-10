// Package api is the implementation of the MetaClean API version 1.
package api

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"meta-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Status(c *gin.Context) {
	c.Status(http.StatusOK)
}

func UploadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	filesUUID := uuid.New().String()
	errors := false

	filesFolder := filepath.Join("/tmp/MetaClean", filesUUID)
	processedFolder := "/tmp/MetaClean/files"
	folderCreationError := os.MkdirAll(filesFolder, 0755)
	filesFolderCreationError := os.MkdirAll(processedFolder, 0755)

	if folderCreationError != nil || filesFolderCreationError != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		for _, file := range files {
			// Move the file to a temp directory
			filePath := filepath.Join(filesFolder, file.Filename)
			fileError := c.SaveUploadedFile(file, filePath)

			if fileError != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				errors = true
				break
			}
		}

		if !errors {
			// Process the image(s) and provide a download link
			cleanFolderCommand := fmt.Sprintf("exiftool -all= -TagsFromFile @ -ColorSpaceTags -Orientation %s", filesFolder)
			exec.Command("/bin/sh", cleanFolderCommand).Output()

			zipErr := utils.Zip(filesFolder, fmt.Sprintf("%s/%s.zip", processedFolder, filesUUID))

			if zipErr == nil {
				scheme := ""
				if c.Request.TLS != nil {
					scheme = "s"
				}

				c.JSON(http.StatusCreated, gin.H{
					"url": fmt.Sprintf("http%s://%s%s?uuid=%s", scheme, c.Request.Host, c.Request.URL.Path, filesUUID),
				})
			}

		}

		// Remove the upload in a minute
		// Allow the user to download the processed files
		time.AfterFunc(time.Minute, func() {
			os.RemoveAll(filesFolder)
		})
	}
}

func DownloadFiles(c *gin.Context) {
	uuid, ok := c.GetQuery("uuid")

	if ok {
		filesFolder := filepath.Join("/tmp/MetaClean/files", fmt.Sprintf("%s.zip", uuid))
		_, filesExistError := os.Stat(filesFolder)

		if filesExistError == nil {
			c.File(filesFolder)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
