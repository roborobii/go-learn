package main

import (
	"fmt"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File\n")

	// 1. Parse Input; type multipart/form-data
	r.ParseMultipartForm(10 << 20)

	// 2. Retrieve file from posted form-data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving File from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// 3. Write temporary file on our server
	// tempFile, err := ioutil.TempFile("", "upload-*.png")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer tempFile.Close()

	// fileBytes, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// tempFile.Write(fileBytes)

	// 4. return whether or not
	fmt.Fprintf(w, "Successfully uploaded\n")

}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Go File Upload")
	setupRoutes()
}
