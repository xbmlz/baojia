package api

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xbmlz/baojia/utils/oss"
)

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileObj, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// filename is uuid.ext
	filename := uuid.New().String() + filepath.Ext(file.Filename)
	err = oss.UploadFile("baojia", filename, fileObj, file.Size)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "upload file failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "upload file success",
		"data": filename,
	})
}

func GetFile(c *gin.Context) {
	// path is /file/:name
	filename := c.Param("name")
	fileUrl, err := oss.GetFileURL("baojia", filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "download file failed," + err.Error(),
		})
		return
	}
	c.Redirect(http.StatusFound, fileUrl.String())
}
