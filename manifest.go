package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os/exec"
)

type Manifest struct {
	Name        string      `json:"name"`
	Version     string      `json:"version"`
	Description string      `json:"description"`
	Home        string      `json:"home"`
	Repository  string      `json:"respository"`
	Binary      BinaryBuild `json:"binary"`
	Source      SourceBuild `json:"source"`
}

type BinaryBuild struct {
	Sha256   string   `json:"sha256"`
	Download string   `json:"download"`
	Actions  []Action `json:"actions"`
}

type SourceBuild struct {
	Requires []string `json:"requires"`
	Actions  []Action `json:"actions"`
}

func (b *BinaryBuild) ChecksumValidation(input []byte) bool {
	cksum := sha256.Sum256(input)
	hexSum := hex.EncodeToString(cksum[:])

	return hexSum == b.Sha256
}

func (s *SourceBuild) CheckRequirementsInstalled() {
	for _, req := range s.Requires {
		if !s.isRequirementInstalled(req) {
			log.Fatalf("%s is required to be installed before building from source", req)
		}
	}
}

func (s *SourceBuild) isRequirementInstalled(req string) bool {
	cmd := exec.Command("which", req)

	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}
