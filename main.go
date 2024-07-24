package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	port                = ":6002"
	upFolder            = "./up"
	accessPassword      = "qingfeng6"
	fileRetentionPeriod = 3 * time.Hour
)

func main() {
	http.HandleFunc("/", handleFileUpload)
	http.ListenAndServe(port, nil)
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// Check for cookie
		if r.Header.Get("Cookie") != "access_password=qingfeng6" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Parse multipart form without specifying a max memory
		err := r.ParseMultipartForm(0) // 0 means no max memory, files will be stored in temporary files
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Create the file in the up folder
		dst, err := os.Create(filepath.Join(upFolder, handler.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set up a timer to delete the file after 3 hours
		go func(filename string) {
			time.Sleep(fileRetentionPeriod)
			os.Remove(filename)
		}(dst.Name())

		fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)

	case "GET":
		// Check if the request is for a file download
		requestedFile := filepath.Join(upFolder, filepath.Base(r.URL.Path))
		if strings.HasPrefix(r.URL.Path, "/") && fileExists(requestedFile) {
			http.ServeFile(w, r, requestedFile)
			return
		}
		http.Error(w, "File not found", http.StatusNotFound)
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
