package main

import (
	"bytes"
	"fmt"
	"ginboy/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestUploadInvalidExtension checks uploading files with invalid extensions
func TestUploadInvalidExtension(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mongoMockTestClient := &database.MongoClient{}
	router := setupRouter(mongoMockTestClient)

	formData := &bytes.Buffer{}
	multipartWriter, err := processTestFile("files", "testfile.txt", formData)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder() // // Create a new recorder to capture the response
	req, _ := http.NewRequest("POST", "/upload", formData)
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType()) // Set the content type of the request to HTTP multipart/ form-data
	router.ServeHTTP(responseRecorder, req)

	// Assert that the HTTP response code is 400 and the error message is as expected
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected HTTP status 400 Bad Request")
	assert.Contains(t, responseRecorder.Body.String(), "only Jpg, Jpeg and png allowed")

}

// TestUploadNoFileAddedInFormData checks uploading files with no valid formData key
func TestUploadNoFileAddedInFormData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mongoMockTestClient := &database.MongoClient{}
	router := setupRouter(mongoMockTestClient)

	formData := &bytes.Buffer{}
	multipartWriter, err := processTestFile("notcorrect", "testfile.txt", formData)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder() // // Create a new recorder to capture the response
	req, _ := http.NewRequest("POST", "/upload", formData)
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType()) // Set the content type of the request to HTTP multipart/ form-data
	router.ServeHTTP(responseRecorder, req)

	// Assert that the HTTP response code is 400 and the error message is as expected
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected HTTP status 400 Bad Request")
	assert.Contains(t, responseRecorder.Body.String(), "no files found")
}

// TODO TestUploadFileWithCorrectData

// processTestFile creates a multipart form file and writes dummy data to it
func processTestFile(formKey, filename string, formData *bytes.Buffer) (*multipart.Writer, error) {
	multipartWriter := multipart.NewWriter(formData)

	part, err := multipartWriter.CreateFormFile(formKey, filename)
	if err != nil {
		return nil, fmt.Errorf("error creating the form file: %w", err)
	}

	_, err = part.Write([]byte("dummy"))
	if err != nil {
		return nil, fmt.Errorf("error writing to the form file: %w", err)
	}

	multipartWriter.Close() // can't add more parts to the form data later in the test!

	return multipartWriter, nil
}
