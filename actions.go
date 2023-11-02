package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	BIN         = "BIN"
	PKG         = "PKG"
	WORKING_DIR = "WKDIR"
	LINK        = "LINK"
)

type Action struct {
	Action []string `json:"action"`
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

	if !special {
		cmd := exec.Command(cmdName, a.Action[1:]...)
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("unable to execute command %v", err)
		}
	} else {
		a.doSpecialAction(val, a.Action[1:])
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
