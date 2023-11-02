package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func BinaryBuilder(wk *Workspace, manifest *Manifest) {
	bin := manifest.Binary
	pathSplit := strings.Split(bin.Download, "/")
	downloadFile := pathSplit[len(pathSplit)-1]
	log.Printf("Download file: %s\n", downloadFile)

	resp, err := http.Get(bin.Download)
	if err != nil {
		log.Fatalf("unable to make GET request to download link %s -- %v", bin.Download, err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if !bin.ChecksumValidation(body) {
		log.Fatalf("!!! ERROR - Checksums do not match !!!\n")
	}

	os.WriteFile(fmt.Sprintf("%s/%s", wk.WorkingDir, downloadFile), body, 0755)
	log.Println("Saved file!")

	for _, action := range bin.Actions {
		action.Prep(wk)
		action.Exec()
	}
}
