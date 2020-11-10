package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// update to manage /file/upload
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			_, _ = io.WriteString(w, "internal error")
			return
		}
		// write the index html to client
		_, err = io.WriteString(w, string(data))
		if err != nil {
			fmt.Printf("write index.html error, err:%s\n", err.Error())
		}
	} else if r.Method == "POST" {
		// read from file
		file, header, err := r.FormFile("file")
		if err != nil {
			// should return
			_, _ = io.WriteString(w, "read error")
			return
		}
		// test if file upload success
		fmt.Println(header.Filename)
		size := 100
		buffer := make([]byte, size)
		_, err = file.ReadAt(buffer, int64(size))
		if err != nil {
			_, _ = io.WriteString(w, "read error")
			return
		}
		fmt.Println(string(buffer))
		// learn defer of go
	}
}
