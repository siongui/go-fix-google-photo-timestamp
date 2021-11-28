package main

import (
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func unixTimestampToGoTime(timestamp string) (gt time.Time, err error) {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return
	}
	gt = time.Unix(i, 0)
	return
}

func getFileModTime(dir, filename string) (t time.Time, err error) {
	fileInfo, err := os.Stat(filepath.Join(dir, filename))
	if err != nil {
		return
	}
	return fileInfo.ModTime(), err
}
