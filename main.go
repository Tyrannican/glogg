package main

import (
	"fmt"
)

type Workspace struct {
	RootDir    string
	WorkingDir string
	Bin        string
	Packages   string
}

func (w *Workspace) Prep(manifest *Manifest) {
	workingDir := CreateDirectory(w.Packages, fmt.Sprintf("%s/%s", manifest.Name, manifest.Version))
	w.WorkingDir = workingDir
}

func main() {
	manifest := LoadTestManifest()
	wkspace := CreateGloggRoot()
	wkspace.Prep(manifest)
	BinaryBuilder(wkspace, manifest)
}
