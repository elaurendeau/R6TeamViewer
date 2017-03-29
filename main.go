package main

import (
	"net/http"
	"log"
	"os"
	"bytes"
)

func main() {
	log.Fatal(http.ListenAndServe(":80", http.FileServer(getWebPath())))
}

func getWebPath() http.Dir {
	var webPath bytes.Buffer
	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	webPath.WriteString(path)
	webPath.WriteString("../web/.")

	return http.Dir("../web/.")
}
