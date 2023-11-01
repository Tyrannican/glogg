package main

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
