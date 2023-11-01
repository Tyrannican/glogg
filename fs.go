package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

func CreateGloggRoot() *Workspace {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable to obtain user home directory: %v", err)
	}

	root := CreateDirectory(home, ".glogg")
	bin := CreateDirectory(root, "bin")
	packages := CreateDirectory(root, "packages")

	return &Workspace{
		WorkingDir: root,
		Bin:        bin,
		Packages:   packages,
	}
}

func CreateDirectory(root, target string) string {
	target = filepath.Join(root, target)
	os.MkdirAll(target, 0644)

	return target
}

func LoadTestManifest() *Manifest {
	fh, err := os.Open("test_synth.json")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	defer fh.Close()

	raw, _ := io.ReadAll(fh)
	var manifest *Manifest

	err = json.Unmarshal(raw, &manifest)
	if err != nil {
		log.Fatalf("unable to deserialize manifest: %v", err)
	}

	return manifest
}
