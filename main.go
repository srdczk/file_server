package main

import (
	"fmt"
	"github.com/srdczk/file_server/handler"
	"net/http"
)

func main() {
	// set handler for update
	http.HandleFunc("/file/upload", handler.UploadHandler)
	err := http.ListenAndServe(":9393", nil)
	if err != nil {
		fmt.Println("listen error: err:%s\n", err.Error())
	}
}
