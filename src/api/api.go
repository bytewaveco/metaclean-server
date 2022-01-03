// Package api is a wrapper to meta-server api functionality.
package api

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"meta-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	errors := false

	filesFolder := filepath.Join("/tmp", uuid.New().String())
	folderCreationError := os.Mkdir(filesFolder, 0755)

	if folderCreationError != nil {
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
			// Process the image and provide a download link
			cleanFolderCommand := fmt.Sprintf("exiftool -all= -TagsFromFile @ -ColorSpaceTags -Orientation %s", filesFolder)
			exec.Command("/bin/sh", cleanFolderCommand).Output()

			zipErr := utils.Zip(filesFolder, "archive.zip")

			if zipErr != nil {
				//

			} else {
				//
			}
			// time.AfterFunc(time.Second, func() {
			// 	os.RemoveAll(filesFolder)
			// })
		}
	}
}
