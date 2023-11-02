package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Workspace struct {
    RootDir string
	WorkingDir string
	Bin        string
	Packages   string
}

func BinaryBuilder(wk *Workspace, manifest *Manifest) {
    workingDir := CreateDirectory(wk.Packages, fmt.Sprintf("%s/%s", manifest.Name, manifest.Version))
    os.Chdir(workingDir)
    wk.WorkingDir = workingDir

    pathSplit := strings.Split(manifest.Binary.Download, "/")
    downloadFile := pathSplit[len(pathSplit) - 1]
    log.Printf("Download file: %s\n", downloadFile)

    resp, err := http.Get(manifest.Binary.Download)
    if err != nil {
        log.Fatalf("unable to make GET request to download link %s -- %v", manifest.Binary.Download, err);
    }

    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if !manifest.Binary.ChecksumValidation(body) {
        log.Fatalf("!!! ERROR - Checksums do not match !!!\n")
    }

    os.WriteFile(fmt.Sprintf("%s/%s", workingDir, downloadFile), body, 0755)
    log.Println("Saved file!")

    for _, action := range manifest.Binary.Actions {
        action.Prep(wk)
        action.Exec()
    }
}

func main() {
    manifest := LoadTestManifest()
    wkspace := CreateGloggRoot()
    BinaryBuilder(wkspace, manifest)
}
