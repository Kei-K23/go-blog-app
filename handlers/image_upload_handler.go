package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
)

type ImagesUploadHandler struct {}

func (ImgU *ImagesUploadHandler) UploadHandler(c echo.Context) error {
	const MAX_UPLOAD_SIZE = 10 << 20 // Set the max upload size to 10 MB

	// Limit request body size
	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, MAX_UPLOAD_SIZE)

	// Parse the form data
	if err := c.Request().ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		return c.String(http.StatusBadRequest, "The uploaded file is too big. Please choose a file that's less than 10MB in size")
	}

	// Retrieve the file from form data
	file, fileHeader, err := c.Request().FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	defer file.Close()

	// Create the uploads folder if it doesn't exist
	err = os.MkdirAll("./images", os.ModePerm)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Create a new file in the uploads directory with a unique name
	filename := fmt.Sprintf("/images/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
	dst, err := os.Create("." + filename)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer dst.Close()

	// Copy the uploaded file to the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Send response with the file location
	response := map[string]string{"location": filename}
	return c.JSON(http.StatusCreated, response)
}

func (ImgU *ImagesUploadHandler) ServeImages(e echo.Context) error {	
	fmt.Println(e.Request().URL)
	
	filepath := e.Param("*")
	
 	return e.File("./images/" + filepath)
}