package helpers

import (
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"
)

func GetFilenamePhoto(file *multipart.FileHeader) string {
	// Set Format Save Photo in Temp
	extension := filepath.Ext(file.Filename)
	epoc := strconv.FormatInt(time.Now().Unix(), 10)
	randomFilename := RandomString(len(file.Filename))
	return randomFilename + "-" + epoc + extension
}