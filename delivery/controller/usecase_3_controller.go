package controller

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"usecase-1/usecase"

	"github.com/gin-gonic/gin"
)

type Usecase3Controller struct {
	usecase3 usecase.Usecase3UseCase
	router   *gin.Engine
}

func (e *Usecase3Controller) createHandler(c *gin.Context) {
	// Retrieve file and filename from the request
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get file from request"})
		return
	}
	defer file.Close()

	// Retrieve filename from form data or use the original filename
	filename := c.Request.FormValue("filename")
	if filename == "" {
		filename = fileHeader.Filename // Use the original filename if not provided
	}

	// Create a directory to save files if it doesn't exist
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create directory"})
		return
	}

	// Create a new file in the uploads directory with the specified filename
	dst, err := os.Create(filepath.Join("./uploads", filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create file"})
		return
	}
	defer dst.Close()

	// Copy the file data to the new file
	if _, err := io.Copy(dst, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func (e *Usecase3Controller) listHandler(c *gin.Context) {
	files, err := filepath.Glob("./uploads/*")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to list files"})
		return
	}

	var filenames []string
	for _, file := range files {
		filenames = append(filenames, filepath.Base(file))
	}

	c.JSON(http.StatusOK, gin.H{"files": filenames})
}

func (e *Usecase3Controller) listByIdHandler(c *gin.Context) {
	filename := c.Param("filename")
	filepath := filepath.Join("./uploads", filename)

	// Check if the file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"message": "File not found"})
		return
	}

	c.File(filepath)
}

func NewU3Controller(usecase usecase.Usecase3UseCase, r *gin.Engine) *Usecase3Controller {
	controller := Usecase3Controller{
		router:   r,
		usecase3: usecase,
	}

	rg := r.Group("/upload/")
	rg.POST("/", controller.createHandler)
	rg.GET("/files", controller.listHandler)
	rg.GET("/files/:filename", controller.listByIdHandler)
	return &controller
}
