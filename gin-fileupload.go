package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./www")
	router.POST("/upload", func(c *gin.Context) {
		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("ERR: %s", err.Error()))
			return
		}

		filename := "./www/upload/"+filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("ERR: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("OK."))
	})
	router.Run(":8086")
}