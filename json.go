package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type GooglePhotoTakeoutMetadata struct {
	Title          string `json:"title"`
	PhotoTakenTime struct {
		Timestamp string `json:"timestamp"`
		Formatted string `json:"formatted"`
	} `json:"photoTakenTime"`
}

func decodeJsonMetadata(dir string, file os.DirEntry) (m GooglePhotoTakeoutMetadata, err error) {
	b, err := os.ReadFile(filepath.Join(dir, file.Name()))
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &m)
	return
}
