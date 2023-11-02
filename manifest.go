package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	BIN                = "BIN"
	PKG                = "PKG"
	WORKING_DIR        = "WKDIR"
	LINK               = "LINK"
	DIR_KEYWORDS       = []string{"BIN", "PKG"}
	OPERATION_KEYWORDS = []string{"LINK"}
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
type Action struct {
	Action []string `json:"action"`
}

func (b *BinaryBuild) ChecksumValidation(input []byte) bool {
	cksum := sha256.Sum256(input)
	hexSum := hex.EncodeToString(cksum[:])

	return hexSum == b.Sha256
}

func (a *Action) Prep(wk *Workspace) {
	action := a.Action

	for i, arg := range action {
		if strings.Contains(arg, BIN) {
			action[i] = strings.Replace(action[i], BIN, wk.Bin, -1)
		} else if strings.Contains(arg, PKG) {
			action[i] = strings.Replace(action[i], PKG, wk.Packages, -1)
		} else if strings.Contains(arg, WORKING_DIR) {
			action[i] = strings.Replace(action[i], WORKING_DIR, wk.WorkingDir, -1)
		}
	}
}

func (a *Action) Exec() {
	cmdName := a.Action[0]
	special, val := a.isSpecialAction(cmdName)

	if special {
		a.doSpecialAction(val, a.Action[1:])
	} else {
		cmd := exec.Command(cmdName, a.Action[1:]...)
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("unable to execute command %v", err)
		}
	}
}

func (a *Action) isSpecialAction(cmd string) (bool, string) {
	for _, special := range []string{LINK} {
		if cmd == special {
			return true, special
		}
	}

	return false, ""
}

func (a *Action) doSpecialAction(name string, args []string) {
	switch name {
	case LINK:
		src, dest := args[0], args[1]
		log.Printf("Symlinking %s to %s", src, dest)
		os.Symlink(src, dest)
	default:
		break
	}
}
