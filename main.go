package main

import "log"

type Workspace struct {
	WorkingDir string
	Bin        string
	Packages   string
}

func main() {
    manifest := LoadTestManifest()
    wkspace := CreateGloggRoot()

    log.Printf("Manifest: %v\nWorkspace: %v\n", manifest, wkspace)
}
