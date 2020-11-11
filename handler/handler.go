package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
			fmt.Println("read form file error: err%s\n", err.Error())
			return
		}
		defer file.Close()
		path := "./tmp/"
		_ = os.MkdirAll(path, os.ModePerm)

		newFile, err := os.OpenFile(path + header.Filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Println("open or create file error: err%s\n", err.Error())
			return
		}

		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Println("copy to file error: err%s\n", err.Error())
			return
		}

		defer newFile.Close()
		// learn defer of go
		h := w.Header()
		h.Set("REDIRECT", "REDIRECT")
		h.Set("CONTEXTPATH", "/file/upload/success")
	}
}
// if file load success, redirect to this
func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "upload succeed")
}
